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

func tailInVicinity(leader *point, tail point) bool {
	if leader == nil {
		return false
	}

	return tail.x >= leader.x-1 && tail.x <= leader.x+1 && tail.y >= leader.y-1 && tail.y <= leader.y+1
}

func moveRight(leader *point, tail point) point {
	if !tailInVicinity(leader, tail) {
		if leader != nil {
			tail.y = leader.y
			tail.x = leader.x - 1
		} else {
			tail.x++
		}
	}

	return tail
}

func moveLeft(leader *point, tail point) point {
	if !tailInVicinity(leader, tail) {
		if leader != nil {
			tail.y = leader.y
			tail.x = leader.x + 1
		} else {
			tail.x--
		}
	}

	return tail
}

func moveUp(leader *point, tail point) point {
	if !tailInVicinity(leader, tail) {
		if leader != nil {
			tail.x = leader.x
			tail.y = leader.y - 1
		} else {
			tail.y++
		}
	}

	return tail
}

func moveDown(leader *point, tail point) point {
	if !tailInVicinity(leader, tail) {
		if leader != nil {
			tail.x = leader.x
			tail.y = leader.y + 1
		} else {
			tail.y--
		}
	}

	return tail
}

func drawTail(leader *point, tail point, action move) point {
	switch action.direction {
	case 'R':
		return moveRight(leader, tail)
	case 'L':
		return moveLeft(leader, tail)
	case 'U':
		return moveUp(leader, tail)
	case 'D':
		return moveDown(leader, tail)
	}

	return tail
}

func part1(moves []move) int {
	parts := make([]point, 2)

	trail := make(map[point]bool)
	last := len(parts) - 1

	for i := range moves {
		for s := 0; s < moves[i].steps; s++ {
			for j := range parts {
				if j == 0 {
					parts[j] = drawTail(nil, parts[j], moves[i])
				} else {
					parts[j] = drawTail(&parts[j-1], parts[j], moves[i])
				}

				if j == last {
					trail[parts[j]] = true
				}
			}
			fmt.Println(parts)
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
	fmt.Println("Part1:", part1(moves))
}
