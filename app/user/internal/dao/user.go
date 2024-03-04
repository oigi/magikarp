package dao

import (
    "context"
    userModel "github.com/oigi/Magikarp/app/user/internal/model"
    "github.com/oigi/Magikarp/grpc/pb/user"
    "github.com/oigi/Magikarp/initialize/mysql"
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
func (db *UserDao) GetUserInfo(req *user.UserLoginReq) (r *userModel.User, err error) {
    err = db.Model(&r).Where("email=?", req.Email).First(&r).Error
    if err != nil {
        err = errors.Wrapf(err, "failed to get user info, email = %v", req.Email)
        return nil, err
    }
    // 解密数据库中存储的密码
    if ok := utils.BcryptCheck(req.Password, r.Password); !ok {
        return nil, errors.Wrap(err, "failed to compare passwords")
    }
    return r, err
}

// CreateUser 创建用户
func (db *UserDao) CreateUser(req *user.UserRegisterReq) (err error) {
    var user userModel.User

    if !errors.Is(db.Where("nick_name = ?", req.NickName).First(&user).Error, gorm.ErrRecordNotFound) { // 判断昵称是否注册
        return errors.New("昵称已注册")
    }

    if !errors.Is(db.Where("email = ?", req.Email).First(&user).Error, gorm.ErrRecordNotFound) {
        return errors.New("邮箱已注册")
    }
    // 密码hash加密 注册
    user = userModel.User{
        Email:    req.Email,
        Password: utils.BcryptHash(req.Password),
        NickName: req.NickName,
    }

    if err = db.Create(&user).Error; err != nil {
        return errors.Wrap(err, "failed to create user")
    }
    return err
}
