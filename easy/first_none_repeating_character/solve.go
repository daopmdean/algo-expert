package main

func FirstNonRepeatingCharacter(str string) int {
	m := map[rune]int{}

	for _, v := range str {
		m[v]++
	}

	for i, v := range str {
		if m[v] == 1 {
			return i
		}
	}

	return -1
}
