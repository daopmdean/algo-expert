package main

import (
	"fmt"
)

func main() {
	array := SpecialArray{
		5,
		2,
		SpecialArray{7, -1},
		3,
		SpecialArray{
			6,
			SpecialArray{-13, 8},
			4,
		},
	}
	fmt.Println(ProductSum(array))
}

type SpecialArray []interface{}

// Tip: Each element of `array` can either be cast
// to `SpecialArray` or `int`.
func ProductSum(array []interface{}) int {
	sum := 0
	for _, ele := range array {
		val, ok := ele.(int)
		if ok {
			sum += val
		} else {
			sum += SpecialArraySum(ele.(SpecialArray), 2)
		}
	}
	return sum
}

func SpecialArraySum(array []interface{}, level int) int {
	sum := 0
	for _, ele := range array {
		value, ok := ele.(int)
		if ok {
			sum += value
		} else {
			sum += SpecialArraySum(ele.(SpecialArray), level+1)
		}
	}
	return sum * level
}
