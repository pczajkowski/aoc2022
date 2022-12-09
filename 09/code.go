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

type headTail struct {
	head point
	tail point
}

func tailInVicinity(tracker headTail) bool {
	return tracker.tail.x >= tracker.head.x-1 && tracker.tail.x <= tracker.head.x+1 && tracker.tail.y >= tracker.head.y-1 && tracker.tail.y <= tracker.head.y+1
}

func moveRight(tracker headTail, trail map[point]bool, steps int) (headTail, map[point]bool) {
	limit := tracker.head.x + steps
	for i := tracker.head.x + 1; i <= limit; i++ {
		tracker.head.x = i

		if !tailInVicinity(tracker) {
			tracker.tail.y = tracker.head.y
			tracker.tail.x = i - 1
			trail[tracker.tail] = true
		}
	}

	return tracker, trail
}

func moveLeft(tracker headTail, trail map[point]bool, steps int) (headTail, map[point]bool) {
	limit := tracker.head.x - steps
	for i := tracker.head.x - 1; i >= limit; i-- {
		tracker.head.x = i

		if !tailInVicinity(tracker) {
			tracker.tail.y = tracker.head.y
			tracker.tail.x = i + 1
			trail[tracker.tail] = true
		}
	}

	return tracker, trail
}

func moveUp(tracker headTail, trail map[point]bool, steps int) (headTail, map[point]bool) {
	limit := tracker.head.y + steps
	for i := tracker.head.y + 1; i <= limit; i++ {
		tracker.head.y = i

		if !tailInVicinity(tracker) {
			tracker.tail.x = tracker.head.x
			tracker.tail.y = i - 1
			trail[tracker.tail] = true
		}
	}

	return tracker, trail
}

func moveDown(tracker headTail, trail map[point]bool, steps int) (headTail, map[point]bool) {
	limit := tracker.head.y - steps
	for i := tracker.head.y - 1; i >= limit; i-- {
		tracker.head.y = i

		if !tailInVicinity(tracker) {
			tracker.tail.x = tracker.head.x
			tracker.tail.y = i + 1
			trail[tracker.tail] = true
		}
	}

	return tracker, trail
}

func drawTail(tracker headTail, action move, trail map[point]bool) (headTail, map[point]bool) {
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

func part1(moves []move) int {
	var tracker headTail
	trail := make(map[point]bool)
	trail[tracker.tail] = true

	for i := range moves {
		tracker, trail = drawTail(tracker, moves[i], trail)
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
