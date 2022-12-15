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

func markOccupied(sensor point, distance int, occupied map[int]map[int]byte) {
	end := sensor.y - distance

	width := distance
	for i := sensor.y; i >= end; i-- {
		if occupied[i] == nil {
			occupied[i] = make(map[int]byte)
		}

		for j := sensor.x - width; j <= sensor.x+width; j++ {
			_, ok := occupied[i][j]
			if !ok {
				occupied[i][j] = '#'
			}
		}
		width--
	}

	width = distance - 1
	end = sensor.y + distance
	for i := sensor.y + 1; i <= end; i++ {
		if occupied[i] == nil {
			occupied[i] = make(map[int]byte)
		}

		for j := sensor.x - width; j <= sensor.x+width; j++ {
			occupied[i][j] = '#'
		}
		width--
	}
}

func drawLines(points [][2]point) map[int]map[int]byte {
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
		markOccupied(points[i][0], distance, occupied)
	}

	return occupied
}

func part1(occupied map[int]map[int]byte, row int) int {
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
	occupied := drawLines(points)
	fmt.Println("Part1:", part1(occupied, 2000000))
}
