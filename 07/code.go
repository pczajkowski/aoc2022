package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type file struct {
	name string
	size int
}

type dir struct {
	name   string
	parent *dir
	dirs   []dir
	files  []file
}

func cd(line string, current *dir, dirsRead map[string]*dir, root *dir) *dir {
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
		current, _ = dirsRead[name]
		if current == nil {
			newDir := dir{name: name, parent: parent}
			parent.dirs = append(parent.dirs, newDir)
			current = &newDir
			fmt.Println(name, current, newDir)
			dirsRead[name] = current
		}
	}

	return current
}

func readInput(file *os.File) dir {
	scanner := bufio.NewScanner(file)
	root := dir{name: "/"}
	dirsRead := make(map[string]*dir)
	var current *dir

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		if strings.HasPrefix(line, "$ cd") {
			current = cd(line, current, dirsRead, &root)
			fmt.Println(line, current)
		}
	}

	return root
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
	fmt.Println(root)
}
