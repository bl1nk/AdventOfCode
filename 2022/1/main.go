package main

import (
	"fmt"
	"strings"

	"github.com/bl1nk/adventofcode/x"
)

func main() {
	input := x.ReadString("./input.txt")
	fmt.Println(mostCalories(input))
}

func mostCalories(input string) int {
	lines := strings.Split(input, "\n")
	var most int
	var sum int
	for _, line := range lines {
		if line == "" {
			if most < sum {
				most = sum
			}
			sum = 0
		} else {
			i := x.MustAtoi(line)
			sum += i
		}
	}
	return most
}
