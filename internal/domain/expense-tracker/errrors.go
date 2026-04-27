package expensetracker

import "errors"

var (
	ErrExpenseNotFound    = errors.New("expense not found")
	ErrInvalidDescription = errors.New("invalid task description")
)
