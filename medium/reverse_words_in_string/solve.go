package main

func ReverseWordsInString(str string) string {
	result := ""

	word := ""
	for i, char := range str {
		if string(char) == " " {
			result = " " + word + result
			word = ""
		} else {
			word += string(char)
		}

		if i == len(str)-1 {
			result = word + result
		}
	}

	return result
}
