package config

import (
	"fmt"
)

type Server struct {
	Address    string `env:"SERVER_ADDRESS" required:"true" example:"custom-domain.com"`
	Port       string `env:"SERVER_PORT" required:"true" example:"3000"`
	ApiPath    string `env:"SERVER_API_PATH" default:"/api" example:"/api"`
	Timeout    int    `env:"SERVER_TIMEOUT" default:"60" example:"60"`
	SwaggerUrl string `env:"SERVER_SWAGGER_URL" required:"true" example:"custom-domain.com"`
	CorsOrigin string `env:"SERVER_CORS_ORIGIN" default:"*" example:"*"`
	CorsMethod string `env:"SERVER_CORS_METHOD" default:"GET,POST,PUT,PATCH,DELETE,HEAD" example:"GET,POST,PUT,PATCH,DELETE,HEAD"`
}

func (s Server) URL() string {
	return fmt.Sprintf("%s:%s", s.Address, s.Port)
}
