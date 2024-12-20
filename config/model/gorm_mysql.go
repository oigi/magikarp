package model

type Mysql struct {
	Path     string `mapstructure:"path" json:"path" yaml:"path"`             // 服务器地址:端口
	Port     string `mapstructure:"port" json:"port" yaml:"port"`             //:端口
	Config   string `mapstructure:"config" json:"config" yaml:"config"`       // 高级配置
	Dbname   string `mapstructure:"db_name" json:"db_name" yaml:"db_name"`    // 数据库名
	Username string `mapstructure:"username" json:"username" yaml:"username"` // 数据库用户名
	Password string `mapstructure:"password" json:"password" yaml:"password"` // 数据库密码
}

func (m *Mysql) Dsn() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Path + ":" + m.Port + ")/" + m.Dbname + "?" + m.Config
}
