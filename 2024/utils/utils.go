package utils

import (
	"fmt"
	"slices"
	"strings"
)

func ParseLines[T any](lines []string, splitOperator string, parseFunc func(string) (T, error)) [][]T {
	var parseResult [][]T
	for _, line := range lines {
		stringElements := strings.Split(line, splitOperator)
		arr := make([]T, len(stringElements))

		for j, element := range stringElements {
			parsedElement, err := parseFunc(element)
			if err != nil {
				fmt.Println("Error parsing element:", err)
				continue
			}
			arr[j] = parsedElement
		}
		parseResult = append(parseResult, arr)
	}
	return parseResult
}

func RowsToColumns(arr [][]int, n int) [][]int {
	columns := make([][]int, n)

	for _, row := range arr {
		for j, num := range row {
			columns[j] = append(columns[j], num)
		}
	}

	return columns
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
