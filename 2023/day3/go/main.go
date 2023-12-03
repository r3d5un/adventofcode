package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	fileContents, err := os.ReadFile("../input.txt")
	if err != nil {
		panic(err)
	}

	input := strings.TrimSuffix(string(fileContents), "\n")

	fmt.Println("Part 1:", Part1(input))
}

func Part1(input string) int {
	result := 0
	grid := NewGrid(input)

	// line by line
	for y := grid.MinY(); y <= grid.MaxY(); y++ {
		var current int // value of current position
		var symbol bool

		// character by character
		for x := grid.MinX(); x <= grid.MaxX(); x++ {
			// if current position is a digit
			if value, ok := grid[Coordinate{x, y}]; ok && value >= '0' && value <= '9' {
				// multiply current by 10 and add the value of the digit
				current = 10*current + int(value-'0')
				// check if any neighbors is a symbol
				symbol = symbol || grid.isSymbolAdjacent(Coordinate{x, y})
			} else {
				if symbol {
					result += current
				}
				current = 0
				symbol = false
			}
		}
		if symbol {
			result += current
		}
	}

	return result
}

type Coordinate struct {
	X int
	Y int
}

func (coordinate Coordinate) GetNeighbours() []Coordinate {
	return []Coordinate{
		{coordinate.X - 1, coordinate.Y},
		{coordinate.X + 1, coordinate.Y},
		{coordinate.X, coordinate.Y - 1},
		{coordinate.X, coordinate.Y + 1},
		{coordinate.X - 1, coordinate.Y - 1},
		{coordinate.X + 1, coordinate.Y - 1},
		{coordinate.X - 1, coordinate.Y + 1},
		{coordinate.X + 1, coordinate.Y + 1},
	}
}

type Grid map[Coordinate]int

func NewGrid(input string) Grid {
	grid := make(Grid)
	lines := strings.Split(input, "\n")

	for y, line := range lines {
		for x, char := range line {
			grid[Coordinate{x, y}] = int(char)
		}
	}

	return grid
}

func (grid Grid) Print() {
	minX, minY, maxX, maxY := grid.Bounds()
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			fmt.Print(string(grid[Coordinate{x, y}]))
		}
		fmt.Println()
	}
}

func (grid Grid) MinX() int {
	minX := math.MaxInt64
	for coordinate := range grid {
		if coordinate.X < minX {
			minX = coordinate.X
		}
	}
	return minX
}

func (grid Grid) MinY() int {
	minY := math.MaxInt64
	for coordinate := range grid {
		if coordinate.Y < minY {
			minY = coordinate.Y
		}
	}
	return minY
}

func (grid Grid) MaxX() int {
	maxX := math.MinInt64
	for coordinate := range grid {
		if coordinate.X > maxX {
			maxX = coordinate.X
		}
	}
	return maxX
}

func (grid Grid) MaxY() int {
	maxY := math.MinInt64
	for coordinate := range grid {
		if coordinate.Y > maxY {
			maxY = coordinate.Y
		}
	}
	return maxY
}

func (grid Grid) Bounds() (minX, minY, maxX, maxY int) {
	return grid.MinX(), grid.MinY(), grid.MaxX(), grid.MaxY()
}

func (grid Grid) isSymbolAdjacent(coordinate Coordinate) bool {
	for _, neighbor := range coordinate.GetNeighbours() {
		if value, ok := grid[neighbor]; ok && value != '.' && !(value >= '0' && value <= '9') {
			return true
		}
	}

	return false
}
