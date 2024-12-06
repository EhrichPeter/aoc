package utils

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func ParseLines(lines []string, splitOperator string) [][]int {
	var parseResult [][]int
	for _, line := range lines {
		stringNumbers := strings.Split(line, splitOperator)
		arr := make([]int, len(stringNumbers))

		for j, element := range stringNumbers {
			num, err := strconv.Atoi(element)
			if err != nil {
				fmt.Println("Error converting string to int:", err)
				continue
			}
			arr[j] = num
		}
		parseResult = append(parseResult, arr)
	}
	return parseResult
}

func RowsToColumns(arr [][]int, n int) [][]int {
	coloumns := make([][]int, n)

	for _, row := range arr {
		for j, num := range row {
			coloumns[j] = append(coloumns[j], num)
		}
	}

	return coloumns
}

func RowIsSortedAscending(row []int) bool {
	return slices.IsSorted(row)
}

func RowIsSortedDescending(row []int) bool {
	var comparison = make([]int, len(row))
	copy(comparison, row)
	slices.Sort(comparison)
	slices.Reverse(comparison)
	return slices.Equal(row, comparison)
}
