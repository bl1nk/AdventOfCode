package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var re = regexp.MustCompile(`^move (\d+) from (\d+) to (\d+)$`)

func main() {
	in, err := os.ReadFile("./2022/5/input.txt")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	part1, err := solvePart1(string(in))
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Printf("Part 1: On top of each stack: %q\n", part1)

	part2, err := solvePart2(string(in))
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Printf("Part 2: On top of each stack: %q\n", part2)
}

func solvePart1(in string) (string, error) {
	state, moves, err := parseStateAndMoves(in)
	if err != nil {
		return "", err
	}
	for _, move := range moves {
		state = advance(state, move)
	}
	return state.solutionString(), nil
}
func solvePart2(in string) (string, error) {
	state, moves, err := parseStateAndMoves(in)
	if err != nil {
		return "", err
	}
	for _, move := range moves {
		state = advance9001(state, move)
	}
	return state.solutionString(), nil
}

func parseStateAndMoves(in string) (state, moves, error) {
	s := strings.Split(in, "\n\n")
	state := parseState(s[0])
	moves, err := parseMoves(s[1])
	return state, moves, err
}

type state struct {
	stacks []stack
}

func (s state) solutionString() string {
	var out string
	for _, stack := range s.stacks {
		out += strings.Trim(stack[len(stack)-1], "[]")
	}
	return out
}

type stack []string
type stacks []stack

type move struct {
	count, from, to int
}
type moves []move

func parseState(s string) state {
	lines := strings.Split(s, "\n")
	reverseSlice(lines)
	stacks := make(stacks, (len(lines[0])+3)/4)
	crateStrings := lines[1:]
	for _, line := range crateStrings {
		for i := 0; i < len(lines[0]); i += 4 {
			currentCrate := string(line[i:(i + 3)])
			if currentCrate == "   " {
				continue
			}
			stacks[i/4] = append(stacks[i/4], currentCrate)
		}
	}

	return state{stacks: stacks}
}

func reverseSlice(s []string) {
	sort.SliceStable(s, func(i, j int) bool { return i > j })
}

func parseMoves(s string) (moves, error) {
	ss := strings.Split(s, "\n")
	moves := make(moves, 0)

	for _, m := range ss {
		matches := re.FindStringSubmatch(m)
		// ???
		if len(matches) < 4 {
			continue
		}

		count, err := strconv.Atoi(matches[1])
		if err != nil {
			return nil, err
		}

		from, err := strconv.Atoi(matches[2])
		if err != nil {
			return nil, err
		}

		to, err := strconv.Atoi(matches[3])
		if err != nil {
			return nil, err
		}

		moves = append(moves, move{
			count: count,
			from:  from,
			to:    to,
		})
	}

	return moves, nil
}

func advance(s state, m move) state {
	for i := 0; i < m.count; i++ {
		fromStack := s.stacks[m.from-1]
		toStack := s.stacks[m.to-1]

		crate, fromStack := fromStack[len(fromStack)-1], fromStack[:len(fromStack)-1]
		toStack = append(toStack, crate)

		s.stacks[m.from-1] = fromStack
		s.stacks[m.to-1] = toStack
	}

	return s
}

func advance9001(s state, m move) state {
	fromStack := s.stacks[m.from-1]
	toStack := s.stacks[m.to-1]

	crate, fromStack := fromStack[len(fromStack)-m.count:], fromStack[:len(fromStack)-m.count]
	toStack = append(toStack, crate...)

	s.stacks[m.from-1] = fromStack
	s.stacks[m.to-1] = toStack
	return s
}
