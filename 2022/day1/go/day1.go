package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input_bytes, err := os.ReadFile("../input.txt")
	if err != nil {
		fmt.Print(err)
	}

	input_text := string(input_bytes)
	group_list := strings.Split(input_text, "\n\n")

	var calories_per_elf_list []int

	for _, group := range group_list {
		var calorie_string_list []string = strings.Split(strings.TrimSpace(group), "\n")
		var calorie_number_list []int

		for _, calorie_number := range calorie_string_list {
			calorie_value, err := strconv.Atoi(calorie_number)
			if err != nil {
				fmt.Print(err)
			}
			calorie_number_list = append(calorie_number_list, calorie_value)
		}
		var sum int = 0
		for _, value := range calorie_number_list {
			sum += value
		}
		calories_per_elf_list = append(calories_per_elf_list, sum)
	}
	sort.Ints(calories_per_elf_list)
	fmt.Println("Answer part one: ", calories_per_elf_list[len(calories_per_elf_list)-1])

	var sum int = 0
	for _, value := range calories_per_elf_list[len(calories_per_elf_list)-3:] {
		sum += value
	}

	fmt.Println("Answer part two: ", sum)
}
