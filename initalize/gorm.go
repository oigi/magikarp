package initalize

import (
	"github.com/oigi/Magikarp/global"
	"github.com/oigi/Magikarp/models/user"
	"gorm.io/gorm"
	"os"
)

func Gorm() *gorm.DB {
	return GormMysql()
}

func RegisterTables() {
	db := global.DB
	err := db.AutoMigrate(
		user.User{},
		//Todo 添加其他的
	)
	if err != nil {
		os.Exit(0)
	}
}
