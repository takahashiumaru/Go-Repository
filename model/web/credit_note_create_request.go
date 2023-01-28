package web

type CreditNoteCreateRequest struct {
	// Required Fields
	Period             string  `json:"period" validate:"required,period_month"`
	DiscountProposalID string  `json:"discount_proposal_id" validate:"required"`
	CustomerID         string  `json:"customer_id" validate:"required"`
	OutletID           string  `json:"outlet_id" validate:"required"`
	ProductID          string  `json:"product_id" validate:"required"`
	DistributorID      string  `json:"distributor_id" validate:"required"`
	InvoiceNo          string  `json:"invoice_no" validate:"required"`
	InvoiceDate        string  `json:"invoice_date" validate:"required,datetime=2006-01-02"`
	Qty                float32 `json:"qty" validate:"required"`
	Price              float32 `json:"price" validate:"required"`
}
