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
}