package cli

import (
	"flag"
	"fmt"
	"log"

	"github.com/robertd2000/expense-tracker/interval/service"
)

type Commands struct {
	expenseService service.ExpenseService
}

func NewCommands(expenseService service.ExpenseService) *Commands {
	return &Commands{expenseService: expenseService}
}


func (c *Commands) Add(args []string) {
	description := flag.String("description", "", "Description of the item (e.g., 'Lunch')")

	// Определяем числовой флаг --amount
	amount := flag.Int("amount", 0, "Amount of the item (e.g., 20)")

	// Парсим аргументы командной строки
	flag.Parse()

	// Проверяем, были ли переданы обязательные флаги
	if *description == "" || *amount == 0 {
		fmt.Println("Error: Both --description and --amount are required.")
		flag.Usage() // Показываем справку по использованию
		return
	}
	_, err := c.expenseService.Add(*description, float64(*amount))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Task with description %s created\n", description)
}
