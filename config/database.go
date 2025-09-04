package config

import "fmt"

type Database struct {
	Host     string `env:"DB_HOST"  required:"true" example:"127.0.0.1"`
	Port     string `env:"DB_PORT"  required:"true" example:"3306"`
	User     string `env:"DB_USER"  required:"true" example:"user"`
	Password string `env:"DB_PASSWORD"  required:"true" example:"password"`
	Name     string `env:"DB_NAME"  required:"true" example:"database_name"`
	Charset  string `env:"DB_CHARSET" default:"utf8mb4" example:"utf8mb4"`
}

func (d Database) DataSourceName() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=UTC&multiStatements=true",
		d.User,
		d.Password,
		d.Host,
		d.Port,
		d.Name,
		d.Charset,
	)
}
