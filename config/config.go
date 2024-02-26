package config

type Server struct {
	JWT     JWT     `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Captcha Captcha `mapstructure:"captcha" json:"captcha" yaml:"captcha"`
	System  System  `mapstructure:"system" json:"system" yaml:"system"`
	// gorm
	Mysql Mysql `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
}
