package core

import (
	"fmt"
	"github.com/oigi/Magikarp/core/internal"
	"github.com/oigi/Magikarp/global"
	"github.com/oigi/Magikarp/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func Zap() (logger *zap.Logger) {
	if ok, _ := utils.PathExists(global.CONFIG.Zap.Director); !ok { // 判断是否有Director文件夹
		fmt.Printf("create %v directory\n", global.CONFIG.Zap.Director)
		_ = os.Mkdir(global.CONFIG.Zap.Director, os.ModePerm)
	}

	cores := internal.Zap.GetZapCores()
	logger = zap.New(zapcore.NewTee(cores...))

	if global.CONFIG.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}
