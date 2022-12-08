package day08

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Solution() {
	file, err := os.Open("./day08/input.txt")
	if err != nil {
		fmt.Println("Error while reading file")
		return
	}
	defer file.Close()

	input := [][]int{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		i := []int{}
		s := strings.Split(scanner.Text(), "")
		for _, ss := range s {
			n, _ := strconv.Atoi(ss)
			i = append(i, n)
		}
		input = append(input, i)
	}

	fmt.Println(PartOne(input))
	fmt.Println(PartTwo(input))
}

func PartOne(input [][]int) int {
	dirs := [][]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	}

	v := 0
	for i, c := range input {
		for j, t := range c {
			b := 0
			for _, d := range dirs {
				x, y := i+d[0], j+d[1]
				for x < len(input) && x > -1 && y < len(input[i]) && y > -1 {
					tt := input[x][y]
					if tt >= t {
						b++
						break
					}
					x, y = x+d[0], y+d[1]
				}
			}
			if b < 4 {
				v++
			}
		}
	}
	return v
}

func PartTwo(input [][]int) int {
	dirs := [][]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	}

	v := 0
	for i, c := range input {
		for j, t := range c {
			b := 1
			for _, d := range dirs {
				x, y := i+d[0], j+d[1]
				s := 0
				for x < len(input) && x > -1 && y < len(input[i]) && y > -1 {
					tt := input[x][y]
					s++
					if tt >= t {
						break
					}
					x, y = x+d[0], y+d[1]
				}
				b *= s
			}
			if v < b {
				v = b
			}
		}
	}
	return v
}
