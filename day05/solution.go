package day05

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Solution() {
	file, err := os.ReadFile("./day05/input.txt")
	if err != nil {
		fmt.Println("Error while reading file")
		return
	}

	fmt.Println(PartOne(string(file)))
	fmt.Println(PartTwo(string(file)))
}

func PartOne(input string) string {
	parsed := parseInput(input)
	cmds := parsed[len(parsed)-1]
	parsed = parsed[:len(parsed)-1]
	stacks := createStacks(parsed)

	for _, v := range cmds {
		if v == "" {
			continue
		}

		count, origin, dest := interpretCmd(v)
		os := stacks[origin-1]
		ds := stacks[dest-1]

		for i := 0; i < count; i++ {
			ds.push(os.pop())
		}
	}

	r := ""
	for _, v := range stacks {
		r += v.peek()
	}

	return r
}

func PartTwo(input string) string {
	parsed := parseInput(input)
	cmds := parsed[len(parsed)-1]
	parsed = parsed[:len(parsed)-1]
	stacks := createStacks(parsed)

	for _, v := range cmds {
		if v == "" {
			continue
		}

		count, origin, dest := interpretCmd(v)
		os := stacks[origin-1]
		ds := stacks[dest-1]

		ss := make([]string, count, count)
		for i := count - 1; i >= 0; i-- {
			ss[i] = os.pop()
		}
		for _, v := range ss {
			ds.push(v)
		}

	}

	r := ""
	for _, v := range stacks {
		r += v.peek()
	}

	return r
}

func parseInput(input string) [][]string {
	out := strings.Split(input, "\n\n")
	stacks := strings.Split(out[0], "\n")
	result := [][]string{}
	for _, v := range stacks {
		itr := 0
		for i := 0; i < len(v)-2; i += 4 {
			if len(result) < itr+1 {
				result = append(result, []string{})
			}
			s := v[i : i+3]
			s = strings.TrimSpace(s)
			if s == "" || !strings.ContainsRune(s, '[') {
				itr++
				continue
			}

			s = strings.Replace(s, "[", "", 1)
			s = strings.Replace(s, "]", "", 1)

			result[itr] = append(result[itr], s)
			itr++

		}
	}

	cmds := strings.Split(out[1], "\n")
	result = append(result, cmds)
	return result
}

func createStacks(input [][]string) []*stack {
	stacks := []*stack{}
	for i := 0; i < len(input); i++ {
		s := &stack{}
		for j := len(input[i]) - 1; j >= 0; j-- {
			s.push(input[i][j])
		}
		stacks = append(stacks, s)
	}
	return stacks
}

func interpretCmd(cmd string) (int, int, int) {
	out := strings.Split(cmd, "from")
	cs := strings.Replace(out[0], "move", "", 1)
	cs = strings.TrimSpace(cs)

	count, err := strconv.Atoi(cs)
	if err != nil {
		return -1, -1, -1
	}

	sts := strings.Split(out[1], "to")
	ost := strings.TrimSpace(sts[0])
	origin, err := strconv.Atoi(ost)
	if err != nil {
		return -1, -1, -1
	}
	dst := strings.TrimSpace(sts[1])
	dest, err := strconv.Atoi(dst)
	if err != nil {
		return -1, -1, -1
	}

	return count, origin, dest
}
