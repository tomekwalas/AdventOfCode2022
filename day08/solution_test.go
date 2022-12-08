package day08

import "testing"

func TestPartOne(t *testing.T) {
	input := [][]int{
		{3, 0, 3, 7, 3},
		{2, 5, 5, 1, 2},
		{6, 5, 3, 3, 2},
		{3, 3, 5, 4, 9},
		{3, 5, 3, 9, 0},
	}

	expected := 21
	actual := PartOne(input)

	if expected != actual {
		t.Errorf("Wanted %v got %v", expected, actual)
	}
}

func TestPartTwo(t *testing.T) {
	input := [][]int{
		{3, 0, 3, 7, 3},
		{2, 5, 5, 1, 2},
		{6, 5, 3, 3, 2},
		{3, 3, 5, 4, 9},
		{3, 5, 3, 9, 0},
	}

	expected := 8
	actual := PartTwo(input)

	if expected != actual {
		t.Errorf("Wanted %v got %v", expected, actual)
	}
}
