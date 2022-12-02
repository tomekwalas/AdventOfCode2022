package day02

import "testing"

func TestPartOne(t *testing.T) {
	input := []string{"A Y", "B X", "C Z"}
	expected := 15

	result := PartOne(input)

	if expected != result {
		t.Errorf("Got %d want %d", result, expected)
	}
}

func TestPartTwo(t *testing.T) {
	input := []string{"A Y", "B X", "C Z"}
	expected := 12

	result := PartTwo(input)

	if expected != result {
		t.Errorf("Got %d want %d", result, expected)
	}
}
