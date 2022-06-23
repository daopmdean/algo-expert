package main

func FourNumberSum(array []int, target int) [][]int {
	result := [][]int{}
	if len(array) < 4 {
		return result
	}

	for i := 0; i < len(array)-3; i++ {
		for j := i + 1; j < len(array)-2; j++ {
			for k := j + 1; k < len(array)-1; k++ {
				for m := k + 1; m < len(array); m++ {
					if array[i]+array[j]+array[k]+array[m] == target {
						result = append(result, []int{array[i], array[j], array[k], array[m]})
					}
				}
			}
		}
	}

	return result
}
