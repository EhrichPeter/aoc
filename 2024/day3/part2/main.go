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
	stringContent = strings.ReplaceAll(stringContent, "\r\n", "")
	stringContent = strings.ReplaceAll(stringContent, "\n", "")

	mulPattern := `mul\((\d+),(\d+)\)`
	doPattern := `do\(\)`
	dontPattern := `don't\(\)`

	combinedPattern := fmt.Sprintf(`%s|%s|%s`, doPattern, dontPattern, mulPattern)

	re, err := regexp.Compile(combinedPattern)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return
	}

	mutlRe, err := regexp.Compile(mulPattern)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return
	}

	doRe, err := regexp.Compile(doPattern)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return
	}

	dontRe, err := regexp.Compile(dontPattern)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return
	}

	matches := re.FindAllStringSubmatch(stringContent, -1)

	var result int
	multEnabled := true

	for _, match := range matches {
		if doRe.MatchString(match[0]) {
			multEnabled = true
			fmt.Println("do")
		} else if dontRe.MatchString(match[0]) {
			multEnabled = false
			fmt.Println("dont")
		} else if mutlRe.MatchString(match[0]) && multEnabled {
			fmt.Println("mul", match[1], match[2])
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

	fmt.Println(result)
}
