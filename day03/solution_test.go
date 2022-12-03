package day03

import "testing"

func TestPartOne(t *testing.T) {
	input := []string{"vJrwpWtwJgWrhcsFMMfFFhFp", "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL", "PmmdzqPrVvPwwTWBwg", "wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn", "ttgJtRGJQctTZtZT", "CrZsJsPPZsGzwwsLwLmpwMDw"}
	expected := 157
	actual := PartOne(input)

	if expected != actual {
		t.Errorf("Expected %v got %v", expected, actual)
	}
}

func TestPartTwo(t *testing.T) {
	input := []string{"vJrwpWtwJgWrhcsFMMfFFhFp", "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL", "PmmdzqPrVvPwwTWBwg", "wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn", "ttgJtRGJQctTZtZT", "CrZsJsPPZsGzwwsLwLmpwMDw"}
	expected := 70
	actual := PartTwo(input)

	if expected != actual {
		t.Errorf("Expected %v got %v", expected, actual)
	}
}
