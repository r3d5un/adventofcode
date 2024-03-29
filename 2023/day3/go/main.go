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
	fmt.Println("Part 2:", Part2(input))
}

func Part1(input string) int {
	result := 0
	grid := NewGrid(input)

	minX, minY, maxX, maxY := grid.Bounds()

	// line by line
	for y := minY; y <= maxY; y++ {
		var current int // value of current position
		var symbol bool

		// character by character
		for x := minX; x <= maxX; x++ {
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

func Part2(input string) int {
	result := 0
	grid := NewGrid(input)
	stars := make(Stars)
	starMap := make(StarMap)
	minX, minY, maxX, maxY := grid.Bounds()

	for y := minY; y <= maxY; y++ {
		var current int
		for x := minX; x <= maxX; x++ {
			if value, ok := grid[Coordinate{x, y}]; ok && value >= '0' && value <= '9' {
				current = 10*current + int(value-'0')
				// get all adjacent stars
				stars.AddAll(grid.getAdjacentStars(Coordinate{x, y})...)
			} else {
				if stars.Len() > 0 {
					for star := range stars {
						starMap[star] = append(starMap[star], current)
					}
				}
				current = 0
				stars.Clear()
			}
		}
		if stars.Len() > 0 {
			for star := range stars {
				starMap[star] = append(starMap[star], current)
			}
		}
		stars.Clear()
	}

	// loop through all stars. If a star was added twice, it is a star that is
	// visible from two values. Multiply the two values and add to result.
	for _, values := range starMap {
		if len(values) == 2 {
			result += values[0] * values[1]
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

func (grid Grid) getAdjacentStars(coordinate Coordinate) (result []Coordinate) {
	for _, neighbor := range coordinate.GetNeighbours() {
		if value, ok := grid[neighbor]; ok && value == '*' {
			result = append(result, neighbor)
		}
	}

	return result
}

type StarMap map[Coordinate][]int

type Stars map[Coordinate]struct{}

func (stars Stars) Len() int {
	return len(stars)
}

func (stars Stars) Clear() {
	for star := range stars {
		delete(stars, star)
	}
}

func (stars Stars) Add(star Coordinate) {
	stars[star] = struct{}{}
}

func (stars Stars) AddAll(other ...Coordinate) {
	for _, star := range other {
		stars.Add(star)
	}
}
