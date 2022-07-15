package main

func LongestBalancedSubstring(str string) int {
	max := 0
	sl := []int{-1}

	for i := range str {
		if str[i] == '(' {
			sl = append(sl, i)
		} else {
			sl = sl[:len(sl)-1]
			if len(sl) == 0 {
				sl = append(sl, i)
			} else {
				balanceSl := sl[len(sl)-1]
				currentLength := i - balanceSl
				if currentLength > max {
					max = currentLength
				}
			}
		}
	}
	return max
}
