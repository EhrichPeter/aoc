package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var REFERENCE_STRING = "XMAS"
var WINDOW_LENGTH = len(REFERENCE_STRING)

func readGrid(filename string) [][]rune {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var grid [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []rune(line))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return grid
}

var directions = [][2]int{
	{0, 1}, {1, 0}, {1, 1}, {1, -1}, // right, down, diagonal down-right, diagonal down-left
	{0, -1}, {-1, 0}, {-1, -1}, {-1, 1}, // left, up, diagonal up-left, diagonal up-right
}

func checkWord(grid [][]rune, x, y int, dir [2]int) bool {
	for k := 0; k < WINDOW_LENGTH; k++ {
		nx, ny := x+k*dir[0], y+k*dir[1]
		if nx < 0 || ny < 0 || nx >= len(grid) || ny >= len(grid[0]) || grid[nx][ny] != rune(REFERENCE_STRING[k]) {
			return false
		}
	}
	return true
}

func countXMAS(grid [][]rune) int {
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			for _, dir := range directions {
				if checkWord(grid, i, j, dir) {
					count++
				}
			}
		}
	}
	return count
}

func main() {
	grid := readGrid("input.txt")
	result := countXMAS(grid)
	fmt.Println(result)
}
