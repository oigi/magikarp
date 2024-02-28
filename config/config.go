package config

var conf *Server

type Server struct {
    JWT     JWT     `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
    Zap     Zap     `mapstructure:"zap" json:"zap" yaml:"zap"`
    Redis   Redis   `mapstructure:"redis" json:"redis" yaml:"redis"`
    Captcha Captcha `mapstructure:"captcha" json:"captcha" yaml:"captcha"`
    System  System  `mapstructure:"system" json:"system" yaml:"system"`

    // gorm
    Mysql Mysql `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
}
