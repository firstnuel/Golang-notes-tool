package notes

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getInput(prompt string, expected []string) string {
	reader := bufio.NewReader(os.Stdin)

	if len(expected) == 0 {
		fmt.Printf("%v", prompt)
		text, err := reader.ReadString('\n')
		if err != nil {
			return ""
		}
		return strings.TrimSpace(text)
	}

	for {
		fmt.Printf("\n%v", prompt)
		text, err := reader.ReadString('\n')
		if err != nil {
			return ""
		}
		text = strings.TrimSpace(text)

		if includesString(expected, text) {
			return text
		}
		fmt.Println("Invalid choice!")
	}
}

func includesString(list []string, value string) bool {
	for _, item := range list {
		if item == value {
			return true
		}
	}
	return false
}
