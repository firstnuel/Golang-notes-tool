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

// New creates a new Notes instance with a specified collection name.
func New(name string) *Notes {
	return &Notes{
		Collection: name,
		NoteMap:    make(map[string]string),
	}
}

// AddNote prompts the user to enter a note and adds it to the Notes instance.
func (n *Notes) AddNote() *Notes {
	note := getInput("\nEnter the note text: ", []string{})
	// Generates a new key based on the current number of notes.
	i := len(n.NoteMap)
	key := fmt.Sprintf("%03d", i+1)
	n.NoteMap[key] = note
	fmt.Println(Green("Note added successfully!"))
	return n
}

// DeleteNote prompts the user for a note index and deletes the corresponding note if it exists.
func (n *Notes) DeleteNote() *Notes {
	index := getInput(Red("\nEnter the number of note to remove or 0 to cancel: "), []string{})
	if index == "0" {
		fmt.Println(Yellow("Delete operation cancelled"))
		return n
	}
	i := len(n.NoteMap)
	idx, err := strconv.Atoi(index)

	if err != nil {
		fmt.Println(Red("Invalid index format!"))
		return n
	}
	if idx > i {
		fmt.Println(Red("Note not found!"))
		return n
	}
	key := fmt.Sprintf("%03d", idx)
	// Checks if the note exists before attempting to delete.
	if _, exists := n.NoteMap[key]; exists {
		delete(n.NoteMap, key)
		fmt.Println(Green("Note deleted successfully!"))
		return n
	}
	fmt.Println(Red("Note not found, invalid key!"))
	return n
}

// ShowNotes displays all notes in the Notes instance.
func (n *Notes) ShowNotes() {
	if len(n.NoteMap) == 0 {
		fmt.Println(Yellow("No notes to show. Add a note first."))
		return
	}
	fmt.Println("\nNotes:")
	for key, value := range n.NoteMap {
		fmt.Printf("%v - %v\n", key, value)
	}
}

// WriteToFile encrypts and saves notes to a file using the provided encryption key.
func (n *Notes) WriteToFile() error {
	filename := n.Collection + ".txt" // Collection name as filename
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	keyword := getInput(Yellow("Enter key to encrypt file: "), []string{})
	for _, note := range n.NoteMap {
		line := encrypt(note, keyword) + "\n" // Encrypts each line
		if _, err := file.WriteString(line); err != nil {
			return fmt.Errorf("failed to write to file: %w", err)
		}
	}
	return nil
}

// ReadFromFile reads and decrypts notes from a file using the provided decryption key.
// Incorrect key usage will result in improperly decrypted notes.
func (n *Notes) ReadFromFile() error {
	filename := n.Collection + ".txt" // Collection name as filename
	file, err := os.Open(filename)

	if os.IsNotExist(err) {
		// Creates the file if it doesn't exist
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

	keyword := getInput(Yellow("Enter key to decrypt file: "), []string{})
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
