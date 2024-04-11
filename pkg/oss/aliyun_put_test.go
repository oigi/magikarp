package oss

import (
	"fmt"
	"github.com/oigi/Magikarp/config"
	"github.com/spf13/viper"
	"os"
	"testing"
)

func TestPut(t *testing.T) {
	file, err := os.Open("/Users/aubyn/Downloads/45-从0开始学游戏开发/05-第四章：脚本语言 (3讲)/第24讲丨如何嵌入脚本语言？.pdf")
	if err != nil {
		return
	}

	v := viper.New()
	// 设置配置文件路径
	v.SetConfigType("yaml")
	v.SetConfigName("oss")
	v.AddConfigPath(".")

	// 读取配置文件
	err = v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error model file: %s \n", err))
	}

	if err = v.Unmarshal(&config.CONFIG); err != nil {
		fmt.Println(err)
	}

	client, err := InitOss()
	if err != nil {
		return
	}

	err = client.Put("video", "test.pdf", file)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
}
