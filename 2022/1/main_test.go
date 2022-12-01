package main

import (
	"os"
	"testing"
)

func TestPart1(t *testing.T) {
	in, err := os.ReadFile("./testdata/example.txt")
	if err != nil {
		t.Fatal(err)
	}

	elves, err := foldCalories(string(in))
	if err != nil {
		t.Fatal(err)
	}

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
	in, err := os.ReadFile("./testdata/example.txt")
	if err != nil {
		t.Fatal(err)
	}

	elves, err := foldCalories(string(in))
	if err != nil {
		t.Fatal(err)
	}

	actual, err := sumN(3, elves)
	if err != nil {
		t.Fatal(err)
	}

	expected := 45000
	if actual != expected {
		t.Errorf("expected %d calories, got %d", expected, actual)
	}
}
