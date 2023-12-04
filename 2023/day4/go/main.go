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
	fmt.Println("Part 1:", Part1(input))
}

func Part1(input string) (sum int) {
	cards := []Card{}
	for _, line := range strings.Split(input, "\n") {
		card := ParseCard(line)
		cards = append(cards, card)
	}

	for _, card := range cards {
		sum += card.CalculateScore()
	}

	return sum
}

type Card struct {
	CardID         string
	WinningNumbers []int
	MyNumbers      []int
}

func ParseCard(line string) (card Card) {
	cardIDSplit := strings.Split(line, ":")

	card.CardID = cardIDSplit[0]

	numbersSplit := strings.Split(cardIDSplit[1], "|")
	winningNumbersSlice := strings.Fields(numbersSplit[0])
	myNumbersSlice := strings.Fields(numbersSplit[1])

	card.WinningNumbers = append(card.WinningNumbers, card.stringSliceToIntSlice(winningNumbersSlice)...)
	card.MyNumbers = append(card.MyNumbers, card.stringSliceToIntSlice(myNumbersSlice)...)

	return card
}

func (c Card) stringSliceToIntSlice(slice []string) (intSlice []int) {
	for _, number := range slice {
		number, err := strconv.Atoi(number)
		if err != nil {
			panic(err)
		}
		intSlice = append(intSlice, number)
	}

	return intSlice
}

func (c Card) CalculateScore() (score int) {
	for _, current := range c.MyNumbers {
		if c.isWinningNumber(current) {
			if score == 0 {
				score = 1
			} else {
				score *= 2
			}
		}
	}

	return score
}

func (c Card) isWinningNumber(number int) (isWinningNumber bool) {
	for _, winningNumber := range c.WinningNumbers {
		if number == winningNumber {
			return true
		}
	}

	return false
}
