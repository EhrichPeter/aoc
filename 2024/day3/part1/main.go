package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file := "input.txt"
	content, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	re, err := regexp.Compile(`mul\((\d+),(\d+)\)`)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return
	}

	stringContent := string(content)

	matches := re.FindAllStringSubmatch(stringContent, -1)

	var result int
	for _, match := range matches {
		num1, err := strconv.Atoi(match[1])
		if err != nil {
			fmt.Println("Error converting string to int:", err)
			return
		}

		num2, err := strconv.Atoi(match[2])
		if err != nil {
			fmt.Println("Error converting string to int:", err)
			return
		}

		result += num1 * num2
	}

	fmt.Println(result)

}
