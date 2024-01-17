package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UUID      uuid.UUID `json:"uuid" gorm:"index;comment:用户UUID"`                // 用户UUID
	Email     string    `json:"mail" gorm:"index;comment:邮箱"`                    //邮箱
	Username  string    `json:"userName" gorm:"index;comment:用户登录名"`             // 用户登录名
	Password  string    `json:"-"  gorm:"comment:用户登录密码"`                        // 用户登录密码
	NickName  string    `json:"nickName" gorm:"comment:用户昵称"`                    // 用户昵称
	HeaderImg string    `json:"headerImg" gorm:"comment:用户头像"`                   // 用户头像
	Enable    int       `json:"enable" gorm:"default:1;comment:用户是否被冻结 1正常 2冻结"` //用户是否被冻结 1正常 2冻结
	Rules     string    `json:"rules" gorm:"comment:权限"`
}

//type UserSubscribe struct {
//	ID     int64 `json:"id" gorm:"primary_key;autoIncrement;comment:主键：ID"` // 主键：ID
//	TypeID int64 `json:"type_id" gorm:"comment:类型ID"`                       // 类型ID
//	UserID int64 `json:"user_id" gorm:"comment:用户ID"`                       // 用户ID
//}
//
//type UserRole struct {
//	ID     int64 `json:"id" gorm:"primary_key;autoIncrement;comment:主键：ID"` // 主键：ID
//	RoleID int64 `json:"role_id" gorm:"comment:角色ID"`                       // 角色ID
//	UserID int64 `json:"user_id" gorm:"comment:用户ID"`                       // 用户ID
//}
