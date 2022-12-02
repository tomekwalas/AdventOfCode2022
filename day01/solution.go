package day01

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func Solution() {
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Error while reading file")
		return
	}
	defer file.Close()

	list := []string{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		list = append(list, scanner.Text())
	}

	fmt.Println(PartOne(list))
	fmt.Println(PartTwo(list))
}

func PartOne(list []string) int {
	elfs := getElfsWithSnacks(list)
	return elfs[len(elfs)-1]
}

func PartTwo(list []string) int {
	elfs := getElfsWithSnacks(list)
	top := elfs[len(elfs)-3:]
	sum := 0

	for _, v := range top {
		sum += v
	}

	return sum
}

func getElfsWithSnacks(list []string) []int {
	elfs := []int{0}
	idx := 0
	for _, v := range list {
		if v == "" {
			elfs = append(elfs, 0)
			idx++
			continue
		}

		number, err := strconv.Atoi(v)
		if err != nil {
			continue
		}
		elfs[idx] += number

	}

	sort.Ints(elfs)
	return elfs

}
