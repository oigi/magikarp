package system

import "github.com/oigi/Magikarp/global"

type JwtBlacklist struct {
	global.MODEL
	Jwt string `gorm:"type:text;comment:jwt"`
}
