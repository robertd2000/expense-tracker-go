package main

import (
	"github.com/robertd2000/expense-tracker/interval/cli"
	"github.com/robertd2000/expense-tracker/interval/repository"
	"github.com/robertd2000/expense-tracker/interval/service"
)

func main() {
	expenseRepository := repository.NewRepository("db.json")
	expenseService := service.NewExpenseService(expenseRepository)

	cli.CLI(expenseService)
}