package config

import (
	"github.com/oigi/Magikarp/config/model"
)

type Server struct {
	JWT     model.JWT     `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Zap     model.Zap     `mapstructure:"zap" json:"zap" yaml:"zap"`
	Redis   model.Redis   `mapstructure:"redis" json:"redis" yaml:"redis"`
	Captcha model.Captcha `mapstructure:"captcha" json:"captcha" yaml:"captcha"`
	System  model.System  `mapstructure:"system" json:"system" yaml:"system"`
	Etcd    *model.Etcd   `mapstructure:"etcd" json:"etcd" yaml:"etcd"`
	// gorm
	Mysql model.Mysql `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
}
