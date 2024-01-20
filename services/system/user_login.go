package system

import (
	"errors"
	"github.com/google/uuid"
	"github.com/oigi/Magikarp/global"
	"github.com/oigi/Magikarp/models"
	"github.com/oigi/Magikarp/utils"
	"gorm.io/gorm"
)

type UserService struct{}

// Login 登陆
func (userService *UserService) Login(u *models.User) (userInter *models.User, err error) {
	var user models.User
	err = global.DB.Where("username = ? OR email = ?", u.Username, u.Email).First(&user).Error
	if err == nil {
		if ok := utils.BcryptCheck(u.Password, user.Password); !ok {
			return nil, errors.New("密码错误")
		}
	}
	return &user, err
}

// Register 注册
func (userService *UserService) Register(u *models.User) (userInter *models.User, err error) {
	var user models.User
	if !errors.Is(global.DB.Where("username = ?", u.Username).First(&user).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return userInter, errors.New("用户名已注册")
	}

	if !errors.Is(global.DB.Where("email = ?", u.Email).First(&user).Error, gorm.ErrRecordNotFound) {
		return userInter, errors.New("邮箱已注册")
	}
	// 附加uuid 密码hash加密 注册
	u.Password = utils.BcryptHash(u.Password)
	u.UUID = uuid.Must(uuid.NewV7())
	err = global.DB.Create(&u).Error
	return u, err
}

// ResetPassword 修改用户密码
//func (userService *UserService) ResetPassword(ID uint) (err error) {
//
//	err = global.DB.Model(&models.User{}).Where("id = ?").Update("password", utils.BcryptHash("123456")).Error
//	return err
//}

func (userService *UserService) ChangePassword(u *models.User, newPassword string) (userInter *models.User, err error) {
	var user models.User
	if err = global.DB.Where("id = ?", u.ID).First(&user).Error; err != nil {
		return nil, err
	}
	if ok := utils.BcryptCheck(u.Password, user.Password); !ok {
		return nil, errors.New("原密码错误")
	}
	user.Password = utils.BcryptHash(newPassword)
	err = global.DB.Save(&user).Error
	return &user, err
}

// DeleteUser 删除用户
func (userService *UserService) DeleteUser(id int) (err error) {
	var user models.User
	err = global.DB.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return err
	}
	// Todo 删除关系
	return err
}

// FindUserById 通过id获取用户信息

// FindUserByUuid 通过uuid获取用户信息

// _GetUserInfoList 分页获取数据

// SetUserAuthority 设置一个用户的权限

// SetUserInfo 设置用户信息

// GetUserInfo 获取用户信息
