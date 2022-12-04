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

func TestPart2(t *testing.T) {
	in, err := os.ReadFile("./testdata/example.txt")
	if err != nil {
		t.Fatal(err)
	}

	actual, err := solvePart2(strings.TrimSpace(string(in)))
	if err != nil {
		t.Fatal(err)
	}
	expected := 4

	if actual != expected {
		t.Errorf("expected %d, got %d", expected, actual)
	}

	for _, s := range []string{"2-4,6-8", "2-3,4-5"} {
		s1, s2, err := parseSections(s)
		if err != nil {
			t.Fatal(err)
		}

		if s1.overlaps(s2) || s2.overlaps(s1) {
			t.Errorf("s1 (%#v) and s2 (%#v) should not overlap", s1, s2)
		}
	}
	for _, s := range []string{"5-7,7-9", "2-8,3-7", "6-6,4-6", "2-6,4-8"} {
		s1, s2, err := parseSections(s)
		if err != nil {
			t.Fatal(err)
		}

		if !s1.overlaps(s2) || !s2.overlaps(s1) {
			t.Errorf("s1 (%#v) and s2 (%#v) should overlap", s1, s2)
		}

	}
}
