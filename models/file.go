package models

import (
	"gorm.io/gorm"
)

type File struct {
	gorm.Model
	FileKey  string `json:"file_key" gorm:"comment:文件键"` // 文件键
	Format   string `json:"format" gorm:"comment:格式"`    // 格式
	Type     string `json:"type" gorm:"comment:类型"`      // 类型
	Duration string `json:"duration" gorm:"comment:时长"`  // 时长
	Size     int64  `json:"size" gorm:"comment:大小"`      // 大小
	UserID   int64  `json:"user_id" gorm:"comment:用户ID"` // 用户ID
}
