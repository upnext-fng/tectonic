package config

type Auth struct {
	Issuer   string `env:"AUTH_ISSUER" default:"application-server" example:"jwt-token-issuer" description:"the principal that issued the token"`
	Audience string `env:"AUTH_AUDIENCE" default:"application-client" example:"jwt-token-audience" description:"the recipients that the token is intended for"`

	AccessSecret   string `env:"AUTH_ACCESS_SECRET" required:"true" example:"178381614efd2b290b555bbb9" description:"the access token signature secret"`
	AccessTime     int    `env:"AUTH_ACCESS_TIME" default:"3600" example:"3600" description:"the duration of validity of the access token"`
	AccessAttempts int    `env:"AUTH_ACCESS_ATTEMPTS" default:"5" example:"5" description:"the maximum failed authentication attempts before suspension"`

	RefreshSecret string `env:"AUTH_REFRESH_SECRET" required:"true" example:"17838161efd2b290u5432344" description:"the refresh token signature secret"`
	RefreshTime   int    `env:"AUTH_REFRESH_TIME" default:"86400" example:"86400" description:"the duration of validity of the refresh token"`
}
