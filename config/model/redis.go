package model

type Redis struct {
	DB       int    `mapstructure:"dao" json:"dao" yaml:"dao"`                // redis的哪个数据库
	Addr     string `mapstructure:"addr" json:"addr" yaml:"addr"`             // 服务器地址:端口
	Password string `mapstructure:"password" json:"password" yaml:"password"` // 密码
}
