package model

import (
	"gorm.io/gorm"
	"time"
)

type Model struct {
	ID        int64          `gorm:"auto_increment;primarykey;AUTO_INCREMENT=1000"`
	CreatedAt time.Time      // 创建时间
	UpdatedAt time.Time      // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // 删除时间
}
