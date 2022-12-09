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

type headTail struct {
	headX, headY int
	tailX, tailY int
}

func tailInVicinity(tracker headTail) bool {
	return tracker.tailX >= tracker.headX-1 && tracker.tailX <= tracker.headX+1 && tracker.tailY >= tracker.headY-1 && tracker.tailY <= tracker.headY+1
}

func moveRight(tracker headTail, trail [][]byte, steps int) (headTail, [][]byte) {
	destination := tracker.headX + steps
	edge := len(trail[tracker.headY]) - 1
	if destination > edge {
		trail[tracker.headY] = append(trail[tracker.headY], make([]byte, destination-edge)...)
	}

	for i := tracker.headX + 1; i <= destination; i++ {
		tracker.headX = i

		if !tailInVicinity(tracker) {
			tracker.tailY = tracker.headY
			tracker.tailX = i - 1
			trail[tracker.tailY][tracker.tailX] = '#'
		}
	}

	return tracker, trail
}

func moveLeft(tracker headTail, trail [][]byte, steps int) (headTail, [][]byte) {
	destination := tracker.headX - steps

	if destination < 0 {
		add := 0 - destination
		destination = 0
		tracker.headX = add
		trail[tracker.headY] = append(make([]byte, add), trail[tracker.headY]...)
	}

	for i := tracker.headX - 1; i >= destination; i-- {
		tracker.headX = i

		if !tailInVicinity(tracker) {
			edge := len(trail[tracker.headY]) - 1
			if edge < i+1 {
				trail[tracker.headY] = append(trail[tracker.headY], make([]byte, i-edge+1)...)
			}

			tracker.tailY = tracker.headY
			tracker.tailX = i + 1
			trail[tracker.tailY][tracker.tailX] = '#'
		}
	}

	return tracker, trail
}

func moveUp(tracker headTail, trail [][]byte, steps int) (headTail, [][]byte) {
	destination := tracker.headY + steps

	edge := len(trail) - 1
	if destination > edge {
		width := len(trail[tracker.headY])
		for i := 0; i < destination-edge; i++ {
			trail = append(trail, make([]byte, width))
		}
	}

	for i := tracker.headY + 1; i <= destination; i++ {
		tracker.headY = i

		edge := len(trail[i]) - 1
		if tracker.headX > edge {
			trail[i] = append(trail[i], make([]byte, tracker.headX-edge)...)
		}

		if !tailInVicinity(tracker) {
			tracker.tailX = tracker.headX
			tracker.tailY = i - 1
			trail[tracker.tailY][tracker.tailX] = '#'
		}
	}

	return tracker, trail
}

func moveDown(tracker headTail, trail [][]byte, steps int) (headTail, [][]byte) {
	destination := tracker.headY - steps
	if destination < 0 {
		add := 0 - destination
		add++
		destination = 0
		tracker.headY = add
		width := len(trail[tracker.headY])

		var toPrepend [][]byte
		for i := 0; i < add; i++ {
			toPrepend = append(toPrepend, make([]byte, width))
		}

		trail = append(toPrepend, trail...)
	}

	for i := tracker.headY - 1; i >= destination; i-- {
		tracker.headY = i

		if !tailInVicinity(tracker) {
			tracker.tailX = tracker.headX
			tracker.tailY = i + 1

			edge := len(trail[tracker.tailY]) - 1
			if edge < tracker.tailX {
				trail[tracker.tailY] = append(trail[tracker.tailY], make([]byte, tracker.tailX-edge)...)
			}
			trail[tracker.tailY][tracker.tailX] = '#'
		}
	}

	return tracker, trail
}

func drawTail(tracker headTail, action move, trail [][]byte) (headTail, [][]byte) {
	switch action.direction {
	case 'R':
		return moveRight(tracker, trail, action.steps)
	case 'L':
		return moveLeft(tracker, trail, action.steps)
	case 'U':
		return moveUp(tracker, trail, action.steps)
	case 'D':
		return moveDown(tracker, trail, action.steps)
	}

	return tracker, trail
}

func calculate(trail [][]byte) int {
	count := 0
	for i := range trail {
		for j := range trail[i] {
			if trail[i][j] == '#' {
				count++
			}
		}
	}

	return count
}

func part1(moves []move) int {
	trail := [][]byte{[]byte{'#'}}
	var tracker headTail

	for i := range moves {
		tracker, trail = drawTail(tracker, moves[i], trail)
	}

	return calculate(trail)
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
