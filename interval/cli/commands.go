package cli

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

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

	fmt.Printf("Task with description %s created\n", *description)
}

func (c *Commands) Delete(args []string) {
	addCmd := flag.NewFlagSet("delete", flag.ExitOnError)
	id := addCmd.String("id", "", "ID of the item")

	addCmd.Parse(os.Args[2:])

	if *id == "" {
		fmt.Println("Error: Both --description and --amount are required.")
		addCmd.Usage()
		return
	}

	i, err := strconv.Atoi(*id)
    if err != nil {
		log.Fatal(err)
    }

	_, err = c.expenseService.Delete(i)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Task with ID %s deleted\n", *id)
}

func (c *Commands) Summary(args []string) {
	addCmd := flag.NewFlagSet("summary", flag.ExitOnError)
	month := addCmd.String("month", "", "ID of the item")

	addCmd.Parse(os.Args[2:])

	if *month == "" {
		summary, err := c.expenseService.GetSummary()

		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Total amount: %f\n", summary)
		
	}

	i, err := strconv.Atoi(*month)
    if err != nil {
		log.Fatal(err)
    }

	summary, err := c.expenseService.GetSummary(i)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Total amount for August: %f\n", summary)
}
