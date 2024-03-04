package user

import (
	"github.com/oigi/Magikarp/consts"
)

type User struct {
	ID                 int64  `gorm:"primarykey; autoIncrement:10000"`                      // 主键ID
	Email              string `json:"email" gorm:"not null;uniqueIndex;comment:用户邮箱"`       // 用户邮箱
	Password           string `json:"-" gorm:"not null;comment:用户登录密码"`                     // 用户登录密码
	NickName           string `json:"nickName" gorm:"not null;default:系统用户;comment:用户昵称"`   // 用户昵称
	Avatar             string `json:"avatar" gorm:"comment:用户头像"`                           // 用户头像
	Sex                int    `json:"sex" gorm:"type:tinyint;default:1;comment:用户性别 1男 2女"` // 用户性别 1男 2女
	Dec                string `json:"dec" gorm:"comment:'个性签名'"`                            // 个性签名
	Enable             int    `json:"enable" gorm:"default:1;comment:用户是否被冻结 0停用 1正常 2冻结"`  // 用户是否被冻结 1正常 2冻结
	DefaultFavoritesID uint   `json:"default_favorites_id" gorm:"comment:默认收藏夹ID"`
	consts.Model
}

type TokenData struct {
	ID           int64  `json:"id"`
	Email        string `json:"email"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
