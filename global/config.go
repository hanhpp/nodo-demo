package global

import (
	"os"
	"strconv"
)

func getEnv(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return fallback
}

func getEnvUint64(key string, fallback uint64) uint64 {
	if val, ok := os.LookupEnv(key); ok {
		valInt, err := strconv.ParseUint(val, 10, 64)
		if err != nil {
			return fallback
		}
		return valInt
	}
	return fallback
}

func getEnvInt(key string, fallback int) int {
	if val, ok := os.LookupEnv(key); ok {
		valInt, err := strconv.Atoi(val)
		if err != nil {
			return fallback
		}
		return valInt
	}
	return fallback
}

func GetEnvConfig() *VecConfig {
	if Config == nil {
		FetchEnvs()
	}

	return Config
}
