package notes

import "fmt"

func encrypt(text, keyword string) string {

	if keyword == "" {
		fmt.Println("No key entered!")
		return ""
	}

	encrypted := make([]byte, len(text))
	keywordLength := len(keyword)

	for i := 0; i < len(text); i++ {
		char := text[i]
		shift := (keyword[i%keywordLength] - 'a') % 26

		if char >= 'a' && char <= 'z' {
			encrypted[i] = 'a' + (char-'a'+shift)%26
		} else if char >= 'A' && char <= 'Z' {
			encrypted[i] = 'A' + (char-'A'+shift)%26
		} else {
			encrypted[i] = char
		}
	}
	return string(encrypted)
}

func decrypt(text, keyword string) string {

	if keyword == "" {
		fmt.Println("No key entered!")
		return ""
	}

	decrypted := make([]byte, len(text))
	keywordLength := len(keyword)

	for i := 0; i < len(text); i++ {
		char := text[i]
		shift := (keyword[i%keywordLength] - 'a') % 26

		if char >= 'a' && char <= 'z' {
			decrypted[i] = 'a' + (char-'a'-shift+26)%26
		} else if char >= 'A' && char <= 'Z' {
			decrypted[i] = 'A' + (char-'A'-shift+26)%26
		} else {
			decrypted[i] = char
		}
	}
	return string(decrypted)
}
