package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

type valve struct {
	name        string
	rate        int
	open        bool
	connections []string
	paths       []vertex
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
	current := &from

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

func generatePaths(vertices []vertex, graph []path, valves map[string]valve) {
	for key, value := range valves {
		value.paths = traverse(vertex{value.name, 0, false}, vertices, graph)
		valves[key] = value
	}
}

func contains(visited []string, name string) bool {
	for i := range visited {
		if visited[i] == name {
			return true
		}
	}

	return false
}

func filtered(vertices []vertex, valves map[string]valve, visited []string) []vertex {
	var result []vertex
	for i := range vertices {
		if contains(visited, vertices[i].name) {
			continue
		}

		val, _ := valves[vertices[i].name]
		if val.rate > 0 {
			result = append(result, vertices[i])
		}
	}

	return result
}

func calculate(moveTo []vertex, valves map[string]valve, visited []string, count int, rate int) int {
	if count >= 30 || len(moveTo) == 0 {
		return rate
	}

	max := 0
	for i := range moveTo {
		currentCount := count + moveTo[i].cost + 1
		if currentCount > 30 {
			continue
		}

		val, _ := valves[moveTo[i].name]
		val.open = true
		valves[moveTo[i].name] = val

		newVisited := make([]string, len(visited))
		copy(newVisited, visited)
		newVisited = append(newVisited, moveTo[i].name)

		canGo := valves[moveTo[i].name].paths
		toCheck := filtered(canGo, valves, newVisited)
		result := calculate(toCheck, valves, newVisited, currentCount, rate+(30-currentCount)*val.rate)
		if result > max {
			max = result
		}
	}

	return max
}

func part1(from vertex, valves map[string]valve) int {
	canGo := valves[from.name].paths
	toCheck := filtered(canGo, valves, []string{})
	result := calculate(toCheck, valves, []string{}, 0, 0)

	return result
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
	generatePaths(vertices, graph, valves)
	fmt.Println("Part1:", part1(vertices[0], valves))
}
