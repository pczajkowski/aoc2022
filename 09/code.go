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

func moveRight(x int, y int, trail [][]byte, steps int) (int, int, [][]byte) {
	x += steps
	edge := len(trail[y]) - 1
	if x > edge {
		trail[y] = append(trail[y], make([]byte, x-edge)...)
	}

	for i := x - 1; i > x-steps; i-- {
		trail[y][i] = '#'
	}

	return x, y, trail
}

func moveLeft(x int, y int, trail [][]byte, steps int) (int, int, [][]byte) {
	x -= steps

	if x < 0 {
		add := 0 - x
		x = 0
		trail[y] = append(make([]byte, add), trail[y]...)
	}

	for i := 1; i < steps; i++ {
		edge := len(trail[y]) - 1
		if edge < x+i {
			trail[y] = append(trail[y], make([]byte, x+i-edge)...)
		}

		trail[y][x+i] = '#'
	}

	return x, y, trail
}

func moveUp(x int, y int, trail [][]byte, steps int) (int, int, [][]byte) {
	width := len(trail[y])
	y += steps
	edge := len(trail) - 1

	if y > edge {
		for i := 0; i < y-edge; i++ {
			trail = append(trail, make([]byte, width))
		}
	}

	for i := y - steps + 1; i < y; i++ {
		edge := len(trail[i]) - 1
		if x > edge {
			trail[i] = append(trail[i], make([]byte, x-edge)...)
		}

		trail[i][x] = '#'
	}

	return x, y, trail
}

func moveDown(x int, y int, trail [][]byte, steps int) (int, int, [][]byte) {
	y -= steps

	if y < 0 {
		add := 0 - y
		y = 0
		width := len(trail[y])

		var toPrepend [][]byte
		for i := 0; i < add; i++ {
			toPrepend = append(toPrepend, make([]byte, width))
		}

		trail = append(toPrepend, trail...)
	}

	for i := 1; i < steps; i++ {
		edge := len(trail[y+i]) - 1
		if edge < x {
			trail[y+i] = append(trail[y+i], make([]byte, x-edge)...)
		}
		trail[y+i][x] = '#'
	}

	return x, y, trail
}

func drawTail(x int, y int, action move, trail [][]byte) (int, int, [][]byte) {
	switch action.direction {
	case 'R':
		return moveRight(x, y, trail, action.steps)
	case 'L':
		return moveLeft(x, y, trail, action.steps)
	case 'U':
		return moveUp(x, y, trail, action.steps)
	case 'D':
		return moveDown(x, y, trail, action.steps)
	}

	return x, y, trail
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
	trail := [][]byte{make([]byte, 1)}

	x := 0
	y := 0

	for i := range moves {
		x, y, trail = drawTail(x, y, moves[i], trail)
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
