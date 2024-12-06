package main

import (
	"fmt"

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

	var similarity []int
	for _, num := range coloumns[0] {
		var counter int
		for _, num2 := range coloumns[1] {
			if num == num2 {
				counter++
			}
		}
		similarity = append(similarity, counter*num)
	}

	var result int
	for _, num := range similarity {
		result += num
	}

	fmt.Println(result)
}
