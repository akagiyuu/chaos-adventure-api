package config

type Config struct {
	Port int `env:"PORT" envDefault:"3000"`

	DatabaseURL string `env:"DATABASE_URL"`

	JWTSecret    string `env:"JWT_SECRET" envDefault:"secret"`
	JWTExpiredIn int    `env:"JWT_EXPIRED_IN" envDefault:"24"` // hour
}
