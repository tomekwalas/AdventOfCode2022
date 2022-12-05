package day05

import "testing"

func TestInputParser(t *testing.T) {
	input := `
    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`
	parsed := parseInput(input)

	if len(parsed) != 4 {
		t.Errorf("Wanted %v got %v", 4, len(parsed))
	}

	if len(parsed[0]) != 2 {
		t.Errorf("Wanted 2 got %v", len(parsed[0]))
	}
	if len(parsed[1]) != 3 {
		t.Errorf("Wanted 2 got %v", len(parsed[1]))
	}
	if len(parsed[2]) != 1 {
		t.Errorf("Wanted 2 got %v", len(parsed[2]))
	}
	if len(parsed[3]) != 4 {
		t.Errorf("Wanted 4 got %v", len(parsed[3]))
	}

}

func TestPartOne(t *testing.T) {
	input := `
    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`
	expected := "CMZ"
	actual := PartOne(input)
	if expected != actual {
		t.Errorf("Wanted %v got %v", expected, actual)
	}
}

func TestPartTwo(t *testing.T) {
	input := `
    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`
	expected := "MCD"
	actual := PartTwo(input)
	if expected != actual {
		t.Errorf("Wanted %v got %v", expected, actual)
	}
}
