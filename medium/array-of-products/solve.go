package main

func ArrayOfProducts(array []int) []int {
	products := make([]int, len(array))

	left := 1
	for i := range array {
		products[i] = left
		left *= array[i]
	}

	right := 1
	for i := len(array) - 1; i >= 0; i-- {
		products[i] *= right
		right *= array[i]
	}

	return products
}
