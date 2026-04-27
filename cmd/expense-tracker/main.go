package main

import (
	"fmt"
	"os"

	"github.com/vishalyadav0987/expense-tracker-cli/interfaces/cli"
	appservice "github.com/vishalyadav0987/expense-tracker-cli/internal/infrastructure/expense-tracker"
	"github.com/vishalyadav0987/expense-tracker-cli/internal/infrastructure/persistence/json"
)

func main() {
	fmt.Println("Managing Expense tracker using CLI")
	filePath := "internal/infrastructure/persistence/json/storage.json"
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		file, err := os.Create(filePath)
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}
		file.Close()
	}

	expenseRepo := json.NewExpenseRepository(filePath)
	expenseService := appservice.NewExpenseService(expenseRepo)

	handler := cli.NewHandler(expenseService)

	handler.Run()
}
