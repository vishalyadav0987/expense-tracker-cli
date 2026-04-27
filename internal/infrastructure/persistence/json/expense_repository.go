package json

import (
	"context"
	"encoding/json"
	"os"
	"sync"
	"time"

	domain "github.com/vishalyadav0987/expense-tracker-cli/internal/domain/expense-tracker"
)

type ExpenseRepository struct {
	filePath string
	mu       sync.Mutex
}

func NewExpenseRepository(filePath string) *ExpenseRepository {
	return &ExpenseRepository{
		filePath: filePath,
	}
}

// ---------- helper functions ----------

func (r *ExpenseRepository) loadData() ([]*domain.Expense, error) {
	file, err := os.ReadFile(r.filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return []*domain.Expense{}, nil
		}
		return nil, err
	}

	var expenses []*domain.Expense
	if len(file) == 0 {
		return []*domain.Expense{}, nil
	}

	if err := json.Unmarshal(file, &expenses); err != nil {
		return nil, err
	}

	return expenses, nil
}

func (r *ExpenseRepository) saveAll(expenses []*domain.Expense) error {
	data, err := json.MarshalIndent(expenses, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(r.filePath, data, 0644)
}

// ---------- interface implementation ----------
func (r *ExpenseRepository) Save(ctx context.Context, exp *domain.Expense) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	expense, err := r.loadData()
	if err != nil {
		return err
	}

	expense = append(expense, exp)

	return r.saveAll(expense)

}

func (r *ExpenseRepository) Delete(ctx context.Context, id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	expense, err := r.loadData()
	if err != nil {
		return err
	}

	// Creating Empty expense
	newExpense := make([]*domain.Expense, 0)

	found := false

	for _, e := range expense {
		if e.ID == id {
			found = true
			continue
		}
		newExpense = append(newExpense, e)
	}

	if !found {
		return domain.ErrExpenseNotFound
	}

	return r.saveAll(newExpense)

}

func (r *ExpenseRepository) GetAll(ctx context.Context) ([]*domain.Expense, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	return r.loadData()
}

func (r *ExpenseRepository) GetSummary(ctx context.Context) (int, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	expenses, err := r.loadData()
	if err != nil {
		return 0, err
	}

	totalAmount := 0
	for _, e := range expenses {
		totalAmount += e.Amount
	}

	return totalAmount, nil
}

func (r *ExpenseRepository) GetMonthSummary(ctx context.Context) (int, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	now := time.Now()
	currentMonth := now.Month()
	currentYear := now.Year()

	expenses, err := r.loadData()
	if err != nil {
		return 0, err
	}

	totalAmount := 0

	for _, e := range expenses {
		if e.CreatedAt.Month() == currentMonth &&
			e.CreatedAt.Year() == currentYear {

			totalAmount += e.Amount
		}
	}

	return totalAmount, nil
}
