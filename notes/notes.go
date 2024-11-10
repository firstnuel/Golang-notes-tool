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

// New creates a new Notes instance
func New(name string) *Notes {
	return &Notes{
		Collection: name,
		NoteMap:    make(map[string]string),
	}
}

// AddNote prompts the user for a note, encrypts it, and adds it to NoteMap
func (n *Notes) AddNote() *Notes {
	note := getInput("\nEnter the note text: ", []string{})
	encryptedNote := EncReverse(note)

	i := len(n.NoteMap)
	key := fmt.Sprintf("%03d", i+1)
	n.NoteMap[key] = encryptedNote
	fmt.Println("Note added and encrypted successfully!")
	return n
}

// DeleteNote prompts the user for an index and deletes the corresponding note
func (n *Notes) DeleteNote() *Notes {
	index := getInput("\nEnter the index text (00X): ", []string{})

	i := len(n.NoteMap)
	idx, err := strconv.Atoi(index)

	if err != nil || idx > i || idx < 1 {
		fmt.Println("Problem with index!")
		return n
	}

	key := fmt.Sprintf("%03d", idx)
	delete(n.NoteMap, key)
	fmt.Println("Note deleted successfully!")
	return n
}

// ShowNotes displays the encrypted notes
func (n *Notes) ShowNotes() {
	if len(n.NoteMap) == 0 {
		fmt.Println("No notes to show, Add note")
		return
	}
	fmt.Println("\nEncrypted Notes:")
	for key, value := range n.NoteMap {
		fmt.Printf("%v - %v\n", key, value)
	}
}

// ShowDecryptedNotes displays decrypted notes by reversing the encryption
func (n *Notes) ShowDecryptedNotes() {
	if len(n.NoteMap) == 0 {
		fmt.Println("No notes to show, Add note")
		return
	}
	fmt.Println("\nDecrypted Notes:")
	for key, value := range n.NoteMap {
		decryptedNote := EncReverse(value) // Decrypt the note
		fmt.Printf("%v - %v\n", key, decryptedNote)
	}
}

// WriteToFile saves encrypted notes to a file
func (n *Notes) WriteToFile() error {
	filename := n.Collection + ".txt"
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	for _, note := range n.NoteMap {
		line := note + "\n" // Already encrypted
		if _, err := file.WriteString(line); err != nil {
			return fmt.Errorf("failed to write to file: %w", err)
		}
	}
	return nil
}

// ReadFromFile reads notes from a file, assuming they are encrypted
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
		n.NoteMap[key] = line // Load as encrypted
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading file: %w", err)
	}
	return nil
}
