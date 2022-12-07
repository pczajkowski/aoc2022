package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

type item struct {
	name string
	size int
}

type dir struct {
	name   string
	parent *dir
	dirs   []*dir
	files  []*item
}

func parentHasDir(parent *dir, name string) *dir {
	for i := range parent.dirs {
		if parent.dirs[i].name == name {
			return parent.dirs[i]
		}
	}

	return nil
}

func cd(line string, current *dir, root *dir) *dir {
	var name string
	n, err := fmt.Sscanf(line, "$ cd %s", &name)
	if n != 1 || err != nil {
		log.Fatal("Can't parse cd:", err)
	}

	if name == "/" {
		current = root
	} else if name == ".." {
		current = current.parent
	} else {
		parent := current
		current = parentHasDir(parent, name)
		if current == nil {
			newDir := dir{name: name, parent: parent}
			parent.dirs = append(parent.dirs, &newDir)
			current = &newDir
		}
	}

	return current
}

func readInput(file *os.File) dir {
	scanner := bufio.NewScanner(file)
	root := dir{name: "/"}
	var current *dir
	read := false

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		if strings.HasPrefix(line, "$ cd") {
			read = false
			current = cd(line, current, &root)
		} else if strings.HasPrefix(line, "$ ls") {
			read = true
			continue
		}

		if read {
			if strings.HasPrefix(line, "dir ") {
				continue
			} else {
				var newFile item
				n, err := fmt.Sscanf(line, "%d %s", &newFile.size, &newFile.name)
				if n != 2 || err != nil {
					log.Fatal("Can't parse cd:", err)
				}

				current.files = append(current.files, &newFile)
			}
		}

	}

	return root
}

func getSizes(root dir, sizes []int) (int, []int) {
	size := 0
	for i := range root.files {
		size += root.files[i].size
	}

	for i := range root.dirs {
		var c int
		c, sizes = getSizes(*root.dirs[i], sizes)
		size += c
	}

	sizes = append(sizes, size)
	return size, sizes
}

func part1(sizes []int) int {
	sum := 0
	for i := range sizes {
		if sizes[i] < 100000 {
			sum += sizes[i]
		}
	}

	return sum
}

func part2(sizes []int, total int, fsSize int, needed int) int {
	unused := fsSize - total
	needed -= unused

	sort.Ints(sizes)
	for i := range sizes {
		if sizes[i] >= needed {
			return sizes[i]
		}
	}

	return 0
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("You need to specify a file!")
	}

	filePath := os.Args[1]
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Failed to open %s!\n", filePath)

	}

	root := readInput(file)
	total, sizes := getSizes(root, []int{})
	fmt.Println("Part1:", part1(sizes))
	fmt.Println("Part2:", part2(sizes, total, 70000000, 30000000))
}
