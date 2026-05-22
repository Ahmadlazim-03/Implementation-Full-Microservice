package config

import (
	"os"
	"strconv"
)

// Config carries all configuration. Each service uses ONLY the fields
// relevant to its own bounded context — this is intentional per
// Database-per-Service: a service must never read another service's DSN.
type Config struct {
	AppEnv   string
	LogLevel string

	// Auth Service (PostgreSQL)
	PostgresAuthDSN string

	// Places Service (MongoDB)
	MongoPlacesURI string
	MongoDBPlaces  string

	// Review Service (MongoDB)
	MongoReviewURI string
	MongoDBReview  string

	// Redis (shared cache)
	RedisAddr     string
	RedisPassword string
	RedisDB       int

	// JWT
	JWTSecret       string
	JWTExpiresHours int

	// Service Ports
	GatewayPort       string
	AuthServicePort   string
	PlacesServicePort string
	ReviewServicePort string

	// Gateway -> backend service URLs
	AuthServiceURL   string
	PlacesServiceURL string
	ReviewServiceURL string
}

func Load() *Config {
	return &Config{
		AppEnv:   getEnv("APP_ENV", "development"),
		LogLevel: getEnv("LOG_LEVEL", "info"),

		PostgresAuthDSN: getEnv("POSTGRES_AUTH_DSN", "postgres://auth_user:auth_pass@localhost:5433/kampusmap_auth?sslmode=disable"),

		MongoPlacesURI: getEnv("MONGO_PLACES_URI", "mongodb://localhost:27018"),
		MongoDBPlaces:  getEnv("MONGO_DB_PLACES", "kampusmap_places"),

		MongoReviewURI: getEnv("MONGO_REVIEW_URI", "mongodb://localhost:27019"),
		MongoDBReview:  getEnv("MONGO_DB_REVIEW", "kampusmap_review"),

		RedisAddr:     getEnv("REDIS_ADDR", "localhost:6379"),
		RedisPassword: getEnv("REDIS_PASSWORD", ""),
		RedisDB:       getEnvInt("REDIS_DB", 0),

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
