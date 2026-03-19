package service

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"test/python_redis_test/internal/model"
)

type ReportService struct{}

func NewReportService() *ReportService {
	return &ReportService{}
}

func (s *ReportService) Handle(task model.Task) model.Result {
	switch task.Type {
	case "generate_report":
		return s.generateReport(task)

	case "recalculate_rating":
		return s.recalculateRating(task)

	case "calculate_scenario":
		return s.calculateScenario(task)

	case "calculate_financial_effect":
		return s.calculateFinancialEffect(task)

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

func (s *ReportService) recalculateRating(task model.Task) model.Result {
	var input model.RatingInput

	if len(task.Payload) == 0 {
		return model.Result{
			TaskID: task.TaskID,
			Status: "error",
			Data:   "empty payload for recalculate_rating",
		}
	}

	if err := json.Unmarshal(task.Payload, &input); err != nil {
		return model.Result{
			TaskID: task.TaskID,
			Status: "error",
			Data:   fmt.Sprintf("invalid rating payload: %v", err),
		}
	}

	resultData := CalculateRating(input)

	raw, err := json.Marshal(resultData)
	if err != nil {
		return model.Result{
			TaskID: task.TaskID,
			Status: "error",
			Data:   fmt.Sprintf("marshal rating result: %v", err),
		}
	}

	return model.Result{
		TaskID: task.TaskID,
		Status: "done",
		Data:   string(raw),
	}
}

func (s *ReportService) calculateScenario(task model.Task) model.Result {
	var input model.ScenarioInput

	if len(task.Payload) == 0 {
		return model.Result{
			TaskID: task.TaskID,
			Status: "error",
			Data:   "empty payload for calculate_scenario",
		}
	}

	if err := json.Unmarshal(task.Payload, &input); err != nil {
		return model.Result{
			TaskID: task.TaskID,
			Status: "error",
			Data:   fmt.Sprintf("invalid scenario payload: %v", err),
		}
	}

	resultData := CalculateScenario(input)

	raw, err := json.Marshal(resultData)
	if err != nil {
		return model.Result{
			TaskID: task.TaskID,
			Status: "error",
			Data:   fmt.Sprintf("marshal scenario result: %v", err),
		}
	}

	return model.Result{
		TaskID: task.TaskID,
		Status: "done",
		Data:   string(raw),
	}
}

func (s *ReportService) calculateFinancialEffect(task model.Task) model.Result {
	var input model.FinancialEffectInput

	if len(task.Payload) == 0 {
		return model.Result{
			TaskID: task.TaskID,
			Status: "error",
			Data:   "empty payload for calculate_financial_effect",
		}
	}

	if err := json.Unmarshal(task.Payload, &input); err != nil {
		return model.Result{
			TaskID: task.TaskID,
			Status: "error",
			Data:   fmt.Sprintf("invalid financial effect payload: %v", err),
		}
	}

	resultData := CalculateFinancialEffect(input)

	raw, err := json.Marshal(resultData)
	if err != nil {
		return model.Result{
			TaskID: task.TaskID,
			Status: "error",
			Data:   fmt.Sprintf("marshal financial effect result: %v", err),
		}
	}

	return model.Result{
		TaskID: task.TaskID,
		Status: "done",
		Data:   string(raw),
	}
}
