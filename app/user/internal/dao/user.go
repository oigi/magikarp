package dao

import (
	"context"
	userModel "github.com/oigi/Magikarp/app/user/internal/model"
	"github.com/oigi/Magikarp/grpc/pb/user"
	"github.com/oigi/Magikarp/pkg/mysql"
	"github.com/oigi/Magikarp/pkg/utils"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type UserDao struct {
	*gorm.DB
}

func NewUserDao(ctx context.Context) *UserDao {
	return &UserDao{mysql.NewDBClient(ctx)}
}

// GetUserInfoById 根据用户id获取用户信息
func (db *UserDao) GetUserInfoById(req *user.GetUserByIdReq) (r *userModel.User, err error) {
	r = &userModel.User{}
	err = db.Model(&r).Where("id=?", req.UserId).First(&r).Error
	if err != nil {
		err = errors.Wrapf(err, "failed to get user info, id = %v", req.UserId)
		return nil, err
	}
	return
}

// GetUserInfo 获取用户信息
func (db *UserDao) GetUserInfo(req *user.UserLoginReq) (r *userModel.User, err error) {
	r = &userModel.User{}
	err = db.Where("email=?", req.Email).First(&r).Error
	if err != nil {
		err = errors.Wrapf(err, "failed to get user info, email = %v", req.Email)
		return nil, err
	}

	// 解密数据库中存储的密码
	if ok := utils.BcryptCheck(req.Password, r.Password); !ok {
		return nil, errors.Wrap(err, "failed to compare passwords")
	}
	return
}

// CreateUser 创建用户
func (db *UserDao) CreateUser(req *user.UserRegisterReq) (id int64, err error) {
	var user userModel.User

	if !errors.Is(db.Where("nick_name = ?", req.NickName).First(&user).Error, gorm.ErrRecordNotFound) {
		return 0, errors.New("昵称已注册")
	}

	user = userModel.User{}

	if !errors.Is(db.Where("email = ?", req.Email).First(&user).Error, gorm.ErrRecordNotFound) {
		return 0, errors.New("邮箱已注册")
	}

	// 密码hash加密并创建新用户，确保使用的是新的user变量
	user = userModel.User{
		Email:    req.Email,
		Password: utils.BcryptHash(req.Password),
		NickName: req.NickName,
	}

	if err = db.Create(&user).Error; err != nil {
		return 0, errors.Wrap(err, "failed to create user")
	}
	return user.ID, nil
}
