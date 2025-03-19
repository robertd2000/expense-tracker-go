package cli

import "github.com/robertd2000/expense-tracker/interval/service"

type Commands struct {
	taskService service.ExpenseService
}

func NewCommands(expenseService service.ExpenseService) *Commands {
	return &Commands{taskService: expenseService}
}