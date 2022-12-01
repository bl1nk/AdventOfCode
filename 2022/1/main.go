package main

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/bl1nk/adventofcode/x"
)

func main() {
	input := x.ReadString("./2022/1/input.txt")
	elves := foldCalories(input)

	mostCalories, err := sumN(1, elves)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Printf("Part 1: The elf carrying the most calories carries %d calories.\n", mostCalories)

	topThree, err := sumN(3, elves)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Printf("Part 2: The three elves carrying the most calories carry %d calories.\n", topThree)

}

func foldCalories(in string) []int {
	var elf int
	elves := []int{0}
	for _, line := range strings.Split(strings.TrimSpace(in), "\n") {
		if line == "" {
			elf++
			elves = append(elves, 0)
			continue
		}

		cal := x.MustAtoi(line)
		elves[elf] += cal
	}

	sort.Ints(elves)
	return elves
}

func sumN(n int, ints []int) (int, error) {
	if n > len(ints) {
		return 0, fmt.Errorf("got a list of %d ints, cannot sum %d elements", len(ints), n)
	}
	toSum := ints[len(ints)-n:]

	s := 0
	for _, i := range toSum {
		s += i
	}
	return s, nil

}
