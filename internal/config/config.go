package config

import "os"

type Config struct {
	Port string
}

func Load() Config {
	return Config{
		Port: getEnv("443", "8080"),
	}
}

func getEnv(key, fallback string) string {

	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return fallback
}
