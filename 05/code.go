package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type move struct {
	what int
	from int
	to   int
}

func parseBoxes(line string, boxes [][]byte) [][]byte {
	counter := 0
	index := 0
	readLetter := false

	for i := range line {
		if line[i] == '\n' {
			break
		}

		counter++
		if counter > 3 {
			index++
			counter = 0
		}

		if line[i] == '[' {
			readLetter = true
			continue
		}

		if line[i] == ']' {
			readLetter = false
			continue
		}

		if readLetter {
			edge := len(boxes) - 1
			if index > edge {
				for i := 0; i < index-edge; i++ {
					boxes = append(boxes, []byte{})
				}
			}

			boxes[index] = append(boxes[index], line[i])
		}
	}

	return boxes
}

func parseMove(line string) move {
	var current move
	n, err := fmt.Sscanf(line, "move %d from %d to %d", &current.what, &current.from, &current.to)
	if n != 3 || err != nil {
		log.Fatal("Can't parse move!", err)
	}

	return current
}

func readInput(file *os.File) ([][]byte, []move) {
	scanner := bufio.NewScanner(file)
	var moves []move
	var boxes [][]byte
	readBoxes := true

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			readBoxes = false
			continue
		}

		if readBoxes {
			boxes = parseBoxes(line, boxes)
			continue
		}

		moves = append(moves, parseMove(line))
	}

	return boxes, moves
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

	boxes, moves := readInput(file)
	fmt.Println(boxes, moves)
}
