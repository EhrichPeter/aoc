package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file := "input.txt"
	content, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	stringContent := string(content)
	stringContent = strings.TrimSpace(stringContent)

	mulPattern := `mul\((\d+),(\d+)\)`
	doPattern := `do\(\)`
	dontPattern := `don't\(\)`

	combinedPattern := fmt.Sprintf(`%s.*?%s`, doPattern, dontPattern)

	mulRe, err := regexp.Compile(mulPattern)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return
	}

	dontRe, err := regexp.Compile(dontPattern)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return
	}

	combinedRe, err := regexp.Compile(combinedPattern)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return
	}

	firstDontMatch := dontRe.FindStringIndex(stringContent)
	if firstDontMatch == nil {
		fmt.Println("No 'don't()' found in the content")
		return
	}

	// Split the string at the first occurrence of "don't()"
	beforeFirstDont := stringContent[:firstDontMatch[0]]
	afterFirstDont := stringContent[firstDontMatch[1]:]

	var result int
	startMatches := mulRe.FindAllStringSubmatch(beforeFirstDont, -1)
	for _, match := range startMatches {
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

	allSections := combinedRe.Split(afterFirstDont, -1)
	for _, section := range allSections {
		remainingMatches := mulRe.FindAllStringSubmatch(section, -1)
		for _, match := range remainingMatches {
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
	}

	fmt.Println("Result:", result)
}
