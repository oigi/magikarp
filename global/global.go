package global

import (
	"github.com/oigi/Magikarp/config"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	DB     *gorm.DB
	VIPER  *viper.Viper
	CONFIG config.Server
)
