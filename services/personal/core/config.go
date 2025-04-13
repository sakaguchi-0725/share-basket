package core

import "os"

type Config struct {
	Env             string
	Port            string
	AuthService     string
	PersonalService string
}

func LoadConfig() *Config {
	return &Config{
		Env:             getEnv("APP_ENV", "dev"),
		Port:            getEnv("PORT", "8080"),
		AuthService:     getEnv("AUTH_SERVICE", "http://localhost:8081"),
		PersonalService: getEnv("PERSONAL_SERVICE", "http://localhost:8082"),
	}
}

func getEnv(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return fallback
}
