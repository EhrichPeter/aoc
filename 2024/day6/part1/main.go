package main

import (
	"bufio"
	"fmt"
	"os"
)

func LoadLinesFromFile(filePath string) ([][]rune, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var grid [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grid = append(grid, []rune(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return grid, nil
}

type GameObject interface {
	Position() [2]int
	PrintMarker()
}

type Tile interface {
	GameObject
	Visited() bool
	Visit()
}

type Floor struct {
	position [2]int
	visited  bool
}

func (f *Floor) Position() [2]int {
	return f.position
}

func (f *Floor) PrintMarker() {
	if f.visited {
		fmt.Print("x")
	} else {
		fmt.Print(".")
	}
}

func (f *Floor) Visited() bool {
	return f.visited
}

func (f *Floor) Visit() {
	f.visited = true
}

type Obstacle struct {
	position [2]int
}

func (o *Obstacle) Position() [2]int {
	return o.position
}

func (o *Obstacle) PrintMarker() {
	fmt.Print("#")
}

func (o *Obstacle) Visited() bool {
	return false
}

func (o *Obstacle) Visit() {
	panic("Cannot visit an obstacle")
}

type Guard struct {
	position    [2]int
	orientation int
}

func (g *Guard) Position() [2]int {
	return g.position
}

func (g *Guard) PrintMarker() {
	fmt.Print(string(g.OrientationToMarker()))
}

func (g *Guard) SetPosition(x, y int) {
	g.position = [2]int{x, y}
}

func (g *Guard) Orientation() int {
	return g.orientation
}

func (g *Guard) SetOrientation(orientation int) {
	g.orientation = orientation
}

func (g *Guard) MarkerToOrientation(marker rune) {
	orientations := map[rune]int{
		'^': 0,
		'>': 1,
		'v': 2,
		'<': 3,
	}

	g.orientation = orientations[marker]
}

func (g *Guard) OrientationToMarker() rune {
	markers := map[int]rune{
		0: '^',
		1: '>',
		2: 'v',
		3: '<',
	}

	return markers[g.orientation]
}

func (g *Guard) Print() {
	fmt.Println("Guard at", g.position, "facing", string(g.OrientationToMarker()))
}

type Game struct {
	guard   Guard
	terrain [][]Tile
}

func InitializeGame(input [][]rune) *Game {
	var terrain [][]Tile
	guard := &Guard{}

	for i, row := range input {
		var objects []Tile
		for j, cell := range row {
			pos := [2]int{i, j}
			switch cell {
			case '.':
				objects = append(objects, &Floor{position: pos})
			case '#':
				objects = append(objects, &Obstacle{position: pos})
			case '^', '>', 'v', '<':
				guard.SetPosition(i, j)
				guard.MarkerToOrientation(cell)
				firstFloor := &Floor{position: pos}
				firstFloor.Visit()
				objects = append(objects, firstFloor)
			}
		}
		terrain = append(terrain, objects)
	}

	return &Game{
		guard:   *guard,
		terrain: terrain,
	}
}

func (g *Game) Print() {
	for _, row := range g.terrain {
		for _, cell := range row {
			cell.PrintMarker()
		}
		fmt.Println()
	}
}

func (g *Game) countVisitedTiles() int {
	count := 0
	for _, row := range g.terrain {
		for _, cell := range row {
			if cell.Visited() {
				count++
			}
		}
	}

	fmt.Println("Visited cells:", count)
	return count
}

func main() {
	grid, err := LoadLinesFromFile("input.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	game := InitializeGame(grid)
	game.Print()
	game.guard.Print()
	game.countVisitedTiles()
}
