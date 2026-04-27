package appservice

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/vishalyadav0987/expense-tracker-cli/internal/application/expense-tracker/dto"
	domain "github.com/vishalyadav0987/expense-tracker-cli/internal/domain/expense-tracker"
)

type ExpenseService struct {
	repo domain.ExpenseRepository
}

func NewExpenseService(repo domain.ExpenseRepository) *ExpenseService {
	return &ExpenseService{repo: repo}
}

func (s *ExpenseService) AddExpense(ctx context.Context, input dto.AddExpenseInputDTO) error {

	if input.Amount <= 0 {
		return domain.ErrInvalidAmount
	}

	exp := &domain.Expense{
		ID:          uuid.New().String(),
		Description: input.Description,
		Amount:      input.Amount,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	return s.repo.Save(ctx, exp)
}

func (s *ExpenseService) ListExpense(ctx context.Context) ([]*dto.ExpenseDTO, error) {

	expenses, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var result []*dto.ExpenseDTO

	for _, e := range expenses {
		result = append(result, &dto.ExpenseDTO{
			ID:          e.ID,
			Description: e.Description,
			Amount:      e.Amount,
			CreatedAt:   e.CreatedAt,
		})
	}

	return result, nil
}

func (s *ExpenseService) DeleteExpense(ctx context.Context, id string) error {

	if id == "" {
		return domain.ErrInvalidID
	}

	return s.repo.Delete(ctx, id)
}

func (s *ExpenseService) GetSummary(ctx context.Context) (int, error) {
	return s.repo.GetSummary(ctx)
}

func (s *ExpenseService) GetMonthSummary(ctx context.Context) (int, error) {
	return s.repo.GetMonthSummary(ctx)
}
