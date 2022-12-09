package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type move struct {
	direction byte
	steps     int
}

func readInput(file *os.File) []move {
	scanner := bufio.NewScanner(file)
	var moves []move

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		var current move
		n, err := fmt.Sscanf(line, "%c %d", &current.direction, &current.steps)
		if n != 2 || err != nil {
			log.Fatal("Can't parse cd:", err)
		}

		moves = append(moves, current)
	}

	return moves
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

	moves := readInput(file)
	fmt.Println(moves)
}
