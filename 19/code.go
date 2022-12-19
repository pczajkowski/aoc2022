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

type inventory struct {
	ore      int
	clay     int
	obsidian int
	geode    int
}

func canProduceClay(plan blueprint, resources inventory) bool {
	return resources.ore >= plan.clayCost
}

func produceClay(plan blueprint, resources inventory) inventory {
	resources.ore -= plan.clayCost
	return resources
}

func canProduceObsidian(plan blueprint, resources inventory) bool {
	return resources.ore >= plan.obsidianCost[0] && resources.clay >= plan.obsidianCost[1]
}

func produceObsidian(plan blueprint, resources inventory) inventory {
	resources.ore -= plan.obsidianCost[0]
	resources.clay -= plan.obsidianCost[1]
	return resources
}
func canProduceGeode(plan blueprint, resources inventory) bool {
	return resources.ore >= plan.geodeCost[0] && resources.obsidian >= plan.geodeCost[1]
}

func produceGeode(plan blueprint, resources inventory) inventory {
	resources.ore -= plan.geodeCost[0]
	resources.obsidian -= plan.geodeCost[1]
	return resources
}

func produce(robots inventory, resources inventory) inventory {
	resources.ore += robots.ore
	resources.clay += robots.clay
	resources.obsidian += robots.obsidian
	resources.geode += robots.geode

	return resources
}

func checkPlan(plan blueprint) int {
	var robots inventory
	robots.ore++

	var resources inventory

	for i := 0; i < 24; i++ {
		resources = produce(robots, resources)

		if canProduceGeode(plan, resources) {
			robots.geode++
			resources = produceGeode(plan, resources)
		}

		if canProduceObsidian(plan, resources) {
			robots.obsidian++
			resources = produceObsidian(plan, resources)
		}

		if canProduceClay(plan, resources) {
			robots.clay++
			resources = produceClay(plan, resources)
		}
	}

	return resources.geode * plan.id
}

func part1(blueprints []blueprint) int {
	sum := 0
	for i := range blueprints {
		sum += checkPlan(blueprints[i])
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

	blueprints := readInput(file)
	fmt.Println("Part1:", part1(blueprints))
}
