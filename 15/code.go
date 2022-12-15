package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type point struct {
	x int
	y int
}

func readInput(file *os.File) [][2]point {
	scanner := bufio.NewScanner(file)
	var points [][2]point

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		var current [2]point
		n, err := fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &current[0].x, &current[0].y, &current[1].x, &current[1].y)
		if n != 4 || err != nil {
			log.Fatal("Can't parse", line, err)
		}
		points = append(points, current)
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
