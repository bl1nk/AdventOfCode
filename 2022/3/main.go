package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	in, err := os.ReadFile("./2022/3/input.txt")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// Find the item type that appears in both compartments of each
	// rucksack. What is the sum of the priorities of those item types?
	part1 := solvePart1(strings.TrimSpace(string(in)))
	fmt.Printf("The sum of the priorities of items that appear in both compartments of each rucksack is: %d\n", part1)

	part2 := solvePart2(strings.TrimSpace(string(in)))
	fmt.Printf("The sum of the priorities of badges of each three-elf groups is: %d\n", part2)
}

func solvePart1(in string) int {
	var sum int
	for _, rucksack := range strings.Split(in, "\n") {
		sum += sumPriorities(rucksack)
	}
	return sum
}

func solvePart2(in string) int {
	var sum int

	rucksacks := strings.Split(in, "\n")
	for idx := 0; idx < len(rucksacks); idx = idx + 3 {
		for _, item := range rucksacks[idx] {
			if strings.ContainsRune(rucksacks[idx+1], item) && strings.ContainsRune(rucksacks[idx+2], item) {
				sum += priority(item)
				break
			}
		}
	}

	return sum
}

func sumPriorities(rucksack string) int {
	sum := 0
	half := len(rucksack) / 2

	counts := make([]map[rune]int, 2)
	counts[0] = make(map[rune]int)
	counts[1] = make(map[rune]int)

	for _, item := range rucksack[half:] {
		counts[0][item] = priority(item)
	}

	for _, item := range rucksack[:half] {
		counts[1][item] = priority(item)
	}

	found := make(map[rune]struct{})
	for first := range counts[0] {
		for second := range counts[1] {
			if first == second {
				found[first] = struct{}{}
			}
		}
	}

	for k := range found {
		sum += priority(k)
	}

	return sum
}

func priority(r rune) int {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	return strings.Index(letters, string(r)) + 1
}
