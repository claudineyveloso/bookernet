package configs

import (
	"os"
	"strconv"
)

type Config struct {
	PublicHost             string
	Port                   string
	DBUser                 string
	DBPassword             string
	DBName                 string
	JWTSecret              string
	JWTExpirationInSeconds int64
}

var Envs = InitConfig()

func InitConfig() Config {
	return Config{
		PublicHost:             getEnv("PUBLIC_HOST", "localhost"),
		Port:                   getEnv("PORT", "5432"),
		DBUser:                 getEnv("DB_USER", "user_bookernet_development"),
		DBPassword:             getEnv("DB_PASSWORD", "pwd_bookernet_development"),
		DBName:                 getEnv("DB_NAME", "bookernet_development"),
		JWTSecret:              getEnv("JWT_SECRET", "not-so-secret-now-is-it?"),
		JWTExpirationInSeconds: getEnvAsInt("JWT_EXPIRATION_IN_SECONDS", 3600*24*7),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getEnvAsInt(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fallback
		}
		return i
	}
	return fallback
}
