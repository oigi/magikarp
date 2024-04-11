package model

type Email struct {
	Host     string `json:"host"yaml:"host"`
	Port     int    `json:"port"yaml:"port"`
	User     string `json:"user"yaml:"user"`
	Password string `json:"password"yaml:"password"`
}
