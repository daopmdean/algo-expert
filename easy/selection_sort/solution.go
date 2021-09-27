package easy

func SelectionSort(array []int) []int {
	for i := 0; i < len(array)-1; i++ {
		minValue := array[i]
		minIndex := i
		for j := i + 1; j < len(array); j++ {
			if minValue > array[j] {
				minValue = array[j]
				minIndex = j
			}
		}
		if i != minIndex {
			array[i], array[minIndex] = array[minIndex], array[i]
		}
	}
	return array
}
