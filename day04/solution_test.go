package day04

import "testing"

func TestPartOne(t *testing.T) {
	input := []string{"2-4,6-8", "2-3,4-5", "5-7,7-9", "2-8,3-7", "6-6,4-6", "2-6,4-8"}
	expected := 2
	actual := PartOne(input)

	if expected != actual {
		t.Errorf("Expected %v got %v", expected, actual)
	}
}

func TestPartTwo(t *testing.T) {
	input := []string{"2-4,6-8", "2-3,4-5", "5-7,7-9", "2-8,3-7", "6-6,4-6", "2-6,4-8"}
	expected := 4
	actual := PartTwo(input)

	if expected != actual {
		t.Errorf("Expected %v got %v", expected, actual)
	}
}
