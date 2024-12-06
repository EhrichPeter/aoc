package day2

import "math"

func CheckAdjacentLevels(row []int) bool {
	for i := 0; i < len(row)-1; i++ {
		adjacentDistance := int(math.Abs(float64(row[i] - row[i+1])))

		if adjacentDistance < 1 || adjacentDistance > 3 {
			return false
		}
	}
	return true
}
