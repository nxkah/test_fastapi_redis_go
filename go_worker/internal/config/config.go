package config

import (
	"os"
	"strconv"
	"time"
)

type Config struct {
	RedisAddr            string
	TasksQueue           string
	ResultQueue          string
	HealthPort           string
	RedisConnectAttempts int
	RedisCennectDelay    time.Duration
}

func MustLoad() Config {
	return Config{
		RedisAddr:            getEnv("REDIS_ADDR", "redis:6379"),
		TasksQueue:           getEnv("TASKS_QUEUE", "tasks_queue"),
		ResultQueue:          getEnv("RESULTS_QUEUE", "results_queue"),
		HealthPort:           getEnv("HEALTH_PORT", "8081"),
		RedisConnectAttempts: getEnvAsInt("REDIS_CONNECT_ATTEMPTS", 20),
		RedisCennectDelay:    time.Duration(getEnvAsInt("REDIS_CONNECT_DELAY_SEC", 2)) * time.Second,
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func getEnvAsInt(key string, fallback int) int {
	v := os.Getenv(key)
	if v == "" {
		return fallback
	}

	n, err := strconv.Atoi(v)
	if err != nil {
		return fallback
	}
	return n
}
