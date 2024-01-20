package core

import (
	"fmt"
	"github.com/oigi/Magikarp/global"
	"github.com/spf13/viper"
)

func Viper() *viper.Viper {
	v := viper.New()
	// 设置配置文件路径
	v.AddConfigPath("config")
	v.SetConfigType("yaml")

	// 读取配置文件
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	if err = v.Unmarshal(&global.CONFIG); err != nil {
		fmt.Println(err)
	}

	return v
}
