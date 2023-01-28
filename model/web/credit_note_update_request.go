package web

type CreditNoteUpdateRequest struct {
	// Required Fields
	Period               string `json:"period" validate:"required,period_month"`
	MarketingStructureID string `json:"marketing_structure_id" validate:"required"`
	Status               string `json:"status" validate:"required,status_credit_note"`
}

type CqrsCNSummaryPostUpdateRequest struct {
	// Required Fields
	ValueWithdraw *float32 `json:"value_withdraw" validate:"required"`
}
