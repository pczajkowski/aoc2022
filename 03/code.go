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

func getPriority(item byte) int {
	if item < 96 {
		return int(item) - 38
	}

	return int(item) - 96
}

func part1(rucksacks []rucksack) int {
	sum := 0
	for i := range rucksacks {
		for key, _ := range rucksacks[i].first {
			if rucksacks[i].second[key] > 0 {
				sum += getPriority(key)
			}
		}
	}

	return sum
}

func checkCompartments(compartment map[byte]int, rucksacks []rucksack, index int) int {
	for key, _ := range compartment {
		if rucksacks[index+1].first[key] == 0 && rucksacks[index+1].second[key] == 0 {
			continue
		}

		if rucksacks[index+2].first[key] == 0 && rucksacks[index+2].second[key] == 0 {
			continue
		}

		return getPriority(key)
	}

	return 0
}

func part2(rucksacks []rucksack) int {
	edge := len(rucksacks) - 2
	index := 0
	sum := 0

	for {
		if index >= edge {
			break
		}

		result := checkCompartments(rucksacks[index].first, rucksacks, index)
		if result == 0 {
			result = checkCompartments(rucksacks[index].second, rucksacks, index)
		}

		sum += result
		index += 3
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
	fmt.Println("Part2:", part2(rucksacks))
}
