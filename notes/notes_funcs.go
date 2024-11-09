package notes

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

func show_operations() {
	green := color.New(color.FgGreen).SprintFunc()

	fmt.Println(green("Select operation (1/2/3/4):"))
	fmt.Print(
		green("1. Show notes.\n"),
		green("2. Add a note.\n"),
		green("3. Delete a note.\n"),
		green("4. Exit.\n"),
	)
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
	yellow := color.New(color.FgYellow).SprintFunc()
	toContinue := getInput(yellow("Perform another operation? (y/n): "), []string{"y", "n", "Y", "N"})
	return strings.ToLower(toContinue) == "y"
}

func Exit() {
	yellow := color.New(color.FgYellow).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	fmt.Println(yellow("\nsaving..."))
	fmt.Println(green("exited.... \nThanks for using our groups Notes Tool."))
}
