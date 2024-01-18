package config

// Config for Redis.
type RedisConfig struct {
	Host     string `env:"REDIS_AUTH_HOST,required"`
	Port     string `env:"REDIS_AUTH_PORT,required"`
	Password string `env:"REDIS_AUTH_PASS,required"`
	DB       int    `env:"REDIS_AUTH_DB,required"`
}

// Config for DB.
type DBConfig struct {
	DBname   string `env:"POSTGRES_DB,required"`
	User     string `env:"POSTGRES_USER,required"`
	Password string `env:"POSTGRES_PASSWORD,required"`
	Host     string `env:"POSTGRES_HOST,required"`
	Port     string `env:"POSTGRES_PORT,required"`
}

// Config for Server.
type ServerConfig struct {
	Salt          string `env:"PASSWORD_SALT,required"`
	JWTKey        string `env:"JWT_SECRET_KEY,required"`
	TokenLifeTime int    `env:"TOKEN_LIFE_TIME,required"`
}

// Config for App.
type Config struct {
	Redis  RedisConfig
	DB     DBConfig
	Server ServerConfig
}
