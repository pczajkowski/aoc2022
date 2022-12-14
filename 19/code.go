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

func canProduceOre(plan blueprint, resources inventory) bool {
	return resources.ore >= plan.oreCost
}

func produceOre(plan blueprint, resources inventory) inventory {
	resources.ore -= plan.oreCost
	return resources
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

func shouldProduce(currentCount int, plan blueprint, robots inventory, robotsWith inventory, resources inventory, resourcesWith inventory, check func(blueprint, inventory) bool) bool {
	if currentCount == 0 {
		return true
	}

	countWithout := 0
	for {
		if check(plan, resources) {
			break
		}

		resources = produce(robots, resources)
		countWithout++
	}

	countWith := 0
	for {
		if check(plan, resourcesWith) {
			break
		}

		resourcesWith = produce(robotsWith, resourcesWith)
		countWith++
	}

	return countWith <= countWithout
}

func shouldProduceOre(plan blueprint, robots inventory, resources inventory) bool {
	robotsWith := robots
	robotsWith.ore++

	resourcesWith := resources
	resourcesWith = produceOre(plan, resourcesWith)

	return shouldProduce(robots.ore, plan, robots, robotsWith, resources, resourcesWith, canProduceClay)
}

func shouldProduceClay(plan blueprint, robots inventory, resources inventory) bool {
	robotsWith := robots
	robotsWith.clay++

	resourcesWith := resources
	resourcesWith = produceClay(plan, resourcesWith)

	return shouldProduce(robots.clay, plan, robots, robotsWith, resources, resourcesWith, canProduceObsidian)
}

func shouldProduceObsidian(plan blueprint, robots inventory, resources inventory) bool {
	robotsWith := robots
	robotsWith.obsidian++

	resourcesWith := resources
	resourcesWith = produceObsidian(plan, resourcesWith)

	return shouldProduce(robots.obsidian, plan, robots, robotsWith, resources, resourcesWith, canProduceGeode)
}

func checkPlan(plan blueprint) int {
	var robots inventory
	robots.ore++

	var resources inventory

	for i := 0; i < 24; i++ {
		newRobots := robots
		if canProduceGeode(plan, resources) {
			newRobots.geode++
			resources = produceGeode(plan, resources)
		} else if canProduceObsidian(plan, resources) {
			if shouldProduceObsidian(plan, robots, resources) {
				newRobots.obsidian++
				resources = produceObsidian(plan, resources)
			}
		} else if robots.clay < plan.obsidianCost[1] && canProduceClay(plan, resources) {
			if shouldProduceClay(plan, robots, resources) {
				newRobots.clay++
				resources = produceClay(plan, resources)
			}
		} else if robots.ore < plan.clayCost && canProduceOre(plan, resources) {
			if shouldProduceClay(plan, robots, resources) {
				newRobots.ore++
				resources = produceOre(plan, resources)
			}
		}

		resources = produce(robots, resources)
		fmt.Println(plan.id, i+1, robots, resources)
		robots = newRobots
	}

	fmt.Println(plan.id, resources.geode)
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
