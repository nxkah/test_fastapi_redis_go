package main

import (
	"context"
	"errors"
	"log"
	"os/signal"
	"syscall"

	"test/python_redis_test/internal/config"
	"test/python_redis_test/internal/httpserver"
	"test/python_redis_test/internal/queue"
	"test/python_redis_test/internal/service"
	"test/python_redis_test/internal/worker"
)

func main() {
	cfg := config.MustLoad()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	redisClient := queue.NewRedisClient(cfg)
	defer func() {
		if err := redisClient.Close(); err != nil {
			log.Printf("redis close error: %v", err)
		}
	}()

	if err := redisClient.WaitForRedis(ctx, cfg.RedisConnectAttempts, cfg.RedisCennectDelay); err != nil {
		log.Fatalf("redis unavaible: %v", err)
	}

	go httpserver.StartHealthServer(cfg.HealthPort)

	reportService := service.NewReportService()

	w := worker.New(
		redisClient,
		reportService,
		cfg.TasksQueue,
		cfg.ResultQueue,
	)

	log.Printf("worker started")

	if err := w.Run(ctx); err != nil && !errors.Is(err, context.Canceled) {
		log.Fatalf("worker stopped with error: %v", err)
	}
	log.Println("worker stopped")
}
