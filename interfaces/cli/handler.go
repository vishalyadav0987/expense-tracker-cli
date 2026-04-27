package cli

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"

	app "github.com/vishalyadav0987/expense-tracker-cli/internal/application/expense-tracker"
	"github.com/vishalyadav0987/expense-tracker-cli/internal/application/expense-tracker/dto"
)

type Handler struct {
	service app.ExpenseService
}

func NewHandler(service app.ExpenseService) *Handler {
	return &Handler{service: service}
}

func RenderTasks(expense []*dto.ExpenseDTO) {
	table := tablewriter.NewWriter(os.Stdout)

	table.Header([]string{"ID", "Date", "Description", "Amount"})

	for _, e := range expense {

		table.Append([]string{
			e.ID,
			e.CreatedAt.Format("2006-01-02"),
			e.Description,
			strconv.Itoa(e.Amount),
		})
	}

	table.Render()
}

func (h *Handler) Run() {
	cmd, err := Parse(os.Args)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	ctx := context.Background()

	switch cmd.Name {

	case "add":
		desc, ok := cmd.Args["description"]
		if !ok {
			PrintWarning("description is required")
			return
		}

		amountStr, ok := cmd.Args["amount"]
		if !ok {
			PrintWarning("amount is required")
			return
		}

		amount, err := strconv.Atoi(amountStr)
		if err != nil {
			PrintWarning("invalid amount")
			return
		}

		input := dto.AddExpenseInputDTO{
			Description: desc,
			Amount:      amount,
		}

		if err := h.service.AddExpense(ctx, input); err != nil {
			PrintError(err)
			return
		}

		PrintSuccess("Expense added successfully")

	case "list":
		expenses, err := h.service.ListExpense(ctx)
		if err != nil {
			PrintError(err)
			return
		}

		if len(expenses) == 0 {
			PrintWarning("No expenses found")
			return
		}

		RenderTasks(expenses)

	case "delete":
		id, ok := cmd.Args["id"]
		if !ok {
			PrintInfo("Usage: expense-tracker delete --id <id>")
			return
		}

		if err := h.service.DeleteExpense(ctx, id); err != nil {
			PrintError(err)
			return
		}

		PrintSuccess("Expense deleted successfully")

	case "summary":

		monthStr, hasMonth := cmd.Args["month"]

		if hasMonth {
			// 🔹 Monthly summary
			total, err := h.service.GetMonthSummary(ctx)
			if err != nil {
				PrintError(err)
				return
			}

			PrintSuccess(fmt.Sprintf("Total expenses for month %s: %d", monthStr, total))
			return
		}

		// 🔹 Total summary
		total, err := h.service.GetSummary(ctx)
		if err != nil {
			PrintError(err)
			return
		}

		PrintSuccess(fmt.Sprintf("Total expenses: %d", total))

	default:
		PrintWarning("Unknown command")
	}
}
