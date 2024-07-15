package vo

type SysConfig struct {
	DB     string `env:"DB"`
	Port   int    `env:"PORT" env-default:"37892"`
	JwtKey string `env:"NUXT_JWT_KEY" env-default:"i5e5toXNRbmb7p6UagD7o4sbFVIFvf6X"`
}
