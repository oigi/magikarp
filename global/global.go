package global

import (
	"github.com/oigi/Magikarp/config"
	"gorm.io/gorm"
)

var (
	DB     *gorm.DB
	CONFIG config.Server
)
