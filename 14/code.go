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

func adjustEdges(left point, right point, current point) (point, point) {
	if left.x > current.x {
		left.x = current.x
	}

	if left.y < current.y {
		left.y = current.y
		right.y = current.y
	}

	if right.x < current.x {
		right.x = current.x
	}

	return left, right
}

func drawHorizontal(first point, second point, rocks map[point]bool) {
	start := first.x
	end := second.x
	if start > second.x {
		start = second.x
		end = first.x
	}

	for i := start + 1; i < end; i++ {
		fmt.Println("h", first, second, i)
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
		fmt.Println("v", first, second, i)
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

func buildRocks(points [][]point) map[point]bool {
	rocks := make(map[point]bool)
	leftEdge := point{500, 0}
	rightEdge := point{500, 0}

	for i := range points {
		edge := len(points[i])
		prev := points[i][0]
		rocks[prev] = true
		leftEdge, rightEdge = adjustEdges(leftEdge, rightEdge, prev)

		for j := 1; j < edge; j++ {
			current := points[i][j]
			rocks[current] = true
			leftEdge, rightEdge = adjustEdges(leftEdge, rightEdge, current)

			drawPath(prev, current, rocks)
			prev = current
		}
	}

	return rocks
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
	rocks := buildRocks(points)
	fmt.Println(rocks)
}
