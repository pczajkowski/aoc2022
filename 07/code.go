package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type item struct {
	name string
	size int
}

type dir struct {
	name   string
	parent *dir
	dirs   []dir
	files  []item
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
	read := false

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		if strings.HasPrefix(line, "$ cd") {
			read = false
			current = cd(line, current, dirsRead, &root)
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

				current.files = append(current.files, newFile)
			}
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
