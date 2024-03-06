package model

type JWT struct {
	SigningKey string `mapstructure:"signing-key" json:"signing-key" yaml:"signing-key"` // jwt签名
}
