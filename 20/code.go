package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readInput(file *os.File) []int {
	scanner := bufio.NewScanner(file)
	var numbers []int

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		var current int
		n, err := fmt.Sscanf(line, "%d", &current)
		if n != 1 || err != nil {
			log.Fatal("Can't parse:", line, err)
		}

		numbers = append(numbers, current)
	}

	return numbers
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

	numbers := readInput(file)
	fmt.Println(numbers)
}
