package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type pair struct {
	first  [2]int
	second [2]int
}

func readInput(file *os.File) []pair {
	scanner := bufio.NewScanner(file)
	var pairs []pair

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		current := pair{}
		n, err := fmt.Sscanf(line, "%d-%d,%d-%d", &current.first[0], &current.first[1], &current.second[0], &current.second[1])
		if n != 4 || err != nil {
			log.Fatal("Problem reading input:", err)
		}

		pairs = append(pairs, current)
	}

	return pairs
}

func contains(first [2]int, second [2]int) bool {
	if first[0] >= second[0] && second[0] < first[1] && first[0] > second[1] && second[1] <= first[1] {
		return true
	}

	return false
}

func part1(pairs []pair) int {
	count := 0
	for i := range pairs {
		if contains(pairs[i].first, pairs[i].second) || contains(pairs[i].second, pairs[i].first) {
			count++
		}
	}

	return count
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

	pairs := readInput(file)
	fmt.Println("Part1:", part1(pairs))
}
