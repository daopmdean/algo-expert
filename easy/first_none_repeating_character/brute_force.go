package main

func FirstNonRepeatingCharacter(str string) int {
	for i, v := range str {
		found := false

		for j := 0; j < len(str); j++ {
			if v == rune(str[j]) && i != j {
				found = true
				break
			}

			if (j == len(str)-1) && !found {
				return i
			}
		}

	}

	return -1
}
