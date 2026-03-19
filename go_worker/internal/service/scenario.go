package service

import (
	"fmt"
	"test/python_redis_test/internal/model"
)

func CalculateScenario(in model.ScenarioInput) model.ScenarioResultData {
	current := CalculateRating(model.RatingInput{
		EmployeeID:         in.EmployeeID,
		DealsCount:         in.DealsCount,
		FinancedVolume:     in.FinancedVolume,
		BankSharePercent:   in.BankSharePercent,
		ExtraProductsCount: in.ExtraProductsCount,
	})

	projected := CalculateRating(model.RatingInput{
		EmployeeID:         in.EmployeeID,
		DealsCount:         in.DealsCount + in.DeltaDealsCount,
		FinancedVolume:     in.FinancedVolume + in.DeltaFinancedVolume,
		BankSharePercent:   in.BankSharePercent + in.DeltaBankSharePercent,
		ExtraProductsCount: in.ExtraProductsCount + in.DeltaExtraProductsCount,
	})

	deltaPoints := projected.TotalPoints - current.TotalPoints
	levelChanged := current.Level != projected.Level

	recommendation := buildScenarioRecommendation(current, projected, deltaPoints)

	return model.ScenarioResultData{
		EmployeeID:     in.EmployeeID,
		Current:        current,
		Projected:      projected,
		DeltaPoints:    deltaPoints,
		LevelChanged:   levelChanged,
		Recommendation: recommendation,
	}
}

func buildScenarioRecommendation(current, projected model.RatingResultData, deltaPoints int) string {
	if projected.Level != current.Level {
		return fmt.Sprintf("Сценарий переводит сотрудника с уровня %s на %s", current.Level, projected.Level)
	}

	if deltaPoints > 0 {
		return fmt.Sprintf("Сценарий улучшает рейтинг на %d баллов", deltaPoints)
	}

	if deltaPoints == 0 {
		return "Сценарий не меняет рейтинг"
	}

	return fmt.Sprintf("Сценарий ухудшает рейтинг на %d баллов", -deltaPoints)
}
