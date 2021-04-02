package shell

import (
	"os"
	"strconv"
	"time"
)

// GetBoolEnv parses a boolean environment variable by the given key, if env is empty, returns the given fallback value
func GetBoolEnv(key string, fallback bool) bool {
	strValue := os.Getenv(key)
	if len(strValue) == 0 {
		return fallback
	}

	value, err := strconv.ParseBool(strValue)
	if err != nil {
		return fallback
	}
	return value
}

// GetDurationEnv parses a duration environment variable by the given key, if env is empty, returns the given fallback value
func GetDurationEnv(key string, fallback time.Duration) time.Duration {
	strValue := os.Getenv(key)
	if len(strValue) == 0 {
		return fallback
	}

	value, err := time.ParseDuration(strValue)
	if err != nil {
		return fallback
	}
	return value
}

// GetStringEnv returns an environment variable by the given key, if env is empty, returns the given fallback value
func GetStringEnv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
