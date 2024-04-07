package model

type Redis struct {
	Dbs      map[string]int `mapstructure:"dbs" json:"dbs" yaml:"dbs"`                // 多个数据库
	Addr     string         `mapstructure:"addr" json:"addr" yaml:"addr"`             // 服务器地址:端口
	Password string         `mapstructure:"password" json:"password" yaml:"password"` // 密码
}
