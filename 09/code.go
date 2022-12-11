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

type point struct {
	x int
	y int
}

func tailInVicinity(leader point, tail point) bool {
	return tail.x >= leader.x-1 && tail.x <= leader.x+1 && tail.y >= leader.y-1 && tail.y <= leader.y+1
}

func catchUp(leader point, tail point) point {
	if tailInVicinity(leader, tail) {
		return tail
	}

	if tail.x > leader.x {
		tail.x--
	} else if tail.x < leader.x {
		tail.x++
	}

	if tail.y > leader.y {
		tail.y--
	} else if tail.y < leader.y {
		tail.y++
	}

	return tail
}

func moveLead(leader point, action move) point {
	switch action.direction {
	case 'R':
		leader.x++
		return leader
	case 'L':
		leader.x--
		return leader
	case 'U':
		leader.y++
		return leader
	case 'D':
		leader.y--
		return leader
	}

	return leader
}

func wayOfTail(moves []move, snake []point) int {
	trail := make(map[point]bool)
	last := len(snake) - 1

	for i := range moves {
		for s := 0; s < moves[i].steps; s++ {
			for j := range snake {
				if j == 0 {
					snake[j] = moveLead(snake[j], moves[i])
				} else {
					snake[j] = catchUp(snake[j-1], snake[j])
				}
			}

			trail[snake[last]] = true
		}
	}

	return len(trail)
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
	snake1 := make([]point, 2)
	fmt.Println("Part1:", wayOfTail(moves, snake1))

	snake2 := make([]point, 10)
	fmt.Println("Part2:", wayOfTail(moves, snake2))
}
