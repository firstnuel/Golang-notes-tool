package notes

func EncReverse(note string) string {
	var result string
	
	for i := 0; i < len(note); i++ {


		//For uppercase
		if note[i] >= 'A' && note[i] <= 'Z' {
			newLetter := 'Z' - ((note[i]-'A') + 'A')
			result += string(newLetter + 'a')
			continue
		}
		//For lowercase
		if note[i] >= 'a' && note[i] <= 'z' {
			newLetter := 'z' - ((note[i]-'a') + 'a')
			result += string(newLetter + 'A')
			continue
		} else {
		result += string(note[i])
		}
	}
	return result
}