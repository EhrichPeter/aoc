package main

import (
	"bufio"
	"fmt"
	"os"
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
	
	vat parseResult [][]int
	for _, line := range lines {
		stringNumbers := strings.Split(line, " ")
		arr := make([]int, len(strings.Split(line, " ")))

		for j, element := range stringNumbers {
			num, err := strconv.Atoi(element)
			if err != nil {
				fmt.Println("Error converting string to int:", err)
				continue
			}
			arr = append(arr,num)
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

	finalArray := make([]int, len(strings.Split(lines[0], " ")))

	fmt.Println(finalArray)
}
