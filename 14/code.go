package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type point struct {
	x int
	y int
}

func readPoints(line string) []point {
	parts := strings.Split(line, " -> ")
	if len(parts) == 0 {
		log.Fatal("Wrong input:", line)
	}

	var points []point
	for i := range parts {
		var current point
		n, err := fmt.Sscanf(parts[i], "%d,%d", &current.x, &current.y)
		if n != 2 || err != nil {
			log.Fatal("Can't parse", line, err)
		}

		points = append(points, current)
	}

	return points
}

func readInput(file *os.File) [][]point {
	scanner := bufio.NewScanner(file)
	var points [][]point

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		points = append(points, readPoints(line))
	}

	return points
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

	points := readInput(file)
	fmt.Println(points)
}
