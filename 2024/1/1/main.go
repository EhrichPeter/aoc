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
	var column1, column2 []int
	for _, line := range lines {
		stringNumbers := strings.Split(line, "   ")

		for i, element := range stringNumbers {
			num, err := strconv.Atoi(element)
			if err != nil {
				fmt.Println("Error converting string to int:", err)
				continue
			}
			if i == 0 {
				column1 = append(column1, num)
			} else if i == 1 {
				column2 = append(column2, num)
			}
		}
	}
	return [][]int{column1, column2}
}

func main() {
	file := "input.txt"
	lines, err := loadLinesFromFile(file)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	columns := parseLines(lines)
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
