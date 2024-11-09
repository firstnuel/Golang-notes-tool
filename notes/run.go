package notes

import (
	"os"

	"github.com/fatih/color"
)

func NotesTool() {
	yellow := color.New(color.FgYellow).SprintFunc()

	collectionName := ValidateArgs(os.Args)
	if collectionName == "" {
		return
	}
	notesCollection := New(collectionName)
	notesCollection.ReadFromFile()

	color.Green("\nWelcome to the notes tool!\n\n")

	for {
		show_operations()
		opChoice := getInput(yellow("Enter choice: "), []string{"1", "2", "3", "4"})

		switch opChoice {
		case "1":
			notesCollection.ShowNotes()
		case "2":
			notesCollection.AddNote()
		case "3":
			notesCollection.DeleteNote()
		case "4":
			notesCollection.WriteToFile()
			Exit()
			return
		}

		if !shouldContinue() {
			notesCollection.WriteToFile()
			Exit()
			return
		}
	}
}
