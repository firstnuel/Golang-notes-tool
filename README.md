# Notes Tool
The Notes Tool is a command-line interface (CLI) application written in Go for managing personal notes. It provides a simple, interactive way to add, view, delete, and encrypt notes before saving them to a file. With color-coded console output and a user-friendly menu, Notes Tool makes it easy to keep track of important information securely.

## Features
- **Add a Note:** Create a text note with an auto-generated key.
- **Show Notes:** Display all saved notes with unique keys.
- **Delete a Note:** Remove a note by specifying its key.
- **File Encryption and Decryption:** Encrypts notes with a keyword before saving to a file and decrypts them when loading.
- **Color-coded Console Output:** Enhances readability of console messages with color coding.

### Prerequisites
- **Go:** Ensure *Go* is installed on your system.

## Installation
1. **Clone the Repository:**

    ```bash
    git clone https://gitea.koodsisu.fi/emmanuelikwunna1/notes.git
    cd notes-tool
    ````

2. **Run the Application:** Build it and execute the binary using:

    ````bash
    go build
    ./notes-tool [TAG]
    ````

## Usage Guide
1. **Starting the Tool**

    ````bash
    ./notes-tool demo
    ````
    Replace "demo" with your chosen tag. This tag is used as the file name to store notes.

2. **Interactive Menu** The tool displays an interactive menu where you can choose operations by entering the corresponding number:

    - `1` Show Notes: View all saved notes.
    - `2` Add a Note: Create a new note by entering note text.
    - `3` Delete a Note: Delete an existing note by its number.
    - `4` Exit: Save notes to file and exit.

## Example Interaction

    % ./notes-tool demo

    Welcome to the notes tool!

    Select operation (1/2/3/4):
    1. Show notes.
    2. Add a note.
    3. Delete a note.
    4. Exit.

    Enter choice: 1   
    No notes to show. Add a note first.

    Perform another operation? (y/n): y
    Select operation (1/2/3/4):
    1. Show notes.
    2. Add a note.
    3. Delete a note.
    4. Exit.

    Enter choice: 2

    Enter the note text: First test note
    Note added successfully!

    Perform another operation? (y/n): y
    Select operation (1/2/3/4):
    1. Show notes.
    2. Add a note.
    3. Delete a note.
    4. Exit.

    Enter choice: 1

    Notes:
    001 - First test note

    Perform another operation? (y/n): n
    Enter key to encrypt file: love

    saving...
    exited.... 
    Thanks for using our groups Notes Tool.


## Explanation of How Data is Stored
The Notes Tool stores all notes in memory using a map (key-value pair). Each note is assigned a unique key, formatted as a three-digit number (e.g., 001, 002, etc.), and the note text is stored as the value.

When you exit the notes tool, the tool writes the notes to a file named after the collection (e.g., demo.txt). The notes are encrypted using the Vigenère cipher before being written to the file. This means that even if someone gains access to the file, they won't be able to read the notes without the correct decryption key.

When you run the tool, it will first check to load the notes from the file. If the file does not exist, it will create a new one with the collection name. If the file exists, the program will ask for the decryption key, and the notes will be decrypted using the same Vigenère cipher algorithm.


## File Encryption and Decryption
The Notes Tool offers simple encryption and decryption using a keyword to secure notes stored in a file.

### Encryption Method
- **Vigenère Cipher:** This tool applies a modified version of the Vigenère cipher for encryption. When exiting and saving notes, the tool prompts for a keyword that is used to shift each character in the notes by a value derived from the keyword. The Vigenère cipher technique involves shifting letters based on corresponding characters in the keyword, providing lightweight encryption that makes it challenging to read the notes without the correct keyword.

- **Encryption Process:** Each character in the note is shifted by a different amount, determined by the characters in the keyword, creating a pattern. Lowercase and uppercase letters are preserved, while non-alphabet characters remain unchanged.

### Decryption
- The decryption process uses the same keyword to reverse the shifts applied during encryption. 
**Note:** *If an incorrect keyword is provided, the tool will attempt decryption, resulting in unreadable or garbled text.*
This encryption mechanism ensures that notes are secured on disk and can only be accessed by those with the correct keyword.

