package app

import (
	"context"

	"github.com/vishalyadav0987/expense-tracker-cli/internal/application/expense-tracker/dto"
)

type ExpenseService interface {
	AddExpense(ctx context.Context, input dto.AddExpenseInputDTO) error
	ListExpense(ctx context.Context) ([]*dto.AddExpenseInputDTO, error)
	DeleteExpense(ctx context.Context, id string) error
	GetSummary(ctx context.Context) (int, error)
	GetMonthSummary(ctx context.Context) (int, error)
}
