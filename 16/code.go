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
}

func readInput(file *os.File) map[string]valve {
	scanner := bufio.NewScanner(file)
	valves := make(map[string]valve)

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

		valves[current.name] = current
	}

	return valves
}

type path struct {
	from string
	to   string
}

func buildGraph(valves map[string]valve) []path {
	var graph []path
	for key, value := range valves {
		for i := range value.connections {
			graph = append(graph, path{key, value.connections[i]})
		}
	}

	return graph
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

	valves := readInput(file)
	graph := buildGraph(valves)
	fmt.Println(graph)
}
