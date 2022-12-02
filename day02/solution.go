package day02

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Solution() {
	file, err := os.Open("./day02/input.txt")
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
	signScore := map[string]int{"X": 1, "Y": 2, "Z": 3}
	resultScore := map[int]int{1: 6, 0: 3, -1: 0}

	rounds := []int{}
	for _, v := range input {
		round := strings.Split(v, " ")
		out := determineOutcome(round[0], round[1])
		rounds = append(rounds, resultScore[out]+signScore[round[1]])
	}

	sum := 0
	for _, v := range rounds {
		sum += v
	}
	return sum
}

func PartTwo(input []string) int {
	wm := map[string]int{"X": -1, "Y": 0, "Z": 1}
	ss := map[string]int{"A": 1, "B": 2, "C": 3}
	rs := map[int]int{1: 6, 0: 3, -1: 0}
	ws := map[string]string{"C": "A", "A": "B", "B": "C"}
	ls := map[string]string{"A": "C", "B": "A", "C": "B"}

	rounds := []int{}
	for _, v := range input {
		round := strings.Split(v, " ")
		out := wm[round[1]]
		if out == 0 {
			rounds = append(rounds, rs[out]+ss[round[0]])
		} else if out == 1 {
			s := ws[round[0]]
			rounds = append(rounds, rs[out]+ss[s])
		} else {
			s := ls[round[0]]
			rounds = append(rounds, rs[out]+ss[s])
		}
	}
	sum := 0
	for _, v := range rounds {
		sum += v
	}
	return sum
}

func determineOutcome(so, sm string) int {
	signsMap := map[string]string{"X": "A", "Y": "B", "Z": "C"}
	winSigns := map[string]string{"A": "C", "B": "A", "C": "B"}
	normalizedSign := signsMap[sm]
	if so == normalizedSign {
		return 0
	}

	if winSigns[normalizedSign] == so {
		return 1
	}

	return -1
}
