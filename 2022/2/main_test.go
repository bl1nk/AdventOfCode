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
	expected := 15
	actual := solvePart1(strings.TrimSpace(string(in)))

	if actual != expected {
		t.Fatalf("Expected %d, got %d", expected, actual)
	}
}
