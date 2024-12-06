package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func loadLinesFromFile(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
func parseLines(lines []string) [][]int {

	var parseResult [][]int
	for _, line := range lines {
		stringNumbers := strings.Split(line, " ")
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

func rowIsSortedAscending(row []int) bool {
	return slices.IsSorted(row)
}

func rowIsSortedDescending(row []int) bool {
	var comparison = make([]int, len(row))
	copy(comparison, row)
	slices.Sort(comparison)
	slices.Reverse(comparison)
	return slices.Equal(row, comparison)
}

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
	lines, err := loadLinesFromFile(file)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	parsedArray := parseLines(lines)

	validReportCounter := 0
	for _, row := range parsedArray {
		if rowIsSortedAscending(row) && checkAdjacentLevels(row) {
			validReportCounter++
		} else if rowIsSortedDescending(row) && checkAdjacentLevels(row) {
			validReportCounter++
		}
	}

	fmt.Println(validReportCounter)
}
