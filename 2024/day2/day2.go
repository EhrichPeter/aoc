package day2

import (
	"math"

	"github.com/ehrichpeter/aoc/2024/utils"
)

func CheckAdjacentLevels(row []int) bool {
	for i := 0; i < len(row)-1; i++ {
		adjacentDistance := int(math.Abs(float64(row[i] - row[i+1])))

		if adjacentDistance < 1 || adjacentDistance > 3 {
			return false
		}
	}
	return true
}

func CheckValidReport(row []int) bool {
	return (utils.RowIsSortedAscending(row) || utils.RowIsSortedDescending(row)) && CheckAdjacentLevels(row)
}
