package main

import (
	"bufio"
	"fmt"
	"os"
)

var REFERENCE_STRING = "MAS"
var WINDOW_LENGTH = len(REFERENCE_STRING)

var directions = [][2]int{
	{1, 1}, {1, -1}, {-1, -1}, {-1, 1},
}

func readGrid(filename string) [][]rune {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var grid [][]rune
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()
		grid = append(grid, []rune(text))
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return grid
}

func checkWord(grid [][]rune, x int, y int, dir [2]int) bool {
	sx, sy := x+(-dir[1]), y+(-dir[0]) //Move starting coordinates inverse of direction by one movement

	for k := 0; k < WINDOW_LENGTH; k++ {
		nx, ny := sx+k*dir[1], sy+k*dir[0]
		if nx < 0 || ny < 0 || ny >= len(grid) || nx >= len(grid[0]) || grid[ny][nx] != rune(REFERENCE_STRING[k]) {
			return false
		}
	}
	return true
}

func countXMAS(grid [][]rune) int {
	count := 0
	for i, row := range grid {
		for j := range row {
			ignoreDirections := make(map[[2]int]bool)
			diagonals := 0
			for _, dir := range directions {
				if ignoreDirections[dir] {
					continue
				}
				if checkWord(grid, i, j, dir) {
					inverseDir := [2]int{-dir[0], -dir[1]}
					ignoreDirections[inverseDir] = true
					diagonals++
				}
			}
			if diagonals == 2 {
				count++
			}
		}
	}
	return count

}

func main() {
	grid := readGrid("input.txt")
	count := countXMAS(grid)
	fmt.Println(count)
}
