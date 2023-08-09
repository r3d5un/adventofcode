package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ElfSections struct {
	Lower int
	Upper int
}

func main() {
	input, err := os.ReadFile("../input.txt")
	if err != nil {
		fmt.Printf("error reading file: %v", err)
		os.Exit(1)
	}

	partOneScore, err := partOne(input)
	if err != nil {
		fmt.Printf("error getting part one score: %v", err)
		os.Exit(1)
	}
	fmt.Println("Part One Score:", partOneScore)

	partTwoScore, err := partTwo(input)
	if err != nil {
		fmt.Printf("error getting part two score: %v", err)
		os.Exit(1)
	}
	fmt.Println("Part Two Score:", partTwoScore)
}

func parseLines(input []byte) (values [][]string) {
	for _, line := range strings.Split(string(input), "\n") {
		if line != "" {
			values = append(values, strings.Split(line, ","))
		}
	}
	return values
}

func getBounds(value string) (low, high int, err error) {
	values := strings.Split(value, "-")
	lower, err := strconv.Atoi(values[0])
	if err != nil {
		fmt.Printf("error converting %s to lower bound int: %v", values[0], err)
		return 0, 0, err
	}
	upper, err := strconv.Atoi(values[1])
	if err != nil {
		fmt.Printf("error converting %s to upper bound int: %v", values[1], err)
		return 0, 0, err
	}

	return lower, upper, nil
}

func partOne(input []byte) (score int, err error) {
	score = 0
	for _, line := range parseLines(input) {
		elf1 := ElfSections{}
		elf1.Lower, elf1.Upper, err = getBounds(line[0])
		if err != nil {
			fmt.Printf("error getting bounds for elf one: %v", err)
			return 0, err
		}

		elf2 := ElfSections{}
		elf2.Lower, elf2.Upper, err = getBounds(line[1])
		if err != nil {
			fmt.Printf("error getting bounds for elf two: %v", err)
			return 0, err
		}

		if (elf1.Lower >= elf2.Lower && elf1.Upper <= elf2.Upper) || (elf2.Lower >= elf1.Lower && elf2.Upper <= elf1.Upper) {
			score++
		}
	}

	return score, nil
}

func partTwo(input []byte) (score int, err error) {
	score = 0
	for _, line := range parseLines(input) {
		elf1 := ElfSections{}
		elf1.Lower, elf1.Upper, err = getBounds(line[0])
		if err != nil {
			fmt.Printf("error getting bounds for elf one: %v", err)
			return 0, err
		}

		elf2 := ElfSections{}
		elf2.Lower, elf2.Upper, err = getBounds(line[1])
		if err != nil {
			fmt.Printf("error getting bounds for elf two: %v", err)
			return 0, err

		}

		if !(elf1.Lower > elf2.Upper || elf1.Upper < elf2.Lower) {
			score++
		}
	}
	return score, nil
}
