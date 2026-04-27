package expensetracker

import "context"

type ExpenseRepository interface {
	Save(ctx context.Context, expense *Expense) error
	Delete(ctx context.Context, id string) error
	GetAll(ctx context.Context) ([]*Expense, error)
	GetSummary(ctx context.Context) (int, error)
	GetMonthSummary(ctx context.Context) (int, error)
}

// 🧠 Why context?

// Because:

// DB call slow ho sakta hai
// Timeout cancel karna pad sakta hai
// Production safe code
