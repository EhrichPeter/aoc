package main

import (
	"fmt"

	"github.com/ehrichpeter/aoc/2024/day2"
	"github.com/ehrichpeter/aoc/2024/utils"
)

func main() {
	file := "input.txt"
	lines, err := utils.LoadLinesFromFile(file)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	matrix := utils.ParseLines(lines, " ")

	validReportCounter := 0

OuterLoop:
	for _, row := range matrix {
		for i := range row {
			modifiedRow := append([]int{}, row[:i]...)
			modifiedRow = append(modifiedRow, row[i+1:]...)
			if day2.CheckValidReport(modifiedRow) {
				validReportCounter++
				continue OuterLoop
			}
		}
	}

	fmt.Println(validReportCounter)
}
