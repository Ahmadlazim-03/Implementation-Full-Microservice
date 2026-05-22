package config

import (
	"os"
	"strconv"
)

type Config struct {
	AppEnv   string
	LogLevel string

	MongoURI       string
	MongoDBAuth    string
	MongoDBPlaces  string
	MongoDBReview  string

	JWTSecret       string
	JWTExpiresHours int

	GatewayPort       string
	AuthServicePort   string
	PlacesServicePort string
	ReviewServicePort string

	AuthServiceURL   string
	PlacesServiceURL string
	ReviewServiceURL string
}

func Load() *Config {
	return &Config{
		AppEnv:   getEnv("APP_ENV", "development"),
		LogLevel: getEnv("LOG_LEVEL", "info"),

		MongoURI:      getEnv("MONGO_URI", "mongodb://localhost:27017"),
		MongoDBAuth:   getEnv("MONGO_DB_AUTH", "kampusmap_auth"),
		MongoDBPlaces: getEnv("MONGO_DB_PLACES", "kampusmap_places"),
		MongoDBReview: getEnv("MONGO_DB_REVIEW", "kampusmap_review"),

		JWTSecret:       getEnv("JWT_SECRET", "dev-secret"),
		JWTExpiresHours: getEnvInt("JWT_EXPIRES_HOURS", 72),

		GatewayPort:       getEnv("GATEWAY_PORT", "8080"),
		AuthServicePort:   getEnv("AUTH_SERVICE_PORT", "8081"),
		PlacesServicePort: getEnv("PLACES_SERVICE_PORT", "8082"),
		ReviewServicePort: getEnv("REVIEW_SERVICE_PORT", "8083"),

		AuthServiceURL:   getEnv("AUTH_SERVICE_URL", "http://localhost:8081"),
		PlacesServiceURL: getEnv("PLACES_SERVICE_URL", "http://localhost:8082"),
		ReviewServiceURL: getEnv("REVIEW_SERVICE_URL", "http://localhost:8083"),
	}
}

func getEnv(key, fallback string) string {
	if v, ok := os.LookupEnv(key); ok && v != "" {
		return v
	}
	return fallback
}

func getEnvInt(key string, fallback int) int {
	if v, ok := os.LookupEnv(key); ok {
		if n, err := strconv.Atoi(v); err == nil {
			return n
		}
	}
	return fallback
}
