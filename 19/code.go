package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type blueprint struct {
	id           int
	oreCost      int
	clayCost     int
	obsidianCost [2]int
	geodeCost    [2]int
}

func readInput(file *os.File) []blueprint {
	scanner := bufio.NewScanner(file)
	var blueprints []blueprint

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		var current blueprint
		n, err := fmt.Sscanf(line, "Blueprint %d: Each ore robot costs %d ore. Each clay robot costs %d ore. Each obsidian robot costs %d ore and %d clay. Each geode robot costs %d ore and %d obsidian.", &current.id, &current.oreCost, &current.clayCost, &current.obsidianCost[0], &current.obsidianCost[1], &current.geodeCost[0], &current.geodeCost[1])
		if n != 7 || err != nil {
			log.Fatal("Can't parse:", line, err)
		}

		blueprints = append(blueprints, current)
	}

	return blueprints
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

	blueprints := readInput(file)
	fmt.Println(blueprints)
}
