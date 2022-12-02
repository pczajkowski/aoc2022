package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type round struct {
	elf   byte
	me    byte
	score int
}

func fight(elf byte, me byte) int {
	if me == 'X' && elf == 'C' {
		return 6
	} else if me == 'Z' && elf == 'B' {
		return 6
	} else if me == 'Y' && elf == 'A' {
		return 6
	}

	return 0
}

func readInput(file *os.File, points map[byte]int) []round {
	scanner := bufio.NewScanner(file)
	var rounds []round

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		var elf byte
		var me byte
		n, err := fmt.Sscanf(line, "%c %c", &elf, &me)
		if n != 2 || err != nil {
			log.Fatal("Can't parse input")
		}

		score := 0
		if points[elf] == points[me] {
			score += 3
		} else {
			score += fight(elf, me)
		}

		score += points[me]
		rounds = append(rounds, round{elf, me, score})
	}

	return rounds
}

func roundsScore(rounds []round) int {
	sum := 0
	for i := range rounds {
		sum += rounds[i].score
	}

	return sum
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

	points := make(map[byte]int)
	points['A'] = 1
	points['X'] = 1
	points['B'] = 2
	points['Y'] = 2
	points['C'] = 3
	points['Z'] = 3

	rounds := readInput(file, points)
	fmt.Println("Part1:", roundsScore(rounds))
}
