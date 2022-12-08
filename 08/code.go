package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readInput(file *os.File) [][]byte {
	scanner := bufio.NewScanner(file)
	var trees [][]byte

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		trees = append(trees, []byte(line))
	}

	return trees
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

	trees := readInput(file)
	fmt.Println(trees)
}
