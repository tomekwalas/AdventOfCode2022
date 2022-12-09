package day09

import (
	"bufio"
	"fmt"
	"os"
)

type movement struct {
	direction string
	steps     int
}

type position struct {
	x int
	y int
}

func Solution() {
	file, err := os.Open("./day09/input.txt")
	if err != nil {
		fmt.Println("Error while reading file")
		return
	}
	defer file.Close()

	input := []movement{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		m := movement{}
		fmt.Sscanf(scanner.Text(), "%s %d", &m.direction, &m.steps)
		input = append(input, m)
	}

	fmt.Println(PartOne(input))
	fmt.Println(PartTwo(input))
}

func PartOne(movements []movement) int {
	head, tail := position{}, position{}
	sdirs := map[string][]int{
		"R": {1, 0},
		"L": {-1, 0},
		"U": {0, 1},
		"D": {0, -1},
	}
	dirs := map[string][]int{
		"R":  {1, 0},
		"L":  {-1, 0},
		"U":  {0, 1},
		"D":  {0, -1},
		"UL": {-1, 1},
		"UR": {1, 1},
		"DR": {1, -1},
		"DL": {-1, -1},
	}
	path := map[[2]int]bool{{0, 0}: true}
	for _, m := range movements {
		dir := dirs[m.direction]
		for i := 0; i < m.steps; i++ {
			head.x += dir[0]
			head.y += dir[1]
			if ok := isInRange(head, tail, dirs); ok {
				continue
			}
			for _, d := range dirs {
				nt := position{tail.x + d[0], tail.y + d[1]}
				if ok := isInRange(head, nt, sdirs); ok {
					tail = nt
					break
				}
			}
			path[[2]int{tail.x, tail.y}] = true
		}
	}

	return len(path)
}

func PartTwo(movements []movement) int {
	rope := make([]position, 10, 10)
	dirs := map[string][]int{
		"R":  {1, 0},
		"L":  {-1, 0},
		"U":  {0, 1},
		"D":  {0, -1},
		"UL": {-1, 1},
		"UR": {1, 1},
		"DR": {1, -1},
		"DL": {-1, -1},
	}
	sdirs := map[string][]int{
		"R": {1, 0},
		"L": {-1, 0},
		"U": {0, 1},
		"D": {0, -1},
	}
	path := map[[2]int]bool{{0, 0}: true}
	for _, m := range movements {
		dir := dirs[m.direction]
		for i := 0; i < m.steps; i++ {
			rope[0].x += dir[0]
			rope[0].y += dir[1]
			for j := 0; j < len(rope)-1; j++ {
				head := rope[j]
				tail := rope[j+1]
				if ok := isInRange(head, tail, dirs); ok {
					continue
				}
				f := false
				for _, d := range dirs {
					nt := position{tail.x + d[0], tail.y + d[1]}
					if ok := isInRange(head, nt, sdirs); ok {
						f = true
						rope[j+1] = nt
						break
					}
				}
				if !f {
					for _, d := range dirs {
						nt := position{tail.x + d[0], tail.y + d[1]}
						if ok := isInRange(head, nt, dirs); ok {
							rope[j+1] = nt
							break
						}
					}
				}
				if j+1 == len(rope)-1 {
					path[[2]int{rope[j+1].x, rope[j+1].y}] = true
				}
			}
		}
	}

	return len(path)
}

func isInRange(head, tail position, dirs map[string][]int) bool {
	for _, d := range dirs {
		hx := tail.x + d[0]
		hy := tail.y + d[1]
		inRange := head.x == hx && head.y == hy
		overlap := head.x == tail.x && head.y == tail.y
		if inRange || overlap {
			return true
		}
	}
	return false
}
