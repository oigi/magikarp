package main

import (
	"github.com/oigi/Magikarp/core"
	"github.com/oigi/Magikarp/global"
	"github.com/oigi/Magikarp/initalize"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

func main() {
	global.VIPER = core.Viper() // 初始化Viper

	global.DB = initalize.Gorm()

	initalize.RegisterTables() // 初始化表
	db, _ := global.DB.DB()
	defer db.Close()

	router := initalize.Routers()
	router.Run(":8080")
}
