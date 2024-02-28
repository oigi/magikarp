package logger

import (
    "fmt"
    "github.com/oigi/Magikarp/config"
    "github.com/oigi/Magikarp/pkg/utils"
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
    "os"
)

func InitLogger() (logger *zap.Logger) {
    if ok, _ := utils.PathExists(config.CONFIG.Zap.Director); !ok { // 判断是否有Director文件夹
        fmt.Printf("create %v directory\n", config.CONFIG.Zap.Director)
        _ = os.Mkdir(config.CONFIG.Zap.Director, os.ModePerm)
    }

    cores := Zap.GetZapCores()
    logger = zap.New(zapcore.NewTee(cores...))

    if config.CONFIG.Zap.ShowLine {
        logger = logger.WithOptions(zap.AddCaller())
    }
    return logger
}
