package easy

import "sort"

func TandemBicycle(redShirtSpeeds []int, blueShirtSpeeds []int, fastest bool) int {
	sum := 0
	sort.Ints(redShirtSpeeds)
	sort.Ints(blueShirtSpeeds)

	if fastest {
		blueShirtSpeeds = reverse(blueShirtSpeeds)
		for i, v := range redShirtSpeeds {
			sum += max(v, blueShirtSpeeds[i])
		}
		return sum
	}

	for i, v := range redShirtSpeeds {
		sum += max(v, blueShirtSpeeds[i])
	}
	return sum
}

func max(red, blue int) int {
	if red > blue {
		return red
	}

	return blue
}

func reverse(slice []int) []int {
	for i := 0; i < len(slice)/2; i++ {
		slice[i], slice[len(slice)-i-1] = slice[len(slice)-i-1], slice[i]
	}
	return slice
}
