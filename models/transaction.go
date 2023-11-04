package models

type Transaction struct {
	ID        int64   `json:"id"`
	AccountID int64   `json:"account_id"`
	Amount    float64 `json:"amount"`
	Remarks   string  `json:"remarks"`
	BeforeBalance float64 `json:"before_balance"`
	AfterBalance float64 `json:"after_balance"`
}
