package cli

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/robertd2000/expense-tracker/interval/service"
)

type Commands struct {
	expenseService service.ExpenseService
}

func NewCommands(expenseService service.ExpenseService) *Commands {
	return &Commands{expenseService: expenseService}
}


func (c *Commands) Add(args []string) {
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	description := addCmd.String("description", "", "Description of the item (e.g., 'Lunch')")
	amount := addCmd.Int("amount", 0, "Amount of the item (e.g., 20)")

	addCmd.Parse(os.Args[2:])

	if *description == "" || *amount == 0 {
		fmt.Println("Error: Both --description and --amount are required.")
		addCmd.Usage()
		return
	}

	_, err := c.expenseService.Add(*description, float64(*amount))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Task with description %s created\n", description)
}
