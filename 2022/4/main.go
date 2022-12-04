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

	part2, err := solvePart2(strings.TrimSpace(string(in)))
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Printf("Part 2: One assignment pair overlaps another %d times\n", part2)
}

func solvePart1(in string) (int, error) {
	lines := strings.Split(in, "\n")
	var sum int
	for _, line := range lines {
		section1, section2, err := parseSections(line)
		if err != nil {
			return 0, err
		}

		if section1.contains(section2) || section2.contains(section1) {
			sum++
		}
	}

	return sum, nil
}

func solvePart2(in string) (int, error) {
	lines := strings.Split(in, "\n")
	var sum int
	for _, line := range lines {
		section1, section2, err := parseSections(line)
		if err != nil {
			return 0, err
		}

		if section1.overlaps(section2) {
			sum++
		}
	}

	return sum, nil
}

func parseSections(line string) (section, section, error) {
	pair := strings.Split(line, ",")

	s1 := strings.Split(pair[0], "-")
	from1, err := strconv.Atoi(s1[0])
	if err != nil {
		return section{}, section{}, err
	}

	to1, err := strconv.Atoi(s1[1])
	if err != nil {
		return section{}, section{}, err
	}

	s2 := strings.Split(pair[1], "-")
	from2, err := strconv.Atoi(s2[0])
	if err != nil {
		return section{}, section{}, err
	}

	to2, err := strconv.Atoi(s2[1])
	if err != nil {
		return section{}, section{}, err
	}

	return section{from: from1, to: to1}, section{from: from2, to: to2}, nil
}

type section struct {
	from int
	to   int
}

func (s section) contains(other section) bool {
	return s.from <= other.from && other.to <= s.to
}

func (s section) overlaps(other section) bool {
	return s.to >= other.from && s.from <= other.to
}
