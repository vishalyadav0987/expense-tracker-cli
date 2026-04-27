package expensetracker

import "errors"

var (
	ErrExpenseNotFound    = errors.New("expense not found")
	ErrInvalidDescription = errors.New("invalid task description")
	ErrInvalidID          = errors.New("invalid id")
	ErrInvalidAmount      = errors.New("invalid amount")
)
