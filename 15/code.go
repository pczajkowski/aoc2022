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

func abs(x int) int {
	if x < 0 {
		return 0 - x
	}

	return x
}

func getDistance(sensor point, beacon point) int {
	return abs(sensor.x-beacon.x) + abs(sensor.y-beacon.y)
}

func markOccupied(sensor point, distance int, occupied map[int]map[int]byte, row int) {
	width := distance - abs(row-sensor.y)
	if occupied[row] == nil {
		occupied[row] = make(map[int]byte)
	}

	for j := sensor.x - width; j <= sensor.x+width; j++ {
		_, ok := occupied[row][j]
		if !ok {
			occupied[row][j] = '#'
		}
	}
}

func drawLines(points [][2]point, row int) map[int]map[int]byte {
	occupied := make(map[int]map[int]byte)
	for i := range points {
		if occupied[points[i][0].y] == nil {
			occupied[points[i][0].y] = make(map[int]byte)
		}
		occupied[points[i][0].y][points[i][0].x] = 's'

		if occupied[points[i][1].y] == nil {
			occupied[points[i][1].y] = make(map[int]byte)
		}
		occupied[points[i][1].y][points[i][1].x] = 'b'

		distance := getDistance(points[i][0], points[i][1])
		markOccupied(points[i][0], distance, occupied, row)
	}

	return occupied
}

func filterPoints(points [][2]point, row int) [][2]point {
	var filtered [][2]point
	for i := range points {
		distance := getDistance(points[i][0], points[i][1])
		if row >= points[i][0].y-distance && row <= points[i][0].y+distance {
			filtered = append(filtered, points[i])
		}
	}

	return filtered
}

func part1(points [][2]point, row int) int {
	filtered := filterPoints(points, row)
	occupied := drawLines(filtered, row)
	count := 0
	for _, value := range occupied[row] {
		if value == '#' {
			count++
		}
	}

	return count
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
	fmt.Println("Part1:", part1(points, 2000000))
}
