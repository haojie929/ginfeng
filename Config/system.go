package config

type System struct {
	AppName   string	`yaml:"app-name"`
	Env       string	`yaml:"env"`
	HTTPAddr  string		`yaml:"http-addr"`
	HTTPSAddr string		`yaml:"https-addr"`
	DbType    string	`yaml:"db-type"`
	DbConnect []string  `yaml:"db-connect"`
	RedisConnect []string  `yaml:"redis-connect"`
}
