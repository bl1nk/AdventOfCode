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

	actual, err := solvePart1(strings.TrimSpace(string(in)))
	if err != nil {
		t.Fatal(err)
	}
	expected := 2

	if actual != expected {
		t.Fatalf("expected %d, got %d", expected, actual)
	}
}
