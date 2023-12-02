package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	rMax = 12
	gMax = 13
	bMax = 14
)

func main() {
	input, err := os.ReadFile("../input.txt")
	if err != nil {
		panic(err)
	}

	part1Sum := 0

	scanner := bufio.NewScanner(strings.NewReader(string(input)))
	for scanner.Scan() {
		possible := true

		gameID, err := getGameID(scanner.Text())
		if err != nil {
			panic(err)
		}

		split := strings.Split(scanner.Text(), ":")
		if err != nil {
			panic(err)
		}

		gameSet, err := parseGameSet(split[1])
		if err != nil {
			panic(err)
		}

		for _, draw := range gameSet {
			switch draw.Color {
			case "red":
				if draw.Number > rMax {
					possible = false
				}
			case "green":
				if draw.Number > gMax {
					possible = false
				}
			case "blue":
				if draw.Number > bMax {
					possible = false
				}
			}
		}

		if possible {
			part1Sum += gameID
		}
	}

	fmt.Println("Part 1 Sum:", part1Sum)
}

func getGameID(line string) (int, error) {
	split := strings.Split(line, ":")
	gameID, err := strconv.Atoi(strings.Fields(split[0])[1])
	if err != nil {
		return 0, err
	}
	return gameID, nil
}

type Draw struct {
	Color  string
	Number int
}

func parseGameSet(line string) (set []Draw, err error) {
	split := strings.Split(line, ";")
	for _, a := range split {
		drawStrings := strings.Split(a, ",")
		for _, b := range drawStrings {
			d := strings.Fields(strings.TrimSpace(b))
			value, err := strconv.Atoi(d[0])
			if err != nil {
				return []Draw{}, err
			}

			draw := Draw{Color: d[1], Number: value}

			set = append(set, draw)
		}
	}

	return set, nil
}
