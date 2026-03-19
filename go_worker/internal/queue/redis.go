package queue

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"test/python_redis_test/internal/config"
	"test/python_redis_test/internal/model"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	client *redis.Client
}

func NewRedisClient(cfg config.Config) *RedisClient {
	rdb := redis.NewClient(&redis.Options{
		Addr:         cfg.RedisAddr,
		DB:           0,
		PoolSize:     20,
		MinIdleConns: 5,
		DialTimeout:  5 * time.Second,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	})

	return &RedisClient{client: rdb}
}

func (r *RedisClient) Close() error {
	return r.client.Close()
}

func (r *RedisClient) WaitForRedis(ctx context.Context, attempts int, delay time.Duration) error {
	var lastErr error

	for i := 1; i <= attempts; i++ {
		if err := r.client.Ping(ctx).Err(); err == nil {
			log.Println("connected to redis")
			return nil
		} else {
			lastErr = err
			log.Printf("waiting for redis: attempt=%d err=%v", i, err)
		}

		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(delay):
		}
	}
	return fmt.Errorf("failed to connect to redis: %w", lastErr)
}

func (r *RedisClient) ConsumeTask(ctx context.Context, queueName string, timeout time.Duration) (*model.Task, error) {
	res, err := r.client.BRPop(ctx, timeout, queueName).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, nil
		}
		return nil, err
	}

	if len(res) != 2 {
		return nil, fmt.Errorf("unexpected brpop response: %+v", res)
	}

	var task model.Task
	if err := json.Unmarshal([]byte(res[1]), &task); err != nil {
		return nil, fmt.Errorf("unmarshal task: %w", err)
	}
	return &task, nil
}

func (r RedisClient) PushResult(ctx context.Context, queueName string, result model.Result) error {
	body, err := json.Marshal(result)
	if err != nil {
		return fmt.Errorf("marshal result: %w", err)
	}

	if err := r.client.LPush(ctx, queueName, body).Err(); err != nil {
		return fmt.Errorf("push result: %w", err)
	}
	return nil
}
