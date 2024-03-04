package model

import (
	"gorm.io/gorm"
	"time"
)

type Model struct {
	ID        int64          `gorm:"primary_key;AUTO_INCREMENT:10000"`
	CreatedAt time.Time      // 创建时间
	UpdatedAt time.Time      // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // 删除时间
}
