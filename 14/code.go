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

func drawHorizontal(first point, second point, rocks map[point]bool) {
	start := first.x
	end := second.x
	if start > second.x {
		start = second.x
		end = first.x
	}

	for i := start + 1; i < end; i++ {
		rocks[point{i, first.y}] = true
	}
}

func drawVertical(first point, second point, rocks map[point]bool) {
	start := first.y
	end := second.y
	if start > second.y {
		start = second.y
		end = first.y
	}

	for i := start + 1; i < end; i++ {
		rocks[point{first.x, i}] = true
	}
}

func drawPath(first point, second point, rocks map[point]bool) {
	if first.x == second.x {
		drawVertical(first, second, rocks)
	} else {
		drawHorizontal(first, second, rocks)
	}
}

func buildRocks(points [][]point) (map[point]bool, int) {
	rocks := make(map[point]bool)
	bottom := 0

	for i := range points {
		edge := len(points[i])
		prev := points[i][0]
		rocks[prev] = true
		if prev.y > bottom {
			bottom = prev.y
		}

		for j := 1; j < edge; j++ {
			current := points[i][j]
			rocks[current] = true
			if current.y > bottom {
				bottom = current.y
			}

			drawPath(prev, current, rocks)
			prev = current
		}
	}

	return rocks, bottom
}

func moveUnit(unit point, rocks map[point]bool) point {
	current := point{unit.x, unit.y + 1}
	if !rocks[current] {
		return current
	}

	current.x--
	if !rocks[current] {
		return current
	}

	current.x += 2
	if !rocks[current] {
		return current
	}

	return unit
}

func fall(unit point, rocks map[point]bool, bottom int) bool {
	for {
		current := moveUnit(unit, rocks)
		if current.x == unit.x && current.y == unit.y {
			rocks[current] = true
			break

		}

		if current.y > bottom {
			return true
		}

		unit = current
	}

	return false
}

func sandstorm(rocks map[point]bool, bottom int) int {
	initialRocks := len(rocks)
	currentRocks := initialRocks

	for {
		if fall(point{500, 0}, rocks, bottom) {
			break
		}

		if len(rocks) == currentRocks {
			break
		}

		currentRocks = len(rocks)
	}

	return len(rocks) - initialRocks
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
	rocks, edges := buildRocks(points)
	fmt.Println("Part1:", sandstorm(rocks, edges))
}
