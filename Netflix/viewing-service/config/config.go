package config

import (
	"log"
	"os"
	"strconv"
	"time"
)

// Config holds all configuration for the Viewing Service
type Config struct {
	// Server configuration
	ServerPort string
	
	// Redis configuration
	RedisAddr     string
	RedisPassword string
	RedisDB       int
	
	// PostgreSQL configuration (optional)
	PostgresEnabled bool
	PostgresHost    string
	PostgresPort    int
	PostgresUser    string
	PostgresPassword string
	PostgresDB      string
	
	// JWT configuration
	JWTSecret        string
	JWTExpiration    time.Duration
	
	// Service URLs
	IdentityServiceURL string
	ContentServiceURL  string
	
	// Concurrency limits
	MaxConcurrentStreams int
	
	// Session configuration
	SessionTimeout       time.Duration
	HeartbeatInterval    time.Duration
}

// LoadConfig loads configuration from environment variables
func LoadConfig() *Config {
	cfg := &Config{
		ServerPort:           getEnv("SERVER_PORT", "8080"),
		RedisAddr:            getEnv("REDIS_ADDR", "localhost:6379"),
		RedisPassword:        getEnv("REDIS_PASSWORD", ""),
		RedisDB:              getEnvAsInt("REDIS_DB", 0),
		PostgresEnabled:      getEnvAsBool("POSTGRES_ENABLED", false),
		PostgresHost:         getEnv("POSTGRES_HOST", "localhost"),
		PostgresPort:         getEnvAsInt("POSTGRES_PORT", 5432),
		PostgresUser:         getEnv("POSTGRES_USER", "postgres"),
		PostgresPassword:     getEnv("POSTGRES_PASSWORD", ""),
		PostgresDB:           getEnv("POSTGRES_DB", "viewing_service"),
		JWTSecret:            getEnv("JWT_SECRET", "default-secret-change-in-production"),
		JWTExpiration:        getEnvAsDuration("JWT_EXPIRATION", "1h"),
		IdentityServiceURL:   getEnv("IDENTITY_SERVICE_URL", "http://localhost:8081"),
		ContentServiceURL:    getEnv("CONTENT_SERVICE_URL", "http://localhost:8082"),
		MaxConcurrentStreams: getEnvAsInt("MAX_CONCURRENT_STREAMS", 4),
		SessionTimeout:       getEnvAsDuration("SESSION_TIMEOUT", "24h"),
		HeartbeatInterval:    getEnvAsDuration("HEARTBEAT_INTERVAL", "30s"),
	}
	
	log.Printf("Config loaded: Server will run on port %s", cfg.ServerPort)
	return cfg
}

// Helper functions to get environment variables with defaults
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	valueStr := os.Getenv(key)
	if valueStr == "" {
		return defaultValue
	}
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		log.Printf("Warning: Invalid value for %s, using default: %d", key, defaultValue)
		return defaultValue
	}
	return value
}

func getEnvAsBool(key string, defaultValue bool) bool {
	valueStr := os.Getenv(key)
	if valueStr == "" {
		return defaultValue
	}
	value, err := strconv.ParseBool(valueStr)
	if err != nil {
		log.Printf("Warning: Invalid value for %s, using default: %t", key, defaultValue)
		return defaultValue
	}
	return value
}

func getEnvAsDuration(key string, defaultValue string) time.Duration {
	valueStr := os.Getenv(key)
	if valueStr == "" {
		valueStr = defaultValue
	}
	duration, err := time.ParseDuration(valueStr)
	if err != nil {
		log.Printf("Warning: Invalid duration for %s, using default: %s", key, defaultValue)
		duration, _ = time.ParseDuration(defaultValue)
	}
	return duration
}
