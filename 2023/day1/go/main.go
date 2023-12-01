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
	input, err := os.ReadFile("../sample2.txt")
	if err != nil {
		panic(err)
	}

	p1, err := solvePuzzle(string(input))
	if err != nil {
		panic(err)
	}

	fmt.Println("Part 1:", p1)
}

func solvePuzzle(input string) (int, error) {
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
