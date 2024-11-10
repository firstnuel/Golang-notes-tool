package notes

import (
	"fmt"
	"os"
)

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Yellow = "\033[33m"
)

func YellowText(text string) string {
	return Yellow + text + Reset
}

func NotesTool() {
	collectionName := ValidateArgs(os.Args)
	if collectionName == "" {
		return
	}
	notesCollection := New(collectionName)
	notesCollection.ReadFromFile()

	fmt.Print("\nWelcome to the notes tool!\n\n")

	for {
		show_operations()
		opChoice := getInput("Enter choice: ", []string{"1", "2", "3", "4"})

		switch opChoice {
		case "1":
			notesCollection.ShowNotes()

			if len(notesCollection.NoteMap) > 0 {
				decryptChoice := getInput(YellowText("Press 0 to decrypt your notes: "), []string{"0", ""})
				if decryptChoice == "0" {
					notesCollection.ShowDecryptedNotes()
				}
			}

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
