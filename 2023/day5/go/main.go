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

	SeedToSoil := ParseMapping(input, "seed-to-soil")
	fmt.Println(SeedToSoil)
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

func ParseMapping(input string, mapping string) (ranges []Range) {
	var values []string
	takeNext := false

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if strings.Contains(line, mapping) {
			takeNext = true
			continue
		}
		if line == "" && takeNext {
			break
		}
		if takeNext {
			values = append(values, line)
		}
	}

	for _, value := range values {
		ranges = append(ranges, ParseRange(value))
	}

	return ranges
}

func ParseRange(value string) Range {
	mapValues := strings.Fields(value)
	fmt.Println(mapValues)

	destinationRangeStart, err := strconv.Atoi(mapValues[0])
	if err != nil {
		panic(err)
	}

	sourceRangeStart, err := strconv.Atoi(mapValues[1])
	if err != nil {
		panic(err)
	}

	rangeLength, err := strconv.Atoi(mapValues[2])
	if err != nil {
		panic(err)
	}

	return Range{
		DestinationRangeStart: destinationRangeStart,
		SourceRangeStart:      sourceRangeStart,
		RangeLength:           rangeLength,
	}
}

type Range struct {
	DestinationRangeStart int
	SourceRangeStart      int
	RangeLength           int
}
