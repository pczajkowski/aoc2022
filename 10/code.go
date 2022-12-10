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
	fmt.Println(commands)
}
