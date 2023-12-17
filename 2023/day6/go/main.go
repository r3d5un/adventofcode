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

func Part1(input string) (answer int) {
	races := Parse(input)
	fmt.Printf("races: %v\n", races)

	winsPerRace := []int{}
	for _, r := range races {
		winsPerRace = append(winsPerRace, len(r.GetWinnableAttempts()))
	}

	return multiply(winsPerRace)
}

type Race struct {
	Time           int
	DistanceRecord int
}

func (r *Race) GetWinnableAttempts() (distances []int) {
	for i := 0; i < r.Time; i++ {
		win, distance := r.TryRace(i)

		if win {
			distances = append(distances, distance)
		}

	}

	return distances
}

func (r *Race) TryRace(buttonHoldTime int) (win bool, distance int) {
	acceleration := 1 * buttonHoldTime
	travelTime := r.Time - buttonHoldTime

	if travelTime <= 0 || acceleration <= 0 {
		return false, 0
	}

	distance = acceleration * travelTime
	win = distance > r.DistanceRecord

	return win, distance
}

func Parse(input string) (races []Race) {
	var timeString string
	var distanceString string

	for _, line := range strings.Split(input, "\n") {
		if strings.Contains(line, "Time:") {
			timeString = strings.Split(line, ":")[1]
		}
		if strings.Contains(line, "Distance:") {
			distanceString = strings.Split(line, ":")[1]
		}
	}

	timeSplits := strings.Fields(timeString)
	distanceSplits := strings.Fields(distanceString)

	for i := 0; i < len(timeSplits); i++ {
		time, err := strconv.Atoi(timeSplits[i])
		if err != nil {
			panic(err)
		}
		distance, err := strconv.Atoi(distanceSplits[i])
		if err != nil {
			panic(err)
		}

		races = append(races, Race{Time: time, DistanceRecord: distance})
	}

	return races
}

func multiply(array []int) int {
	product := 1

	for _, value := range array {
		product *= value
	}

	return product
}
