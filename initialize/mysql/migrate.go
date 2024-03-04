package mysql

import (
	"github.com/oigi/Magikarp/app/user/internal/model/user"
	"os"
)

func migration() {
	err := _db.AutoMigrate(
		user.User{},
		//Todo 添加其他的
	)
	if err != nil {
		os.Exit(0)
	}
}
