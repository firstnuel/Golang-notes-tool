package notes

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/fatih/color"
)

type Notes struct {
	Collection string
	NoteMap    map[string]string
}

func New(name string) *Notes {
	return &Notes{
		Collection: name,
		NoteMap:    make(map[string]string),
	}
}

func (n *Notes) AddNote() *Notes {
	note := getInput("\nEnter the note text: ", []string{})

	i := len(n.NoteMap)
	key := fmt.Sprintf("%03d", i+1)
	n.NoteMap[key] = note
	color.Green("Note added successfully!")
	return n
}

func (n *Notes) DeleteNote() *Notes {
	yellow := color.New(color.FgYellow).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()

	index := getInput(red("\nEnter the number of note to remove or 0 to cancel: "), []string{})
	fmt.Println(yellow("Delete operation cancelled"))
	if index == "0" {
		return n
	}

	i := len(n.NoteMap)
	idx, err := strconv.Atoi(index)

	if err != nil {
		color.Red("Problem with index!")
		return n
	}
	if idx > i {
		color.Red("Note not found!")
		return n
	}

	key := fmt.Sprintf("%03d", idx)
	delete(n.NoteMap, key)
	color.Green("Note deleted successfully!")
	return n
}

func (n *Notes) ShowNotes() {

	if len(n.NoteMap) == 0 {
		color.Yellow("No notes to show, Add note")
		return
	}
	fmt.Println("\nNotes:")
	for key, value := range n.NoteMap {
		fmt.Printf("%v - %v\n", key, value)
	}
}

func (n *Notes) WriteToFile() error {
	yellow := color.New(color.FgYellow).SprintFunc()

	filename := n.Collection + ".txt"
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	keyword := getInput(yellow("Entey key to encrypt file: "), []string{})
	for _, note := range n.NoteMap {
		line := encrypt(note, keyword) + "\n"
		if _, err := file.WriteString(line); err != nil {
			return fmt.Errorf("failed to write to file: %w", err)
		}
	}
	return nil
}

func (n *Notes) ReadFromFile() error {
	yellow := color.New(color.FgYellow).SprintFunc()

	filename := n.Collection + ".txt"
	file, err := os.Open(filename)

	if os.IsNotExist(err) {
		file, err = os.Create(filename)
		if err != nil {
			return fmt.Errorf("failed to create file: %w", err)
		}
		file.Close()
		return nil

	} else if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	keyword := getInput(yellow("Entey key to decrypt file: "), []string{})
	n.NoteMap = make(map[string]string)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		key := fmt.Sprintf("%03d", len(n.NoteMap)+1)
		n.NoteMap[key] = decrypt(line, keyword)
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading file: %w", err)
	}
	return nil
}
