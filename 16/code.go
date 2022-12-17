package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
)

type valve struct {
	name        string
	rate        int
	open        bool
	connections []string
}

type vertex struct {
	name    string
	cost    int
	visited bool
}

const maxValue int = 100000

func readInput(file *os.File) ([]vertex, map[string]valve) {
	scanner := bufio.NewScanner(file)
	valves := make(map[string]valve)
	var vertices []vertex

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		var current valve
		if strings.Contains(line, "valves") {
			n, err := fmt.Sscanf(line, "Valve %s has flow rate=%d", &current.name, &current.rate)
			if n != 2 || err != nil {
				log.Fatal("Can't parse (valves):", line, err)
			}

			re := regexp.MustCompile(`valves .*`)
			parts := re.FindString(line)
			parts = strings.TrimLeft(parts, "valves ")
			current.connections = strings.Split(parts, ", ")
		} else {
			var connection string
			n, err := fmt.Sscanf(line, "Valve %s has flow rate=%d; tunnel leads to valve %s", &current.name, &current.rate, &connection)
			if n != 3 || err != nil {
				log.Fatal("Can't parse:", line, err)
			}

			current.connections = append(current.connections, connection)
		}

		vertices = append(vertices, vertex{current.name, maxValue, false})
		valves[current.name] = current
	}

	return vertices, valves
}

type path struct {
	from string
	to   string
	cost int
}

func buildGraph(valves map[string]valve) []path {
	var graph []path
	for key, value := range valves {
		for i := range value.connections {
			graph = append(graph, path{key, value.connections[i], 1})
		}
	}

	return graph
}

func setCost(name string, cost int, vertices []vertex) {
	for i := range vertices {
		if vertices[i].name == name {
			vertices[i].cost = cost
			break
		}
	}
}

func getCost(name string, vertices []vertex) int {
	for i := range vertices {
		if vertices[i].name == name {
			return vertices[i].cost
		}
	}

	return 0
}

func getNext(vertices []vertex) *vertex {
	min := maxValue
	var current *vertex
	for i := range vertices {
		if vertices[i].visited {
			continue
		}

		if vertices[i].cost <= min {
			min = vertices[i].cost
			current = &vertices[i]
		}
	}

	return current
}

func traverse(from vertex, vertices []vertex, graph []path) []vertex {
	newVertices := make([]vertex, len(vertices))
	copy(newVertices, vertices)
	current := &vertex{from.name, 0, false}

	for {
		for j := range graph {
			if graph[j].from != current.name {
				continue
			}

			var tentativeCost int
			if current.cost == maxValue {
				tentativeCost = maxValue
			} else {
				tentativeCost = current.cost + 1
			}

			if tentativeCost < getCost(graph[j].to, newVertices) {
				setCost(graph[j].to, tentativeCost, newVertices)
			}
		}

		current.visited = true

		current = getNext(newVertices)
		if current == nil {
			break
		}
	}

	return newVertices
}

func filter(vertices []vertex, valves map[string]valve) []vertex {
	var result []vertex
	for i := range vertices {
		val, _ := valves[vertices[i].name]
		if !val.open && val.rate > 0 {
			result = append(result, vertices[i])
		}
	}

	return result
}

func moveTo(vertices []vertex, valves map[string]valve) *vertex {
	filtered := filter(vertices, valves)
	if len(filtered) == 0 {
		return nil
	}

	sort.Slice(filtered, func(i, j int) bool {
		return valves[filtered[i].name].rate-filtered[i].cost > valves[filtered[j].name].rate-filtered[j].cost
	})

	return &filtered[0]
}

func part1(vertices []vertex, graph []path, valves map[string]valve) int {
	count := 0
	rate := 0
	current := &vertices[0]
	limit := 30

	for {
		if count >= limit {
			break
		}

		val, _ := valves[current.name]
		if !val.open && val.rate > 0 {
			count++
			val.open = true
			rate += (limit - count) * val.rate
			valves[current.name] = val
			fmt.Println(current, count)
		}

		canGo := traverse(*current, vertices, graph)

		current = moveTo(canGo, valves)
		if current == nil {
			break
		}

		count += current.cost
	}

	return rate
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

	vertices, valves := readInput(file)
	graph := buildGraph(valves)
	fmt.Println("Part1:", part1(vertices, graph, valves))
}
