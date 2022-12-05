package main

import (
	"os"
	"testing"
)

func TestReverseSlice(t *testing.T) {
	in := []string{"A", "B", "C"}
	reverseSlice(in)
	expected := []string{"C", "B", "A"}
	for idx := range in {
		if in[idx] != expected[idx] {
			t.Errorf("expected %q, got %q", expected[idx], in[idx])
		}
	}
}

func TestPart1(t *testing.T) {
	in, err := os.ReadFile("./testdata/example.txt")
	if err != nil {
		t.Fatal(err)
	}

	actual, err := solvePart1(string(in))
	if err != nil {
		t.Fatal(err)
	}
	expected := "CMZ"

	if actual != expected {
		t.Fatalf("expected %q, got %q", expected, actual)
	}
}

func TestPart2(t *testing.T) {
	in, err := os.ReadFile("./testdata/example.txt")
	if err != nil {
		t.Fatal(err)
	}

	actual, err := solvePart2(string(in))
	if err != nil {
		t.Fatal(err)
	}
	expected := "MCD"

	if actual != expected {
		t.Fatalf("expected %q, got %q", expected, actual)
	}
}
