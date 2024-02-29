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
    err = dao.Model(&userModel.User{}).Where("user_name=?", req.Email).
        First(&r).Error
    if err != nil {
        err = errors.Wrapf(err, "failed to get user info, userName = %v", req.Email)
    }
    return
}

// CreateUser 创建用户
func (dao *UserDao) CreateUser(req *user.UserRegisterReq) (err error) {
    var user userModel.User

    if !errors.Is(dao.Where("nick_name = ?", user.NickName).First(&user).Error, gorm.ErrRecordNotFound) { // 判断昵称是否注册
        return errors.New("昵称已注册")
    }

    if !errors.Is(dao.Where("email = ?", user.Email).First(&user).Error, gorm.ErrRecordNotFound) {
        return errors.New("邮箱已注册")
    }
    // 密码hash加密 注册
    user.Password = utils.BcryptHash(user.Password)
    if err = dao.Create(&user).Error; err != nil {
        return errors.Wrap(err, "failed to create user")
    }
    return err
}