package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fileContents, err := os.ReadFile("../sample.txt")
	if err != nil {
		panic(err)
	}

	input := strings.TrimSuffix(string(fileContents), "\n")
	Part1(input)
}

func Part1(input string) {
	seeds := ParseSeeds(input)
	fmt.Println(seeds)
}

func ParseSeeds(input string) (seeds []int) {
	var seedLine string
	for _, line := range strings.Split(input, "\n") {
		if strings.Contains(line, "seeds:") {
			seedLine = strings.Split(line, ":")[1]
			break
		}
	}

	seedStrings := strings.Fields(seedLine)

	for _, seedString := range seedStrings {
		seed, err := strconv.Atoi(seedString)
		if err != nil {
			panic(err)
		}
		seeds = append(seeds, seed)
	}

	return seeds
}
