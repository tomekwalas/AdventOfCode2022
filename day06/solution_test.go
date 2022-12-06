package day06

import "testing"

func TestPartOne(t *testing.T) {
	cases := []struct {
		input    string
		expected int
	}{
		{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", 7},
		{"bvwbjplbgvbhsrlpgdmjqwftvncz", 5},
		{"nppdvjthqldpwncqszvftbrmjlhg", 6},
		{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 10},
		{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 11},
	}

	for _, c := range cases {
		actual := PartOne(c.input)

		if c.expected != actual {
			t.Errorf("Wanted %v got %v", c.expected, actual)
		}
	}
}

func TestPartTwo(t *testing.T) {
	cases := []struct {
		input    string
		expected int
	}{
		{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", 19},
		{"bvwbjplbgvbhsrlpgdmjqwftvncz", 23},
		{"nppdvjthqldpwncqszvftbrmjlhg", 23},
		{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 29},
		{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 26},
	}

	for _, c := range cases {
		actual := PartTwo(c.input)

		if c.expected != actual {
			t.Errorf("Wanted %v got %v", c.expected, actual)
		}
	}
}
