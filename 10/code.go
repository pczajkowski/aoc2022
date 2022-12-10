package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type op byte

const (
	addx op = iota
	noop
)

type command struct {
	instruction op
	value       int
}

func readInput(file *os.File) []command {
	scanner := bufio.NewScanner(file)
	var commands []command

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		var current command
		if line[0] == 'n' {
			current.instruction = noop
		} else {
			current.instruction = addx
			n, err := fmt.Sscanf(line, "addx %d", &current.value)
			if n != 1 || err != nil {
				log.Fatal("Can't parse input:", err, line)
			}
		}

		commands = append(commands, current)
	}

	return commands
}

func part1(commands []command) int {
	cycles := 0
	currentLimit := 20
	absoluteLimit := 221
	result := 0
	x := 1

	for i := range commands {
		if cycles > absoluteLimit {
			break
		}

		if commands[i].instruction == noop {
			cycles++
		} else {
			cycles += 2
			if cycles >= currentLimit {
				result += currentLimit * x
				currentLimit += 40
			}

			x += commands[i].value
		}
	}

	return result
}

func inSprite(pixel int, sprite int) bool {
	return pixel >= sprite-1 && pixel <= sprite+1
}

func putMark(pixel int, sprite int) byte {
	if inSprite(pixel, sprite) {
		return '#'
	}

	return '.'
}

func part2(commands []command) [][]byte {
	sprite := 1
	currentLimit := 40
	period := 40
	absoluteLimit := 220
	cycles := 0
	totalCycles := 0
	var result [][]byte
	currentRow := make([]byte, period)

	for i := range commands {
		if cycles > absoluteLimit {
			break
		}

		end := 1
		if commands[i].instruction == addx {
			end++
		}

		for j := 0; j < end; j++ {
			cycles++
			currentRow[cycles-1] = putMark(cycles-1, sprite)
			if cycles == currentLimit {
				result = append(result, currentRow)
				totalCycles += cycles
				currentRow = make([]byte, period)
				cycles = 0
			}
		}

		if commands[i].instruction == addx {
			sprite += commands[i].value
		}
	}

	return result
}

func printScreen(screen [][]byte) {
	for i := range screen {
		fmt.Println(string(screen[i]))
	}
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

	commands := readInput(file)
	fmt.Println("Part1:", part1(commands))
	fmt.Println("Part2:")
	printScreen(part2(commands))
}
