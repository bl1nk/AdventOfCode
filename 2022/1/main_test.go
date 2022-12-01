package main

import (
	"testing"

	"github.com/bl1nk/adventofcode/x"
)

func TestMain(t *testing.T) {
	calories := x.ReadString("./testdata/example.txt")
	actual := mostCalories(calories)
	expected := 24000
	if actual != expected {
		t.Errorf("expected %d calories, got %d", expected, actual)
	}
}
