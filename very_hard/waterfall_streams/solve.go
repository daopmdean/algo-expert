package main

func WaterfallStreams(array [][]float64, source int) []float64 {
	rowAbove := array[0]

	rowAbove[source] = -1

	for row := 1; row < len(array); row++ {
		currentRow := array[row]
		for idx := range rowAbove {
			valueAbove := rowAbove[idx]

			hasWaterAbove := valueAbove < 0
			hasBlock := currentRow[idx] == 1

			if !hasWaterAbove {
				continue
			}

			if !hasBlock {
				currentRow[idx] += valueAbove
				continue
			}

			splitWater := valueAbove / 2

			var rightIdx = idx
			for rightIdx+1 < len(rowAbove) {
				rightIdx += 1
				if rowAbove[rightIdx] == 1.0 {
					break
				}
				if currentRow[rightIdx] != 1.0 {
					currentRow[rightIdx] += splitWater
					break
				}
			}

			var leftIdx = idx
			for leftIdx-1 >= 0 {
				leftIdx -= 1
				if rowAbove[leftIdx] == 1.0 {
					break
				}
				if currentRow[leftIdx] != 1.0 {
					currentRow[leftIdx] += splitWater
					break
				}
			}
		}
		rowAbove = currentRow
	}

	finalPercentages := make([]float64, 0, len(rowAbove))
	for _, num := range rowAbove {
		if num == 0 {
			finalPercentages = append(finalPercentages, num)
		} else {
			finalPercentages = append(finalPercentages, num*-100)
		}
	}

	return finalPercentages
}
