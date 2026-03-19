package service

import (
	"fmt"
	"strings"
	"test/python_redis_test/internal/model"
)

type levelBenefit struct {
	BonusGain       float64
	MortgageSaving  float64
	CashbackBenefit float64
	DMSValue        float64
}

var benefitByLevel = map[string]levelBenefit{
	"Silver": {
		BonusGain:       0,
		MortgageSaving:  0,
		CashbackBenefit: 10_000,
		DMSValue:        0,
	},
	"Gold": {
		BonusGain:       80_000,
		MortgageSaving:  120_000,
		CashbackBenefit: 35_000,
		DMSValue:        40_000,
	},
	"Black": {
		BonusGain:       180_000,
		MortgageSaving:  250_000,
		CashbackBenefit: 60_000,
		DMSValue:        90_000,
	},
}

func CalculateFinancialEffect(in model.FinancialEffectInput) model.FinancialEffectResultData {
	currentLevel := normalizeLevel(in.CurrentLevel)
	targetLevel := normalizeLevel(in.TargetLevel)

	currentBenefit, ok := benefitByLevel[currentLevel]
	if !ok {
		currentBenefit = benefitByLevel["Silver"]
		currentLevel = "Silver"
	}

	targetBenefit, ok := benefitByLevel[targetLevel]
	if !ok {
		targetBenefit = currentBenefit
		targetLevel = currentLevel
	}

	bonusGain := targetBenefit.BonusGain - currentBenefit.BonusGain
	mortgageSaving := targetBenefit.MortgageSaving - currentBenefit.MortgageSaving
	cashbackBenefit := targetBenefit.CashbackBenefit - currentBenefit.CashbackBenefit
	dmsValue := targetBenefit.DMSValue - currentBenefit.DMSValue

	total := bonusGain + mortgageSaving + cashbackBenefit + dmsValue

	return model.FinancialEffectResultData{
		EmployeeID:         in.EmployeeID,
		CurrentLevel:       currentLevel,
		TargetLevel:        targetLevel,
		BonusGain:          bonusGain,
		MortgageSaving:     mortgageSaving,
		CashbackBenefit:    cashbackBenefit,
		DMSValue:           dmsValue,
		TotalAnnualBenefit: total,
		Recommendation:     buildBenefitRecommendation(currentLevel, targetLevel, total),
	}
}

func normalizeLevel(level string) string {
	level = strings.TrimSpace(strings.ToLower(level))

	switch level {
	case "silver":
		return "Silver"

	case "gold":
		return "Gold"

	case "black":
		return "Black"

	default:
		return "Silver"
	}
}

func buildBenefitRecommendation(currentLevel, targetLevel string, total float64) string {
	if currentLevel == targetLevel {
		return fmt.Sprintf("Сотрудник уже находится на уровне %s, дополнительная выгода не меняется", currentLevel)
	}
	if total <= 0 {
		return fmt.Sprintf("Переход с уровня %s на %s не дает дополнительной годовой выгоды", currentLevel, targetLevel)
	}

	return fmt.Sprintf("Переход с уровня %s на %s дает дополнительную годовую выгоду %.0f ₽", currentLevel, targetLevel, total)
}
