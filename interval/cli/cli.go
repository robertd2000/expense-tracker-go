package cli

import (
	"fmt"
	"os"
)

func CLI() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("No arguments provided")
		return
	}

	switch command := args[0]; command {
	case "add":
		commands.Add(args)
	case "update":
		commands.Update(args)
	case "delete":
		commands.Delete(args)
	case "list":
		commands.List(args)
	case "summary":
		commands.Summary(args)
	default:
		fmt.Println("Invalid command")
	}
}