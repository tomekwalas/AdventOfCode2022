package day04

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Solution() {
	file, err := os.Open("./day04/input.txt")
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
	i := 0

	for _, v := range input {
		pair := strings.Split(v, ",")
		elf1 := strings.Split(pair[0], "-")
		elf2 := strings.Split(pair[1], "-")
		range1 := []int{toInt(elf1[0]), toInt(elf1[1])}
		range2 := []int{toInt(elf2[0]), toInt(elf2[1])}

		if ok := inside(range1, range2); ok {
			i++
		}
	}

	return i
}

func PartTwo(input []string) int {
	i := 0

	for _, v := range input {
		pair := strings.Split(v, ",")
		elf1 := strings.Split(pair[0], "-")
		elf2 := strings.Split(pair[1], "-")
		range1 := []int{toInt(elf1[0]), toInt(elf1[1])}
		range2 := []int{toInt(elf2[0]), toInt(elf2[1])}

		if ok := intersect(range1, range2); ok {
			i++
		}
	}

	return i
}

func toInt(s string) int {
	out, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return out
}

func intersect(range1, range2 []int) bool {
	i1 := range2[0] >= range1[0] && range2[0] <= range1[1]
	i2 := range2[1] >= range1[0] && range2[1] <= range1[1]
	i3 := range1[0] >= range2[0] && range1[0] <= range2[1]
	i4 := range1[1] >= range2[0] && range1[1] <= range2[1]
	return i1 || i2 || i3 || i4
}

func inside(range1, range2 []int) bool {
	l1 := range1[1] - range1[0]
	l2 := range2[1] - range2[0]

	low := range1
	high := range2
	if l1 > l2 {
		low = range2
		high = range1
	}

	return low[0] >= high[0] && low[1] <= high[1]
}
