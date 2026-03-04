package config

type Config struct {
	Port int `env:"PORT" envDefault:"3000"`

	DatabaseURL string `env:"DATABASE_URL"`

	JWTExpiredIn int    `env:"JWT_EXPIRED_IN" envDefault:"24"` // hour
}
