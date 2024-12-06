package main

import (
	"fmt"
	"math"
	"slices"

	"github.com/ehrichpeter/aoc/2024/utils"
)

func main() {
	file := "input.txt"
	lines, err := utils.LoadLinesFromFile(file)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	matrix := utils.ParseLines(lines, "   ")
	columns := utils.RowsToColumns(matrix, 2)

	slices.Sort(columns[0])
	slices.Sort(columns[1])

	var distances []int
	for i, num := range columns[0] {
		distances = append(distances, int(math.Abs(float64(num-columns[1][i]))))
	}

	var result int
	for _, num := range distances {
		result += num
	}

	fmt.Println(result)
}
