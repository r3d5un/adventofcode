package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	Win  = 6
	Loss = 0
	Draw = 3

	Rock     = 1
	Paper    = 2
	Scissors = 3
)

func main() {
	input, err := os.ReadFile("../input.txt")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}

	partOneScore, err := partOne(parseLines(input))
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("Part 1: %d\n", partOneScore)

}

func partOne(input [][]string) (score int, err error) {
	myChoice := map[string]int{
		"X": Rock,
		"Y": Paper,
		"Z": Scissors,
	}

	score = 0

	for _, line := range input {
		if _, ok := myChoice[line[1]]; !ok {
			return 0, fmt.Errorf("invalid player input: %s", line[0])
		}

		score += myChoice[line[1]]

		switch line[1] {
		case "X":
			switch line[0] {
			case "A":
				score += Draw
			case "B":
				score += Loss
			case "C":
				score += Win
			default:
				return 0, fmt.Errorf("invalid opponent input: %s", line[0])
			}
		case "Y":
			switch line[0] {
			case "A":
				score += Win
			case "B":
				score += Draw
			case "C":
				score += Loss
			default:
				return 0, fmt.Errorf("invalid opponent input: %s", line[0])
			}
		case "Z":
			switch line[0] {
			case "A":
				score += Loss
			case "B":
				score += Win
			case "C":
				score += Draw
			default:
				return 0, fmt.Errorf("invalid opponent input: %s", line[0])
			}
		default:
			return 0, fmt.Errorf("invalid opponent input: %s", line[0])
		}
	}

	return score, nil
}

func parseLines(input []byte) (values [][]string) {
	for _, line := range strings.Split(string(input), "\n") {
		if line != "" {
			values = append(values, strings.Split(line, " "))
		}
	}
	return values
}
