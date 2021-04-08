package config

type Mysql struct {
	DriverName   string		`yaml:"driver_name"`
	Path         string		`yaml:"path"`
	Config       string		`yaml:"config"`
	Dbname       string		`yaml:"db-name"`
	Username     string		`yaml:"username"`
	Password     string		`yaml:"password"`
	MaxIdleConns int		`yaml:"max-idle-conns"`
	MaxOpenConns int		`yaml:"max-open-conns"`
	SingularTable bool		`yaml:"singular-table"`
}

