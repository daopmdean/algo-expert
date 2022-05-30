package main

import "strconv"

func RunLengthEncoding(str string) string {
	if len(str) == 1 {
		return "1" + str
	}

	result := ""

	count := 0
	char := rune(str[0])
	for i, v := range str {
		if i == 0 {
			count++
			continue
		}

		if v == char {
			if count == 9 {
				count = 1
				result += "9" + string(v)
			} else {
				count++
			}
		} else {
			countStr := strconv.Itoa(count)
			result += countStr + string(char)
			char = v
			count = 1
		}

		if i == len(str)-1 {
			countStr := strconv.Itoa(count)
			result += countStr + string(char)
		}
	}

	return result
}
