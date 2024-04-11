package config

import (
	"github.com/oigi/Magikarp/config/model"
)

type Config struct {
	JWT     *model.JWT     `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Zap     *model.Zap     `mapstructure:"zap" json:"zap" yaml:"zap"`
	Captcha *model.Captcha `mapstructure:"captcha" json:"captcha" yaml:"captcha"`
	System  *model.System  `mapstructure:"system" json:"system" yaml:"system"`
	Etcd    *model.Etcd    `mapstructure:"etcd" json:"etcd" yaml:"etcd"`
	Kafka   *model.Kafka   `mapstructure:"kafka" json:"kafka" yaml:"kafka"`
	Oss     *model.Oss     `mapstructure:"oss" json:"oss" yaml:"oss"`
	Email   *model.Email   `mapstructure:"email" json:"email" yaml:"email"`

	Mongo *model.Mongo `mapstructure:"mongo" json:"mongo" yaml:"mongo"`
	Mysql *model.Mysql `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Redis *model.Redis `mapstructure:"redis" json:"redis" yaml:"redis"`
}
