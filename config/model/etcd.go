package model

type Etcd struct {
	Address  string              `mapstructure:"address"json:"address"yaml:"address"`
	Services map[string]*Service `mapstructure:"services" json:"services" yaml:"services"`
	Server   *Server             `mapstructure:"server" json:"server" yaml:"server"`
	Jaeger   *Jaeger             `yaml:"jaeger"`
	Domain   map[string]*Domain  `yaml:"domain"`
}

type Domain struct {
	Name string `yaml:"name"`
}

type Jaeger struct {
	Addr string `yaml:"addr"`
}

type Service struct {
	Name        string   `yaml:"name"`
	LoadBalance bool     `yaml:"loadBalance"`
	Addr        []string `yaml:"addr"`
	Metrics     []string `yaml:"metrics"`
}

type Server struct {
	Port      string `yaml:"port"`
	Version   string `yaml:"version"`
	JwtSecret string `yaml:"jwtSecret"`
	Metrics   string `yaml:"metrics"`
}
