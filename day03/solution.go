package day03

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func Solution() {
	file, err := os.Open("./day03/input.txt")
	if err != nil {
		fmt.Println("Error while reading file")
		return
	}
	defer file.Close()

	input := []string{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	fmt.Println(PartOne(input))
	fmt.Println(PartTwo(input))
}

func PartOne(input []string) int {
	ci := []int{}
	for _, v := range input {
		mid := len(v) / 2
		f := v[:mid]
		l := v[mid:]

		for _, r := range f {
			if !strings.ContainsRune(l, r) {
				continue
			}
			if ok := matchRegexp("[a-z]", string(r)); ok {
				ci = append(ci, int(r)-96)
			}

			if ok := matchRegexp("[A-Z]", string(r)); ok {
				ci = append(ci, int(r)-64+26)
			}
			break
		}
	}
	sum := 0
	for _, v := range ci {
		sum += v
	}
	return sum
}

func PartTwo(input []string) int {
	ci := []int{}
	for i := 0; i < len(input)-2; i += 3 {
		f := input[i]
		m := input[i+1]
		l := input[i+2]
		for _, r := range f {
			if !strings.ContainsRune(m, r) || !strings.ContainsRune(l, r) {
				continue
			}

			if ok := matchRegexp("[a-z]", string(r)); ok {
				ci = append(ci, int(r)-96)
			}

			if ok := matchRegexp("[A-Z]", string(r)); ok {
				ci = append(ci, int(r)-64+26)
			}
			break
		}
	}
	sum := 0
	for _, v := range ci {
		sum += v
	}
	return sum
}

func matchRegexp(exp, value string) bool {
	match, err := regexp.MatchString(exp, value)
	if err != nil {
		return false
	}
	return match
}
