package main

// abcdefghijklmnopqrstuvwxyz
func CaesarCipherEncryptor(str string, key int) string {
	result := []rune{}

	for _, v := range str {
		shiftedLocation := findShiftedLocation(int(v), key)
		shiftedLetter := rune(shiftedLocation)
		result = append(result, shiftedLetter)
	}

	return string(result)
}

func findShiftedLocation(start, key int) int {
	// range 97 122
	if start+key <= 122 {
		return start + key
	}

	mod := key % 26
	if start+mod <= 122 {
		return start + mod
	}

	return 96 + (start + mod - 122)
}
