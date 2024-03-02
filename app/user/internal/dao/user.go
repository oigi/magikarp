package dao

import (
	"context"
	"github.com/oigi/Magikarp/grpc/pb/user"
	"github.com/oigi/Magikarp/initialize/mysql"
	userModel "github.com/oigi/Magikarp/models/user"
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

// GetUserInfo 获取用户信息
func (dao *UserDao) GetUserInfo(req *user.UserLoginReq) (r *userModel.User, err error) {
	err = dao.Model(&userModel.User{}).Where("email=?", req.Email).
		First(&r).Error
	if err != nil {
		err = errors.Wrapf(err, "failed to get user info, userName = %v", req.Email)
	}
	return
}

// CreateUser 创建用户
func (dao *UserDao) CreateUser(req *user.UserRegisterReq) (err error) {
	var user userModel.User
	if req.NickName == "" {
		return errors.New("昵称不能为空")
	}

	if req.Email == "" {
		return errors.New("邮箱不能为空")
	}

	if !errors.Is(dao.Where("nick_name = ?", req.NickName).First(&user).Error, gorm.ErrRecordNotFound) { // 判断昵称是否注册
		return errors.New("昵称已注册")
	}

	if !errors.Is(dao.Where("email = ?", req.Email).First(&user).Error, gorm.ErrRecordNotFound) {
		return errors.New("邮箱已注册")
	}
	// 密码hash加密 注册
	user = userModel.User{
		Email:    req.Email,
		Password: utils.BcryptHash(req.Password),
		NickName: req.NickName,
	}

	if err = dao.Create(&user).Error; err != nil {
		return errors.Wrap(err, "failed to create user")
	}
	return err
}
