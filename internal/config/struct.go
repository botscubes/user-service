package config

// Config for Redis.
type RedisConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

// Config for DB.
type DBConfig struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Port     string `yaml:"port"`
	Host     string `yaml:"host"`
	DBname   string `yaml:"dbname"`
}

// Config for Server.
type ServerConfig struct {
	Salt   string `yaml:"salt"`
	JWTKey string `yaml:"jwtkey"`
}

// Config for App.
type Config struct {
	Redis  RedisConfig  `yaml:"redis"`
	DB     DBConfig     `yaml:"db"`
	Server ServerConfig `yaml:"server"`
}
