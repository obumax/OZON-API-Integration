package config

import "os"

// Config содержит конфигурационные параметры приложения
type Config struct {
	OzonClientID string
	OzonAPIKey   string
	LogLevel     string
	Port         string
}

// New создает и возвращает новый экземпляр Config, загружая параметры из переменных окружения
func New() *Config {
	return &Config{
		OzonClientID: os.Getenv("OZON_CLIENT_ID"),
		OzonAPIKey:   os.Getenv("OZON_API_KEY"),
		LogLevel:     getEnvOrDefault("LOG_LEVEL", "info"),
		Port:         getEnvOrDefault("PORT", "8080"),
	}
}

// getEnvOrDefault возвращает значение переменной окружения или значение по умолчанию
func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
