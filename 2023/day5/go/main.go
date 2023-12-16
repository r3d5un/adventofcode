package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fileContents, err := os.ReadFile("../input.txt")
	if err != nil {
		panic(err)
	}

	input := strings.TrimSuffix(string(fileContents), "\n")
	fmt.Printf("Part 1: %d\n", Part1(input))
}

func Part1(input string) int {
	mapNames := []string{
		"seed-to-soil",
		"soil-to-fertilizer",
		"fertilizer-to-water",
		"water-to-light",
		"light-to-temperature",
		"temperature-to-humidity",
		"humidity-to-location",
	}

	mappings := make(map[string]RangeSet)

	for _, ms := range mapNames {
		mappings[ms] = NewRangeSet(input, ms)
	}

	tracker := make(map[int][]int)
	for _, seed := range ParseSeeds(input) {
		tracker[seed] = []int{seed}
		t := seed

		for _, v := range mapNames {
			rs := mappings[v]

			t = rs.GetDestinationValue(tracker[seed][len(tracker[seed])-1])
			tracker[seed] = append(tracker[seed], t)
		}
	}

	var lowestLocation int
	for _, v := range tracker {
		if lowestLocation == 0 {
			lowestLocation = v[len(v)-1]
		}
		if v[len(v)-1] < lowestLocation {
			lowestLocation = v[len(v)-1]
		}
	}

	return lowestLocation
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

func (r Range) IsInSourceRange(value int) bool {
	return value >= r.SourceRangeStart && value <= (r.SourceRangeStart+r.RangeLength)
}

func (r Range) IsInDestinationRange(value int) bool {
	return value >= r.DestinationRangeStart && value <= (r.DestinationRangeStart+r.RangeLength)
}

func (r Range) GetDestinationValue(v int) int {
	return r.DestinationRangeStart + (v - r.SourceRangeStart)
}

type RangeSet []Range

func (rs *RangeSet) Add(r Range) {
	*rs = append(*rs, r)
}

func (rs *RangeSet) GetDestinationValue(v int) int {
	for _, r := range *rs {
		if r.IsInSourceRange(v) {
			return r.GetDestinationValue(v)
		}
	}

	return v
}

func NewRangeSet(input string, mapping string) (rs RangeSet) {
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
		rs.Add(ParseRange(value))
	}

	return rs
}
