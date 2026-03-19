package model

type FinancialEffectInput struct {
	EmployeeID   string `json:"employee_id"`
	CurrentLevel string `json:"current_level"`
	TargetLevel  string `json:"target_level"`
}

type FinancialEffectResultData struct {
	EmployeeID         string  `json:"employee_id"`
	CurrentLevel       string  `json:"curent_level"`
	TargetLevel        string  `json:"target_level"`
	BonusGain          float64 `json:"bonus_gain"`
	MortgageSaving     float64 `json:"mortgage_saving"`
	CashbackBenefit    float64 `json:"cashback_benefit"`
	DMSValue           float64 `json:"dms_value"`
	TotalAnnualBenefit float64 `json:"total_annual_benefit"`
	Recommendation     string  `json:"recommendation"`
}
