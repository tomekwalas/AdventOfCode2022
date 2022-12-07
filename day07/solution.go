package day07

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type node struct {
	parent   *node
	children []*node
	isDir    bool
	name     string
	size     int
}

func (n *node) addFile(name string, size int) {
	nn := &node{isDir: false, name: name, size: size, parent: n}
	n.children = append(n.children, nn)
}

func (n *node) addFolder(name string) {
	nn := &node{isDir: true, name: name, parent: n}
	n.children = append(n.children, nn)
}

func (n *node) findFolder(name string) *node {
	if n == nil {
		return nil
	}

	for _, c := range n.children {
		if c.name == name {
			return c
		}
	}
	return nil
}

func Solution() {
	file, err := os.Open("./day07/input.txt")
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

func PartOne(instructions []string) int {
	dir := createDirectoryTree(instructions)
	f := &[]int{}
	getDirectorySizes(dir, f)

	out := 0
	for _, v := range *f {
		if v <= 100_000 {
			out += v
		}
	}

	return out
}

func PartTwo(instructions []string) int {
	dir := createDirectoryTree(instructions)
	f := &[]int{}
	t := getDirectorySizes(dir, f)
	total := 70_000_000
	required := 30_000_000

	out := t
	for _, v := range *f {
		if total-t+v > required && v < out {
			out = v
		}
	}

	return out
}

func createDirectoryTree(instructions []string) *node {
	var head *node
	var root *node
	for _, v := range instructions {
		if strings.Contains(v, "$ ls") {
			continue
		} else if strings.Contains(v, "$ cd") {
			var f string
			fmt.Sscanf(v, "$ cd %s", &f)
			if f == ".." {
				head = head.parent
				continue
			}

			if f == "/" {
				head = &node{name: f, isDir: true}
				root = head
				continue
			}

			if n := head.findFolder(f); n != nil {
				head = n
			}
		} else if strings.Contains(v, "dir") {
			var f string
			fmt.Sscanf(v, "dir %s", &f)
			head.addFolder(f)
		} else {
			var s int
			var n string
			fmt.Sscanf(v, "%d %s", &s, &n)
			head.addFile(n, s)
		}
	}

	return root
}

func getDirectorySizes(head *node, folders *[]int) int {
	out := 0
	for _, c := range head.children {
		if c.isDir {
			s := getDirectorySizes(c, folders)
			*folders = append(*folders, s)
			out += s
		} else {
			out += c.size
		}
	}
	return out
}
