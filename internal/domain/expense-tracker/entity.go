package expensetracker

import (
	"fmt"
	"time"
)

type Expense struct {
	ID          string    `json:"id"`
	Description string    `json:"description"`
	Amount      int       `json:"amount"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func NewExpense(id, description string, amount int) (*Expense, error) {
	if description == "" {
		return nil, ErrInvalidDescription
	}

	now := time.Now()
	fmt.Println("Generated ID:", id)

	return &Expense{
		ID:          id,
		Description: description,
		Amount:      amount,
		CreatedAt:   now,
		UpdatedAt:   now,
	}, nil
}
