package model

import "fmt"

type Mongo struct {
	Host     string `mapstructure:"host" json:"host" yaml:"host"`
	Port     int    `mapstructure:"port"json:"port" yaml:"port"`
	Username string `mapstructure:"username"json:"username"yaml:"username"`
	Password string `mapstructure:"password"json:"password"yaml:"password"`
}

func (m *Mongo) Uri() string {
	return fmt.Sprintf("mongodb://%s:%s@%s:%d",
		m.Username, m.Password, m.Host, m.Port)
}
