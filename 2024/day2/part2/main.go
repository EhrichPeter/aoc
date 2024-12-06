package main

import (
	"fmt"
	"math"

	"github.com/ehrichpeter/aoc/2024/day2"
	"github.com/ehrichpeter/aoc/2024/utils"
)

func checkAdjacentLevels(row []int) bool {
	for i := 0; i < len(row)-1; i++ {
		adjacentDistance := int(math.Abs(float64(row[i] - row[i+1])))

		if adjacentDistance < 1 || adjacentDistance > 3 {
			return false
		}
	}
	return true
}

func main() {
	file := "input.txt"
	lines, err := utils.LoadLinesFromFile(file)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	parsedArray := utils.ParseLines(lines, " ")

	validReportCounter := 0
	for _, row := range parsedArray {
		if day2.CheckAdjacentLevels(row) && (utils.RowIsSortedAscending(row) || utils.RowIsSortedDescending(row)) {
			validReportCounter++
		}
	}

	fmt.Println(validReportCounter)
}
