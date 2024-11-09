package notes

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	fmt.Println("Note added successfully!")
	return n
}

func (n *Notes) DeleteNote() *Notes {
	index := getInput("\nEnter the index text (00X): ", []string{})

	i := len(n.NoteMap)
	idx, err := strconv.Atoi(index)

	if err != nil {
		fmt.Println("Problem with index!")
		return n
	}
	if idx > i {
		fmt.Println("Note not found!")
		return n
	}

	key := fmt.Sprintf("%03d", idx)
	delete(n.NoteMap, key)
	fmt.Println("Note deleted successfully!")
	return n
}

func (n *Notes) ShowNotes() {

	if len(n.NoteMap) == 0 {
		fmt.Println("No notes to show, Add note")
		return
	}
	fmt.Println()
	for key, value := range n.NoteMap {
		fmt.Printf("%v - %v\n", key, value)
	}
}

func (n *Notes) WriteToFile() error {
	filename := n.Collection + ".txt"
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	for _, note := range n.NoteMap {
		line := note + "\n"
		if _, err := file.WriteString(line); err != nil {
			return fmt.Errorf("failed to write to file: %w", err)
		}
	}
	return nil
}

func (n *Notes) ReadFromFile() error {

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

	n.NoteMap = make(map[string]string)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		key := fmt.Sprintf("%03d", len(n.NoteMap)+1)
		n.NoteMap[key] = line
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading file: %w", err)
	}
	return nil
}
