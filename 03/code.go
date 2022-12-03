package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type rucksack struct {
	first  map[byte]int
	second map[byte]int
}

func readInput(file *os.File) []rucksack {
	scanner := bufio.NewScanner(file)
	var rucksacks []rucksack

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		edge := len(line)
		half := edge / 2
		start := 0

		current := rucksack{make(map[byte]int), make(map[byte]int)}
		for {
			if half >= edge {
				break
			}

			current.first[line[start]]++
			current.second[line[half]]++
			start++
			half++
		}

		rucksacks = append(rucksacks, current)
	}

	return rucksacks
}

func part1(rucksacks []rucksack) int {
	sum := 0
	for i := range rucksacks {
		for key, _ := range rucksacks[i].first {
			if rucksacks[i].second[key] > 0 {
				if key < 96 {
					sum += int(key) - 38
				} else {
					sum += int(key) - 96
				}
			}
		}
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

	rucksacks := readInput(file)
	fmt.Println("Part1:", part1(rucksacks))
}
