package main

import (
	"testing"

	"github.com/bl1nk/adventofcode/x"
)

func TestPart1(t *testing.T) {
	in := x.ReadString("./testdata/example.txt")

	elves := foldCalories(in)
	actual, err := sumN(1, elves)
	if err != nil {
		t.Fatal(err)
	}

	expected := 24000
	if actual != expected {
		t.Errorf("expected %d calories, got %d", expected, actual)
	}
}

func TestPart2(t *testing.T) {
	in := x.ReadString("./testdata/example.txt")

	elves := foldCalories(in)
	actual, err := sumN(3, elves)
	if err != nil {
		t.Fatal(err)
	}

	expected := 45000
	if actual != expected {
		t.Errorf("expected %d calories, got %d", expected, actual)
	}
}
