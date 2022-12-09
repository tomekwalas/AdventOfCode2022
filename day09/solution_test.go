package day09

import "testing"

func TestPartOne(t *testing.T) {
	input := []movement{
		{"R", 4},
		{"U", 4},
		{"L", 3},
		{"D", 1},
		{"R", 4},
		{"D", 1},
		{"L", 5},
		{"R", 2},
	}

	expected := 13
	actual := PartOne(input)

	if expected != actual {
		t.Errorf("Wanted %v got %v", expected, actual)
	}
}

func TestPartTwo(t *testing.T) {
	cases := []struct {
		input    []movement
		expected int
	}{
		{[]movement{
			{"R", 4},
			{"U", 4},
			{"L", 3},
			{"D", 1},
			{"R", 4},
			{"D", 1},
			{"L", 5},
			{"R", 2},
		},
			1,
		},
		{[]movement{
			{"R", 5},
			{"U", 8},
			{"L", 8},
			{"D", 3},
			{"R", 17},
			{"D", 10},
			{"L", 25},
			{"U", 20},
		},
			36,
		},
	}
	for _, c := range cases {
		actual := PartTwo(c.input)

		if c.expected != actual {
			t.Errorf("Wanted %v got %v", c.expected, actual)
		}
	}
}
