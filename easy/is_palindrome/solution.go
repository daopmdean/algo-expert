package easy

func IsPalindrome(str string) bool {
	start := 0
	end := len(str) - 1

	for start < end {
		if str[start] != str[end] {
			return false
		}
		start += 1
		end -= 1
	}
	return true
}
