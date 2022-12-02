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

var strats = map[string]strat{
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
	fmt.Printf("Total score according to strategy guide: %d", part1)
}

func solvePart1(in string) int {
	total := 0
	for _, round := range strings.Split(in, "\n") {
		choice := round[2]
		total += strats[round]
		total += pointsForShape[shapes[choice]]
	}

	return total

}
