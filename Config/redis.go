package config

type Redis struct {
	Addr        string `yaml:"addr"`
	Password    string `yaml:"password"`
	DB          int    `yaml:"db"`
	MaxIdle     int    `yaml:"max-idle"`
	MaxActive   int    `yaml:"max-active"`
	IdleTimeout int    `yaml:"idle-timeout"`
}
