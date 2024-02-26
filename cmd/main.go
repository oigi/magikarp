package main

import (
	"github.com/oigi/Magikarp/core"
	"github.com/oigi/Magikarp/global"
	"github.com/oigi/Magikarp/initialize"
	"go.uber.org/zap"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

func main() {
	global.VIPER = core.Viper() // 初始化Viper

	global.LOG = core.Zap() // 初始化zap日志库
	zap.ReplaceGlobals(global.LOG)

	global.DB = initialize.Gorm()
	initialize.Redis()          // 初始化redis
	initialize.RegisterTables() // 初始化表
	db, _ := global.DB.DB()
	defer db.Close()

	router := initialize.Routers()
	router.Run(":8080")
}
