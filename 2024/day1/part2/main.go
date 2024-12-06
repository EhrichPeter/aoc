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

	matrix := utils.ParseLines(lines, "   ")
	columns := utils.RowsToColumns(matrix, 2)

	var occurences = make(map[int]int)
	for _, num := range columns[1] {
		occurences[num]++
	}

	var similarity int
	for _, num := range columns[0] {
		similarity += num * occurences[num]
	}

	fmt.Println(similarity)
}
