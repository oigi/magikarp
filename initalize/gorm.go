package initalize

import (
	"github.com/oigi/Magikarp/global"
	"github.com/oigi/Magikarp/models/system"
	"gorm.io/gorm"
	"os"
)

func Gorm() *gorm.DB {
	return GormMysql()
}

func RegisterTables() {
	db := global.DB
	err := db.AutoMigrate(
		system.User{},
	)
	if err != nil {
		os.Exit(0)
	}
}
