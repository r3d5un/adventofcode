package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	input, err := os.ReadFile("../input.txt")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}

	rucksacks := parseLines(input)

	partOneSum := 0
	for _, line := range rucksacks {
		part1, part2, err := splitIntoCompartments(line)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			os.Exit(1)
		}

		intersection := getIntersectionOfStrings(part1, part2)

		charVal, err := convertRuneToIntValue(intersection[0])
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			os.Exit(1)
		}

		partOneSum += charVal
	}
	fmt.Printf("Part 1 sum is %d\n", partOneSum)

	partTwoSum := 0
	for i := 0; i < len(rucksacks); i += 3 {
		char, err := getBadge(rucksacks[i], rucksacks[i+1], rucksacks[i+2])
		charVal, err := convertRuneToIntValue(char)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			os.Exit(1)
		}
		partTwoSum += charVal
	}
	fmt.Printf("Part 2 sum is %d\n", partTwoSum)
}

func parseLines(input []byte) (values []string) {
	for _, line := range strings.Split(string(input), "\n") {
		if line != "" {
			values = append(values, line)
		}
	}
	return values
}

func splitIntoCompartments(contents string) (part1 string, part2 string, err error) {
	contentLength := len(contents)
	if contentLength < 0 {
		return "", "", fmt.Errorf("Empty string")
	}
	part1 = contents[:contentLength/2]
	part2 = contents[contentLength/2:]

	return part1, part2, nil
}

func getIntersectionOfStrings(a, b string) []rune {
	h := make(map[rune]bool)
	for _, c := range a {
		h[c] = true
	}
	var ret []rune
	for _, c := range b {
		if h[c] {
			ret = append(ret, c)
		}
	}
	return ret
}

func convertRuneToIntValue(char rune) (value int, err error) {
	if !unicode.IsLetter(char) && !unicode.IsDigit(char) {
		return 0, fmt.Errorf("Invalid character")
	}

	initUpperVal := 'A'
	initLowerChar := 'a'

	if unicode.IsUpper(char) {
		return int(char-initUpperVal) + 27, nil
	}

	return int(char-initLowerChar) + 1, nil
}

func getBadge(part1, part2, part3 string) (char rune, err error) {
	h := make(map[rune]bool)
	var ret []rune

	// Note: Could reuse getIntersectionOfStrings, but I wanted the function to be self-contained
	for _, c := range part1 {
		h[c] = true
	}
	for _, c := range part2 {
		if h[c] {
			ret = append(ret, c)
		}
	}
	for _, char := range ret {
		if strings.Contains(part3, string(char)) {
			return char, nil
		}
	}

	return 0, fmt.Errorf("No badge found")
}
