package service

import "test/python_redis_test/internal/model"

const (
	silverMax = 49
	goldMin   = 50
	goldMax   = 89
	blackMin  = 90
)

func CalculateRating(in model.RatingInput) model.RatingResultData {
	dealsPoints := in.DealsCount * 2
	volumePoints := int(in.FinancedVolume/1_000_000) * 4
	bankSharePoints := int(in.BankSharePercent/5) * 3
	extraProductPoints := in.ExtraProductsCount * 1

	total := dealsPoints + volumePoints + bankSharePoints + extraProductPoints

	level := "Silver"
	nextLevel := "Gold"
	pointsToNext := 0
	progressPercent := 0.0

	switch {
	case total >= blackMin:
		level = "Black"
		nextLevel = ""
		pointsToNext = 0
		progressPercent = 100

	case total >= goldMin:
		level = "Gold"
		nextLevel = "Black"
		pointsToNext = blackMin - total
		progressPercent = float64(total-goldMin) / float64(blackMin-goldMin) * 100

	default:
		level = "Silver"
		nextLevel = "Gold"
		pointsToNext = goldMin - total
		progressPercent = float64(total) / float64(goldMin) * 100
	}

	if progressPercent < 0 {
		progressPercent = 0
	}

	if progressPercent > 100 {
		progressPercent = 100
	}

	return model.RatingResultData{
		EmployeeID:          in.EmployeeID,
		DealsPoints:         dealsPoints,
		VolumePoints:        volumePoints,
		BankSharePoints:     bankSharePoints,
		ExtraProductsPoints: extraProductPoints,
		TotalPoints:         total,
		Level:               level,
		PointsToNextLevel:   pointsToNext,
		NextLevel:           nextLevel,
		ProgressPercent:     progressPercent,
	}
}
