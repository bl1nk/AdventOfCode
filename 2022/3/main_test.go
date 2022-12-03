package main

import (
	"os"
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	in, err := os.ReadFile("./testdata/example.txt")
	if err != nil {
		t.Fatal(err)
	}

	expected := 157
	actual := solvePart1(strings.TrimSpace(string(in)))

	if actual != expected {
		t.Fatalf("expected %d, got %d", expected, actual)
	}
}

func TestPart2(t *testing.T) {
	in, err := os.ReadFile("./testdata/example.txt")
	if err != nil {
		t.Fatal(err)
	}

	expected := 70
	actual := solvePart2(strings.TrimSpace(string(in)))

	if actual != expected {
		t.Fatalf("expected %d, got %d", expected, actual)
	}
}
