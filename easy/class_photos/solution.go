package easy

import "sort"

func ClassPhotos(redShirtHeights []int, blueShirtHeights []int) bool {
	sort.Ints(redShirtHeights)
	sort.Ints(blueShirtHeights)
	arrLen := len(redShirtHeights)
	redFront := true

	if redShirtHeights[arrLen-1] == blueShirtHeights[arrLen-1] {
		return false
	} else if redShirtHeights[arrLen-1] > blueShirtHeights[arrLen-1] {
		redFront = false
	}

	for i := 0; i < arrLen; i++ {
		if redFront {
			if redShirtHeights[i] >= blueShirtHeights[i] {
				return false
			}
		} else {
			if redShirtHeights[i] <= blueShirtHeights[i] {
				return false
			}
		}
	}

	return true
}
