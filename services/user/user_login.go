package user

import (
	"errors"
	"github.com/oigi/Magikarp/global"
	"github.com/oigi/Magikarp/models/user"
	"github.com/oigi/Magikarp/utils"
	"gorm.io/gorm"
)

type UService struct{}

// Login 登陆
func (userService *UService) Login(u *user.User) (userInter *user.User, err error) {
	var user user.User
	err = global.DB.Where("email = ?", u.Email).First(&user).Error
	if err == nil {
		if ok := utils.BcryptCheck(u.Password, user.Password); !ok {
			return nil, errors.New("密码错误")
		}
	}
	return &user, err
}

// Register 注册
func (userService *UService) Register(u *user.User) (userInter *user.User, err error) {
	var user user.User
	if !errors.Is(global.DB.Where("nick_name = ?", u.NickName).First(&user).Error, gorm.ErrRecordNotFound) { // 判断昵称是否注册
		return userInter, errors.New("昵称已注册")
	}

	if !errors.Is(global.DB.Where("email = ?", u.Email).First(&user).Error, gorm.ErrRecordNotFound) {
		return userInter, errors.New("邮箱已注册")
	}
	// 密码hash加密 注册
	u.Password = utils.BcryptHash(u.Password)
	err = global.DB.Create(&u).Error
	return u, err
}

// ChangePassword 改密
func (userService *UService) ChangePassword(u *user.User, newPassword string) (userInter *user.User, err error) {
	var user user.User
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
func (userService *UService) DeleteUser(id int) (err error) {
	var user user.User
	err = global.DB.Where("id = ?", id).Update("enable", 0).Error
	if err != nil {
		return err
	}

	// 更新 NickName 为 "该用户已删除"
	err = global.DB.Model(&user).Where("id = ?", id).Update("nick_name", "该用户已删除").Error
	if err != nil {
		return err
	}
	return nil
}
