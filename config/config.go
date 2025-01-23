package config

import "os"

type Config struct {
	AppConfig AppConfig
	HttpConf  ServerConfig
	DBConf    DBConfig
}

type AppConfig struct {
	LogLevel string
}

type ServerConfig struct {
	Host     string
	HttpPort string
}

type DBConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	SSLMode  string
	DBName   string
	LogLevel string
}

func NewConfig() *Config {
	return &Config{
		AppConfig: AppConfig{
			LogLevel: getEnv("APP_LOG_LEVEL", "debug"),
		},

		HttpConf: ServerConfig{
			Host:     getEnv("HTTP_HOST", "0.0.0.0"),
			HttpPort: getEnv("HTTP_PORT", "9000"),
		},

		DBConf: DBConfig{
			Host:     getEnv("DB_HOST", "192.168.0.210"),
			Port:     getEnv("DB_PORT", "5432"),
			Username: getEnv("DB_USER", "test_task"),
			Password: getEnv("DB_PASSWORD", "test_task"),

			SSLMode:  getEnv("DB_SSLMODE", "disable"),
			DBName:   getEnv("DB_NAME", "test_task"),
			LogLevel: getEnv("DB_LOGLEVEL", "info"),
		},
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
