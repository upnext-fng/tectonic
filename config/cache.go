package config

type Cache struct {
	Host     string `env:"CACHE_HOST" default:"127.0.0.1" example:"127.0.0.1"`
	Port     int    `env:"CACHE_PORT"  default:"6379" example:"6379"`
	User     string `env:"CACHE_USER"  default:"default" example:"user"`
	Password string `env:"CACHE_PASSWORD"  default:"password" example:"password"`
	TTL      int    `env:"CACHE_TTL"  required:"true" example:"86400"`
}
