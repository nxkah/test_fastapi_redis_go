package model

type ScenarioInput struct {
	EmployeeID         string  `json:"employee_id"`
	DealsCount         int     `json:"deals_count"`
	FinancedVolume     float64 `json:"financed_volume"`
	BankSharePercent   float64 `json:"bank_share_percent"`
	ExtraProductsCount int     `json:"extra_product_count"`

	DeltaDealsCount         int     `json:"delta_deals_count"`
	DeltaFinancedVolume     float64 `json:"delta_financed_volume"`
	DeltaBankSharePercent   float64 `json:"delta_bank_share_percent"`
	DeltaExtraProductsCount int     `json:"delta_extra_products_count"`
}

type ScenarioResultData struct {
	EmployeeID     string           `json:"employee_id"`
	Current        RatingResultData `json:"current"`
	Projected      RatingResultData `json:"projected"`
	DeltaPoints    int              `json:"delta_points"`
	LevelChanged   bool             `json:"level_changed"`
	Recommendation string           `json:"recommendation"`
}
