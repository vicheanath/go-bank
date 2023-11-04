package models

// create account type enumberation
const (
	AccountTypeSaving = iota + 1
	AccountTypeChecking
)



type Account struct {
	ID      int64   `json:"id"`
	Name    string  `json:"name"`
	Balance float64 `json:"balance"`
	TypeID  int64   `json:"type_id"`
}