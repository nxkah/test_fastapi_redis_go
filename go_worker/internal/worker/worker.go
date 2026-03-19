package worker

import (
	"context"
	"log"
	"test/python_redis_test/internal/model"
	"test/python_redis_test/internal/queue"
	"time"
)

type taskHandler interface {
	Handle(task model.Task) model.Result
}

type Worker struct {
	queue        *queue.RedisClient
	handler      taskHandler
	tasksQueue   string
	resultsQueue string
}

func New(q *queue.RedisClient, h taskHandler, tasksQueue, resultsQueue string) *Worker {
	return &Worker{
		queue:        q,
		handler:      h,
		tasksQueue:   tasksQueue,
		resultsQueue: resultsQueue,
	}
}

func (w *Worker) Run(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		task, err := w.queue.ConsumeTask(ctx, w.tasksQueue, 5*time.Second)
		if err != nil {
			log.Printf("consume task error: %v", err)
			time.Sleep(time.Second)
			continue
		}
		if task == nil {
			continue
		}
		log.Printf("task received: task_id=%s type=%s user_id=%s", task.TaskID, task.Type, task.UserID)

		result := w.handler.Handle(*task)

		if err := w.queue.PushResult(ctx, w.resultsQueue, result); err != nil {
			log.Printf("push result error: task_id=%s err=%v", result.TaskID, err)
			time.Sleep(time.Second)
			continue
		}
		log.Printf("result sent: task_id=%s status=%s", result.TaskID, result.Status)
	}
}
