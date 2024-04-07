package model

import "fmt"

type Mongo struct {
	Host     string `mapstructure:"host" json:"host" yaml:"host"`
	Port     int    `mapstructure:"port"json:"port" yaml:"port"`
	Username string `mapstructure:"username"json:"username"yaml:"username"`
	Password string `mapstructure:"password"json:"password"yaml:"password"`
	Database string `mapstructure:"database"json:"database"yaml:"database"`
	Collection string `mapstructure:"collection"json:"collection"yaml:"collection"`
	Title string `mapstructure:"title"json:"title"yaml:"title"`
}

func (m *Mongo) Uri() string {
	return fmt.Sprintf("mongodb://%s:%s@%s:%d/%s",
		m.Username, m.Password, m.Host, m.Port, m.Database)
}
