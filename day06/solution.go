package day06

import (
	"fmt"
	"os"
)

func Solution() {
	file, err := os.ReadFile("./day06/input.txt")
	if err != nil {
		fmt.Println("Error while reading file")
		return
	}

	fmt.Println(PartOne(string(file)))
	fmt.Println(PartTwo(string(file)))
}
func PartOne(input string) int {
	return getPacketMarker(input, 4)
}

func PartTwo(input string) int {
	return getPacketMarker(input, 14)
}

func getPacketMarker(input string, chunk int) int {
	m := map[string]int{}
	i := 0
	for i < len(input) {
		cc := string(input[i])
		if pp, ok := m[cc]; ok {
			m = map[string]int{}
			i = pp + 1
			continue
		}

		if len(m) == chunk-1 {
			return i + 1
		}

		m[cc] = i
		i++
	}
	return 0
}
