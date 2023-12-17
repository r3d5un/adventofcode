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
	fmt.Printf("Part 2: %d\n", Part2(input))
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

func Part2(input string) (answer int) {
	race := ParsePart2(input)
	fmt.Printf("race: %v\n", race)

	for i := 0; i < race.Time; i++ {
		win, _ := race.TryRace(i)

		if win {
			answer++
		}
	}

	return answer
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

func ParsePart2(input string) (race Race) {
	var timeString string
	var distanceString string

	for _, line := range strings.Split(input, "\n") {
		if strings.Contains(line, "Time:") {
			timeString = strings.Split(line, ":")[1]
			timeString = strings.Join(strings.Fields(timeString), "")
		}
		if strings.Contains(line, "Distance:") {
			distanceString = strings.Split(line, ":")[1]
			distanceString = strings.Join(strings.Fields(distanceString), "")
		}
	}

	time, err := strconv.Atoi(timeString)
	if err != nil {
		panic(err)
	}
	distance, err := strconv.Atoi(distanceString)
	if err != nil {
		panic(err)
	}

	return Race{Time: time, DistanceRecord: distance}
}

func multiply(array []int) int {
	product := 1

	for _, value := range array {
		product *= value
	}

	return product
}
