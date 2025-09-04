package config

type Storage struct {
	Engine        string `env:"STORAGE_ENGINE"  required:"true" example:"local"`
	Bucket        string `env:"STORAGE_BUCKET"  required:"true" example:"/legal-counsel/assets"`
	BaseUrl       string `env:"STORAGE_BASE_URL"  required:"true" example:"http://localhost:3000/api"`
	MaxFileSize   int    `env:"STORAGE_MAX_FILE_SIZE"  default:"1024" example:"1024 (KB)"`
	MaxFileUpload int    `env:"STORAGE_MAX_FILE_UPLOAD"  default:"10" example:"10"`
}
