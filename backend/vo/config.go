package vo

type AppConfig struct {
	Version         string `env:"VERSION"`
	DB              string `env:"DB"`
	Port            int    `env:"PORT" env-default:"37892"`
	JwtKey          string `env:"JWT_KEY"`
	UploadDir       string `env:"UPLOAD_DIR"`
	LogLevel        string `env:"LOG_LEVEL" env-default:"INFO"`
	EnableSwagger   bool   `env:"ENABLE_SWAGGER" env-default:"false"`
	EnableSQLOutput bool   `env:"ENABLE_SQL_OUTPUT" env-default:"false"`
}
