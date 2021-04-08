package config

type JWT struct {
	SigningKey  string	`yaml:signing-key`
	ExpiresTime int64
	BufferTime  int64
}
