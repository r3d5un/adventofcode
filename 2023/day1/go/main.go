package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	input, err := os.ReadFile("../input.txt")
	if err != nil {
		panic(err)
	}

	p1, err := Part1(string(input))
	if err != nil {
		panic(err)
	}

	fmt.Println("Part 1:", p1)

	p2, err := Part2(string(input))
	if err != nil {
		panic(err)
	}

	fmt.Println("Part 2:", p2)
}

func Part1(input string) (int, error) {
	var results []int
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		decem, err := getFirstNumber(scanner.Text())
		if err != nil {
			continue
		}
		decem = decem * 10

		revesed := reverseString(scanner.Text())
		uno, err := getFirstNumber(revesed)
		if err != nil {
			continue
		}

		result := decem + uno

		results = append(results, result)
	}

	return sumSlice(results), nil
}

func reverseString(s string) string {
	var result string
	for _, c := range s {
		result = string(c) + result
	}
	return result
}

func getFirstNumber(s string) (int, error) {
	for _, c := range s {
		if unicode.IsDigit(c) {
			i, err := strconv.Atoi(string(c))
			if err != nil {
				return 0, err
			}

			return i, nil
		}
	}
	return 0, fmt.Errorf("no number found")
}

func sumSlice(s []int) int {
	var result int
	for _, i := range s {
		result += i
	}

	return result
}

func Part2(input string) (int, error) {
	results := []int{}
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		matches := getMatches(scanner.Text())
		if len(matches) == 0 {
			continue
		}

		decem := getFirstMatch(matches).value
		uno := getLastMatch(matches).value

		result := decem*10 + uno

		results = append(results, result)
	}

	return sumSlice(results), nil
}

type Match struct {
	idx   int
	value int
}

func getFirstMatch(matches []Match) Match {
	var result Match
	if len(matches) > 0 {
		result = matches[0]
	}
	for _, m := range matches {
		if m.idx < result.idx {
			result = m
		}
	}

	return result
}

func getLastMatch(matches []Match) (result Match) {
	if len(matches) > 0 {
		result = matches[len(matches)-1]
	}
	for _, m := range matches {
		if m.idx > result.idx {
			result = m
		}
	}

	return result
}

func getMatches(s string) []Match {
	strNum := []string{
		"one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
		"1", "2", "3", "4", "5", "6", "7", "8", "9",
	}
	matches := []Match{}

	for _, str := range strNum {
		idx := strings.Index(s, str)
		if idx != -1 {
			value := convertToIntegerValue(str)
			matches = append(matches, Match{idx: idx, value: value})
		}
	}

	for _, str := range strNum {
		idx := strings.LastIndex(s, str)
		if idx != -1 {
			value := convertToIntegerValue(str)
			matches = append(matches, Match{idx: idx, value: value})
		}
	}

	return matches
}

func convertToIntegerValue(s string) int {
	switch s {
	case "one", "1":
		return 1
	case "two", "2":
		return 2
	case "three", "3":
		return 3
	case "four", "4":
		return 4
	case "five", "5":
		return 5
	case "six", "6":
		return 6
	case "seven", "7":
		return 7
	case "eight", "8":
		return 8
	case "nine", "9":
		return 9
	default:
		return 0
	}
}
