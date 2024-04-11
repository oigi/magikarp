package oss

import (
	"fmt"
	"github.com/oigi/Magikarp/config"
	"github.com/spf13/viper"
	"testing"
)

func TestInitOss(t *testing.T) {
	v := viper.New()
	// 设置配置文件路径
	v.SetConfigType("yaml")
	v.SetConfigName("oss")
	v.AddConfigPath(".")

	// 读取配置文件
	err := v.ReadInConfig()
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

	if client == nil || client.Client == nil || client.Bucket == nil {
		fmt.Println("asdasdasd")
		t.Fatal("OSS client or bucket is nil")
	}

	fmt.Println(client.Bucket)
}
