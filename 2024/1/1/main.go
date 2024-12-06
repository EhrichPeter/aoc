package main

import (
	"fmt"
	"math"
	"slices"

	"github.com/ehrichpeter/aoc/utils"
)

func main() {
	file := "input.txt"
	lines, err := utils.LoadLinesFromFile(file)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	matrix := utils.ParseLines(lines, "   ")
	coloumns := utils.RowsToColumns(matrix, 2)

	slices.Sort(coloumns[0])
	slices.Sort(coloumns[1])

	var distances []int
	for i, num := range coloumns[0] {
		distances = append(distances, int(math.Abs(float64(num-coloumns[1][i]))))
	}

	var result int
	for _, num := range distances {
		result += num
	}

	fmt.Println(result)
}
