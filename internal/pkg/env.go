package pkg

import (
	"log"
	"os"
	"strconv"
	"strings"
)

type Environment string

const (
	Local      = "local"
	Production = "production"
)

func GetEnv(key string, defaultValue string) string {
	valueStr := readEnv(key)
	if valueStr == "" {
		return defaultValue
	}

	return valueStr
}

func GetEnvAsInt(key string, defaultValue int) int {
	valueStr := readEnv(key)
	if valueStr == "" {
		return defaultValue
	}

	value, err := strconv.Atoi(valueStr)
	if err != nil {
		log.Fatalf("Cannot parse %s as int: %v", key, err)
	}

	return value
}

func GetEnvAsBool(key string, defaultValue bool) bool {
	valueStr := readEnv(key)
	if valueStr == "" {
		return defaultValue
	}

	value, err := strconv.ParseBool(valueStr)
	if err != nil {
		log.Fatalf("Cannot parse %s as bool: %v", key, err)
	}

	return value
}

func readEnv(key string) string {
	return strings.TrimSpace(os.Getenv(key))
}
