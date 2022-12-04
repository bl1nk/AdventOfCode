package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	in, err := os.ReadFile("./2022/4/input.txt")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	part1, err := solvePart1(strings.TrimSpace(string(in)))
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Printf("Part 1: One assignment pair fully contains the other %d times\n", part1)
}

func solvePart1(in string) (int, error) {
	lines := strings.Split(in, "\n")
	var sum int
	for _, line := range lines {
		sections := strings.Split(line, ",")

		section1, err := parseSections(sections[0])
		if err != nil {
			return 0, err
		}

		section2, err := parseSections(sections[1])
		if err != nil {
			return 0, err
		}

		if section1.contains(section2) || section2.contains(section1) {
			sum++
		}
	}

	return sum, nil
}

func parseSections(pair string) (section, error) {
	boundaries := strings.Split(pair, "-")

	from, err := strconv.Atoi(boundaries[0])
	if err != nil {
		return section{}, err
	}

	to, err := strconv.Atoi(boundaries[1])
	if err != nil {
		return section{}, err
	}

	return section{from: from, to: to}, nil
}

type section struct {
	from int
	to   int
}

func (s *section) contains(other section) bool {
	return s.from <= other.from && other.to <= s.to
}
