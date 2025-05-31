package env

import (
	"os"
	"strconv"
)

func GetString(key, fallback string) string {
	value := os.Getenv(key)

	return value
}

func GetInt(key string, fallback int) int {
	value := os.Getenv(key)

	intValue, err := strconv.Atoi(value)
	if err != nil {
		return fallback
	}

	return intValue
}
