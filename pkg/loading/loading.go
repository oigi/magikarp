package loading

import (
    "github.com/oigi/Magikarp/config"
    "github.com/oigi/Magikarp/pkg/logger"
    "github.com/oigi/Magikarp/pkg/viper"
    "go.uber.org/zap"
)

// Loading 全局loading
func Loading() {
    config.VIPER = viper.Viper()     // 初始化Viper
    config.LOG = logger.InitLogger() // 初始化zap日志库
    zap.ReplaceGlobals(config.LOG)
}
