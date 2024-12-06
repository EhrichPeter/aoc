package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var REFERENCE_STRING = "XMAS"
var WINDOW_LENGTH = len(REFERENCE_STRING)

func slideOver(input string) int {
	var hits int
	for i := 0; i < len(input)-WINDOW_LENGTH; i++ {
		if input[i:i+WINDOW_LENGTH] == REFERENCE_STRING {
			hits++
		}
	}
	return hits
}

func traverseHorizontalForward(rows []string) int {
	copyRows := make([]string, len(rows))
	copy(copyRows, rows)

	var hits int
	for _, row := range reverseRows(copyRows) {
		hits += slideOver(row)
	}
	return hits
}

func traverseHorizontalBackward(rows []string) int {
	copyRows := make([]string, len(rows))
	copy(copyRows, rows)

	var hits int
	for _, row := range copyRows {
		hits += slideOver(row)
	}
	return hits
}

func traverseVerticalForward(rows []string) int {
	copyRows := make([]string, len(rows))
	copy(copyRows, rows)

	var hits int
	for i := 0; i < len(copyRows[0]); i++ {
		var column string
		for _, row := range copyRows {
			column += string(row[i])
		}
		hits += slideOver(column)
	}
	return hits
}

func traverseVerticalBackward(rows []string) int {
	copyRows := make([]string, len(rows))
	copy(copyRows, rows)

	var hits int
	for i := 0; i < len(copyRows[0]); i++ {
		var column string
		for _, row := range reverseRows(copyRows) {
			column += string(row[i])
		}
		hits += slideOver(column)
	}
	return hits
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func reverseRows(rows []string) []string {
	var reversedRows []string
	for _, row := range rows {

		reversedRows = append([]string{reverseString(row)}, reversedRows...)
	}
	return reversedRows
}

func main() {

	file, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	rows := strings.Split(string(file), "\n")

	thfHits := traverseHorizontalForward(rows)
	fmt.Println("Hits with traverseHorizontalForward:", thfHits)

	thbHits := traverseHorizontalBackward(rows)
	fmt.Println("Hits with traverseHorizontalBackward:", thbHits)

	tvfHits := traverseVerticalForward(rows)
	fmt.Println("Hits with traverseVerticalForward:", tvfHits)

	tvbHits := traverseVerticalBackward(rows)
	fmt.Println("Hits with traverseVerticalBackward:", tvbHits)
}
