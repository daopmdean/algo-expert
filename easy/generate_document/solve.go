package main

func GenerateDocument(characters string, document string) bool {
	if len(characters) < len(document) {
		return false
	}

	if len(document) == 0 {
		return true
	}

	for _, v := range document {
		if countChar(characters, v) < countChar(document, v) {
			return false
		}
	}

	return true
}

func countChar(str string, char rune) int {
	count := 0
	for _, v := range str {
		if char == v {
			count++
		}
	}
	return count
}
