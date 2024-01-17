package models

import "time"

type Captcha struct {
	UUID       string    `json:"uuid" gorm:"primary_key;comment:主键：UUID"` // 主键：UUID
	Code       string    `json:"code" gorm:"not null;comment:验证码"`        // 验证码
	ExpireTime time.Time `json:"expire_time" gorm:"comment:过期时间"`         // 过期时间
}
