package outils

import "os"

// getEnv get key environment variable if exist, otherwise return defaultValue
func GetEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}
