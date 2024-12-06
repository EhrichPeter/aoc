package main

import (
	"fmt"

	"github.com/ehrichpeter/aoc/2024/utils"
)

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
