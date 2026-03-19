package model

type RatingInput struct {
	EmployeeID         string  `json:"employee_id"`
	DealsCount         int     `json:"deals_count"`
	FinancedVolume     float64 `json:"financed_volume"`
	BankSharePercent   float64 `json:"bank_share_percent"`
	ExtraProductsCount int     `json:"extra_product_count"`
}

type RatingResultData struct {
	EmployeeID          string  `json:"employee_id"`
	DealsPoints         int     `json:"deals_points"`
	VolumePoints        int     `json:"volume_points"`
	BankSharePoints     int     `json:"bank_share_points"`
	ExtraProductsPoints int     `json:"extra_product_points"`
	TotalPoints         int     `json:"total_points"`
	Level               string  `json:"level"`
	PointsToNextLevel   int     `json:"points_to_next_level"`
	NextLevel           string  `json:"next_level"`
	ProgressPercent     float64 `json:"progress_percent"`
}
