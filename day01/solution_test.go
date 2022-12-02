package day01

import "testing"

func TestPartOne(t *testing.T) {
	list := []string{"10", "20", "", "50", "", "30", "40"}
	expected := 70
	result := PartOne(list)

	if result != expected {
		t.Errorf("Expected %d got %d", expected, result)
	}
}

func TestPartTwo(t *testing.T) {
	list := []string{"10", "20", "", "50", "", "30", "40", "", "5", "5", "10", "", "1", "1", "1"}
	expected := 150
	result := PartTwo(list)

	if result != expected {
		t.Errorf("Expected %d got %d", expected, result)
	}
}
