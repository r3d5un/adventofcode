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

	fmt.Printf("Part 1: %d\n", Part1(input))
}

func Part1(input string) int {
	hands := Parse(input)

	for _, hand := range hands {
		fmt.Println(hand)
	}

	return 0
}

type Hand struct {
	Cards []Card
	Bid   int
	Rank  int
}

type Card struct {
	kind  string
	value int
}

func NewCard(cs string) (c Card, err error) {
	switch cs {
	case "A":
		return Card{"A", 13}, nil
	case "K":
		return Card{"K", 12}, nil
	case "Q":
		return Card{"Q", 11}, nil
	case "J":
		return Card{"J", 10}, nil
	case "T":
		return Card{"T", 9}, nil
	case "9":
		return Card{"9", 8}, nil
	case "8":
		return Card{"8", 7}, nil
	case "7":
		return Card{"7", 6}, nil
	case "6":
		return Card{"6", 5}, nil
	case "5":
		return Card{"5", 4}, nil
	case "4":
		return Card{"4", 3}, nil
	case "3":
		return Card{"3", 2}, nil
	case "2":
		return Card{"2", 1}, nil
	default:
		return Card{"0", 0}, err
	}

}

func Parse(input string) []Hand {
	hands := []Hand{}
	for _, line := range strings.Split(input, "\n") {
		hand := Hand{}

		handString := strings.Fields(line)

		cardString := handString[0]
		cardStrings := strings.Split(cardString, "")
		for _, cs := range cardStrings {
			card, err := NewCard(cs)
			if err != nil {
				panic(err)
			}

			hand.Cards = append(hand.Cards, card)
		}

		bidString := handString[1]
		bid, err := strconv.Atoi(bidString)
		if err != nil {
			panic(err)
		}
		hand.Bid = bid

		hands = append(hands, hand)
	}

	return hands
}
