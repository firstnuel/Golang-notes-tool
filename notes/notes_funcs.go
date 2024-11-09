package notes

import (
	"fmt"
	"strings"
)

func show_operations() {
	fmt.Println("Select operation (1/2/3/4):")
	fmt.Print("1. Show notes.\n2. Add a note.\n3. Delete a note.\n4. Exit.\n")
}

func ValidateArgs(args []string) string {
	if len(args) < 2 {
		fmt.Println("Usage: ./notes-tool [TAG]")
		return ""
	} else if len(args) > 2 {
		fmt.Println("Usage: ./notes-tool [TAG]")
		return ""
	}
	return strings.ToLower(args[1:][0])
}

func shouldContinue() bool {
	toContinue := getInput("Perform another operation? (y/n): ", []string{"y", "n", "Y", "N"})
	return strings.ToLower(toContinue) == "y"
}

func Exit() {
	fmt.Println()
	fmt.Println("saving...")
	fmt.Println("exited.... Thanks for using our groups Notes Tool.")
}
