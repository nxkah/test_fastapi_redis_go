package service

import (
	"fmt"
	"strings"
	"test/python_redis_test/internal/model"
	"time"
)

type ReportService struct{}

func NewReportService() *ReportService {
	return &ReportService{}
}

func (s *ReportService) Handle(task model.Task) model.Result {
	switch task.Type {
	case "generate_report":
		return s.generateReport(task)
	default:
		return model.Result{
			TaskID: task.TaskID,
			Status: "error",
			Data:   fmt.Sprintf("unknown task type: %s", task.Type),
		}
	}
}

func (s *ReportService) generateReport(task model.Task) model.Result {
	userID := strings.TrimSpace(task.UserID)
	if userID == "" {
		userID = "unknown"
	}

	result := fmt.Sprintf("report generated for user %s at %s", userID, time.Now().Format(time.RFC3339))

	return model.Result{
		TaskID: task.TaskID,
		Status: "done",
		Data:   result,
	}
}
