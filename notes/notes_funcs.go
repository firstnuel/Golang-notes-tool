package notes

import (
	"fmt"
	"strings"
)

const (
	Reset  = "\033[0m"
	red    = "\033[31m"
	green  = "\033[32m"
	yellow = "\033[33m"
)

// Green formats the text in green.
func Green(text string) string {
	return fmt.Sprintf("%s%s%s", green, text, Reset)
}

// Red formats the text in red.
func Red(text string) string {
	return fmt.Sprintf("%s%s%s", red, text, Reset)
}

// Yellow formats the text in yellow.
func Yellow(text string) string {
	return fmt.Sprintf("%s%s%s", yellow, text, Reset)
}

// Prints out operations
func show_operations() {
	fmt.Println(Green("Select operation (1/2/3/4):"))
	fmt.Print(
		Green("1. Show notes.\n"),
		Green("2. Add a note.\n"),
		Green("3. Delete a note.\n"),
		Green("4. Exit.\n"),
	)
}

// Validates if the program is executed correctly
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
	toContinue := getInput(Yellow("Perform another operation? (y/n): "), []string{"y", "n", "Y", "N"})
	return strings.ToLower(toContinue) == "y"
}

func Exit() {
	fmt.Println(Yellow("\nsaving..."))
	fmt.Println(Green("exited.... \nThanks for using our groups Notes Tool."))
}
