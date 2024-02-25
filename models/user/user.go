package user

import (
	"github.com/oigi/Magikarp/global"
)

type User struct {
	global.MODEL
	UID                int    `json:"uid" gorm:"index;comment:用户UID"`                      // 用户UID
	Password           string `json:"-"  gorm:"comment:用户登录密码"`                            // 用户登录密码
	NickName           string `json:"nickName" gorm:"default:系统用户;comment:用户昵称"`           // 用户昵称
	HeaderImg          string `json:"headerImg" gorm:"comment:用户头像"`                       // 用户头像
	Sex                int    `json:"sex" gorm:"default:1;comment:用户性别 1男 2女"`             //用户性别
	Email              string `json:"email"  gorm:"comment:用户邮箱"`                          // 用户邮箱
	Enable             int    `json:"enable" gorm:"default:1;comment:用户是否被冻结 0停用 1正常 2冻结"` //用户是否被冻结 1正常 2冻结
	DefaultFavoritesID uint   `json:"default_favorites_id" gorm:"comment:默认收藏夹ID"`
}
