package easy

func BubbleSort(array []int) []int {
	for i := 0; i < len(array)-1; i++ {
		for j := i + 1; j < len(array); j++ {
			if array[i] > array[j] {
				temp := array[i]
				array[i] = array[j]
				array[j] = temp
			}
		}
	}
	return array
}
