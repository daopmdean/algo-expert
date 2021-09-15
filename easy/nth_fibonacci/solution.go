package nthfibonacci

func GetNthFib(n int) int {
	if n == 1 {
		return 0
	} else if n == 2 {
		return 1
	} else {
		return GetNthFib(n-1) + GetNthFib(n-2)
	}
}

func GetNthFibSolution2(n int) int {
	if n == 1 {
		return 0
	} else if n == 2 {
		return 1
	} else {
		firstNum := 0
		secondNum := 1
		result := 0

		for i := 3; i <= n; i++ {
			result = firstNum + secondNum
			firstNum = secondNum
			secondNum = result
		}
		return result
	}
}
