package main

import (
	"fmt"
	"os"
)

const (
	startOfPacketLength  = 4
	startOfMessageLength = 14
)

func main() {
	input, err := os.ReadFile("../input.txt")
	if err != nil {
		panic(err)
	}

	pos, _ := findMarker(input, startOfPacketLength)
	fmt.Printf("Part 1: The number of processed characters: %d\n", pos+startOfPacketLength)

	// Part 2
	pos, _ = findMarker(input, startOfMessageLength)
	fmt.Printf("Part 1: The number of processed characters: %d\n", pos+startOfMessageLength)
}

func findMarker(input []byte, markerSize int) (position int, marker []byte) {
	for i := 0; i < len(input); i++ {
		if len(input) > i+markerSize {
			buf := input[i : i+markerSize]
			if isUnique(buf) {
				return i, buf
			}
		}
	}

	return 0, nil
}

func isUnique(buf []byte) bool {
	m := make(map[byte]int)
	for _, c := range buf {
		m[c]++
	}
	for _, v := range m {
		if v > 1 {
			return false
		}
	}

	return true
}
