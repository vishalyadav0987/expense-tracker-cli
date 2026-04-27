package dto

import "time"

type ExpenseDTO struct {
	ID          string    `json:"Id"`
	CreatedAt   time.Time `json:"Date"`
	Description string    `json:"description"`
	Amount      int       `json:"amount"`
}
