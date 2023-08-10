package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Instruction struct {
	From   int
	Amount int
	To     int
}

func main() {
	input, err := os.ReadFile("../input.txt")
	if err != nil {
		fmt.Printf("Unable to read file: %s\n", err)
		os.Exit(1)
	}

	stacks, instructions := parseLines(string(input))
	fmt.Printf("Part1: %s\n", partOne(stacks, instructions))

	stacks, instructions = parseLines(string(input))
	fmt.Printf("Part2: %s\n", partTwo(stacks, instructions))
}

func partOne(stacks map[int][]rune, instructions []Instruction) (answer string) {
	for _, instruction := range instructions {
		for quantity := 0; quantity < instruction.Amount; quantity++ {
			movingCrate := stacks[instruction.From][len(stacks[instruction.From])-1]
			stacks[instruction.From] = stacks[instruction.From][:len(stacks[instruction.From])-1]

			stacks[instruction.To] = append(stacks[instruction.To], movingCrate)
		}
	}

	for i := 1; i < len(stacks)+1; i++ {
		answer += string(stacks[i][len(stacks[i])-1])
	}

	return answer
}

func partTwo(stacks map[int][]rune, instructions []Instruction) (answer string) {
	for _, instruction := range instructions {
		fromIndex := len(stacks[instruction.From]) - instruction.Amount
		stacks[instruction.To] = append(stacks[instruction.To], stacks[instruction.From][fromIndex:]...)
		stacks[instruction.From] = stacks[instruction.From][:fromIndex]
	}

	for i := 1; i < len(stacks)+1; i++ {
		answer += string(stacks[i][len(stacks[i])-1])
	}

	return answer
}

func parseLines(lines string) (stacks map[int][]rune, instructions []Instruction) {
	parts := strings.Split(lines, "\n\n")

	stackPart := strings.Split(parts[0], "\n")

	// Parse and create stacks
	stackIDStr := strings.Split(stackPart[len(stackPart)-1], " ")
	stacks = make(map[int][]rune)
	for _, stackID := range stackIDStr {
		if stackID != "" {
			id, err := strconv.Atoi(stackID)
			if err != nil {
				fmt.Printf("Unable to convert stackID to int: %s\n", err)
				os.Exit(1)
			}
			stacks[id] = []rune{}
		}
	}

	// Perform initial stack setup
	// TODO: This whole thing is just pure awfulness... Future me... sorry...
	containers := stackPart[:len(stackPart)-1]
	cIdx := 1
	for key := 1; key < len(stacks)+1; key++ {
		for i := range containers {
			d := string(containers[len(containers)-1-i][cIdx])

			runes := []rune{}
			if d != " " {
				runes = append(runes, rune(d[0]))
			}

			stacks[key] = append(stacks[key], runes...)
		}
		cIdx += 4
	}

	// Parse instructions
	instructions = []Instruction{}
	for _, line := range strings.Split(parts[1], "\n") {
		instruction := Instruction{}
		_, err := fmt.Sscanf(
			line,
			"move %d from %d to %d",
			&instruction.Amount,
			&instruction.From,
			&instruction.To,
		)
		if err != nil {
			panic(err)
		}

		instructions = append(instructions, instruction)
	}

	return stacks, instructions
}
