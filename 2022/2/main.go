package main

import (
	"fmt"
	"os"
	"strings"
)

type shape = string

const (
	rock     shape = "rock"
	paper    shape = "paper"
	scissors shape = "scissors"
)

type strat = int

const (
	win  strat = 6
	draw strat = 3
	lose strat = 0
)

var stratsPart1 = map[string]strat{
	"A X": draw,
	"A Y": win,
	"A Z": lose,
	"B X": lose,
	"B Y": draw,
	"B Z": win,
	"C X": win,
	"C Y": lose,
	"C Z": draw,
}

var stratsPart2 = map[string]strat{
	"A X": lose,
	"B X": lose,
	"C X": lose,
	"A Y": draw,
	"B Y": draw,
	"C Y": draw,
	"A Z": win,
	"B Z": win,
	"C Z": win,
}

var pointsForShape = map[shape]int{
	rock:     1,
	paper:    2,
	scissors: 3,
}

var shapes = map[byte]shape{
	'A': rock,
	'X': rock,
	'B': paper,
	'Y': paper,
	'C': scissors,
	'Z': scissors,
}

func main() {
	in, err := os.ReadFile("./2022/2/input.txt")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	part1 := solvePart1(strings.TrimSpace(string(in)))
	fmt.Printf("Total score according to strategy guide from part 1: %d\n", part1)

	part2 := solvePart2(strings.TrimSpace(string(in)))
	fmt.Printf("Total score according to strategy guide from part 2: %d\n", part2)
}

func solvePart1(in string) int {
	total := 0
	for _, round := range strings.Split(in, "\n") {
		choice := round[2]
		total += stratsPart1[round]
		total += pointsForShape[shapes[choice]]
	}

	return total
}

func solvePart2(in string) int {
	total := 0
	for _, round := range strings.Split(in, "\n") {
		choice := pickShape(round)
		total += stratsPart2[round]
		total += pointsForShape[choice]
	}

	return total
}

func pickShape(round string) shape {
	theirs := shapes[round[0]]
	strat := stratsPart2[round]
	switch strat {
	case draw:
		return theirs
	case win:
		switch theirs {
		case rock:
			return paper
		case paper:
			return scissors
		case scissors:
			return rock
		}
	case lose:
		switch theirs {
		case rock:
			return scissors
		case paper:
			return rock
		case scissors:
			return paper
		}
	}
	panic("???")
}
