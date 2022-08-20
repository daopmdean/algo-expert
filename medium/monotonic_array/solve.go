package main

func IsMonotonic(array []int) bool {
	nonDec := true
	nonInc := true
	for i := 1; i < len(array); i++ {
		if array[i] < array[i-1] {
			nonInc = false
		}
		if array[i] > array[i-1] {
			nonDec = false
		}
	}

	return nonInc || nonDec
}
