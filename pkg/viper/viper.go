package viper

import (
	"fmt"
	"github.com/oigi/Magikarp/config"
	"github.com/spf13/viper"
)

func Viper() *viper.Viper {
	v := viper.New()
	// 设置配置文件路径
	v.AddConfigPath("model")
	v.SetConfigType("yaml")

	// 读取配置文件
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error model file: %s \n", err))
	}

	// 解析整个配置到 model.CONFIG
	if err = v.Unmarshal(&config.CONFIG); err != nil {
		fmt.Println(err)
	}

	return v
}
