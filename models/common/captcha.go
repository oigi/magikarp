package common

import "time"

// Captcha 表示 captcha 表
type Captcha struct {
    UUID       string    `json:"uuid" gorm:"comment:UUID"`
    Code       string    `json:"code" gorm:"comment:验证码"`
    ExpireTime time.Time `json:"expire_time" gorm:"comment:过期时间"`
}
