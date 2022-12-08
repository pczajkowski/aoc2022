package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readInput(file *os.File) [][]byte {
	scanner := bufio.NewScanner(file)
	var trees [][]byte

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		trees = append(trees, []byte(line))
	}

	return trees
}

func visibleFromLeft(x int, y int, trees [][]byte) bool {
	visible := true
	for i := 0; i < x; i++ {
		if trees[y][i] >= trees[y][x] {
			visible = false
			break
		}
	}

	return visible
}

func visibleFromRight(x int, y int, trees [][]byte, limit int) bool {
	visible := true
	for i := x + 1; i < limit; i++ {
		if trees[y][i] >= trees[y][x] {
			visible = false
			break
		}
	}

	return visible
}

func visibleFromTop(x int, y int, trees [][]byte) bool {
	visible := true
	for i := 0; i < y; i++ {
		if trees[i][x] >= trees[y][x] {
			visible = false
			break
		}
	}

	return visible
}

func visibleFromBottom(x int, y int, trees [][]byte, limit int) bool {
	visible := true
	for i := y + 1; i < limit; i++ {
		if trees[i][x] >= trees[y][x] {
			visible = false
			break
		}
	}

	return visible
}

func isVisible(x int, y int, trees [][]byte, width int, height int) bool {
	if visibleFromLeft(x, y, trees) {
		return true
	}

	if visibleFromRight(x, y, trees, width) {
		return true
	}

	if visibleFromTop(x, y, trees) {
		return true
	}

	if visibleFromBottom(x, y, trees, height) {
		return true
	}

	return false
}

func part1(trees [][]byte) int {
	width := len(trees[0])
	height := len(trees)

	visible := 2*height + (width-2)*2
	for y := 1; y < height-1; y++ {
		for x := 1; x < width-1; x++ {
			if isVisible(x, y, trees, width, height) {
				visible++
			}
		}
	}

	return visible
}

func numberFromLeft(x int, y int, trees [][]byte) int {
	count := 0
	for i := x - 1; i >= 0; i-- {
		count++
		if trees[y][i] >= trees[y][x] {
			break
		}
	}

	return count
}

func numberFromRight(x int, y int, trees [][]byte, limit int) int {
	count := 0
	for i := x + 1; i < limit; i++ {
		count++
		if trees[y][i] >= trees[y][x] {
			break
		}
	}

	return count
}

func numberFromTop(x int, y int, trees [][]byte) int {
	count := 0
	for i := y - 1; i >= 0; i-- {
		count++
		if trees[i][x] >= trees[y][x] {
			break
		}
	}

	return count
}

func numberFromBottom(x int, y int, trees [][]byte, limit int) int {
	count := 0
	for i := y + 1; i < limit; i++ {
		count++
		if trees[i][x] >= trees[y][x] {
			break
		}
	}

	return count
}

func numberVisible(x int, y int, trees [][]byte, width int, height int) int {
	count := 1

	count *= numberFromLeft(x, y, trees)
	count *= numberFromRight(x, y, trees, width)
	count *= numberFromTop(x, y, trees)
	count *= numberFromBottom(x, y, trees, height)

	return count
}

func part2(trees [][]byte) int {
	width := len(trees[0])
	height := len(trees)
	max := 0

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			count := numberVisible(x, y, trees, width, height)
			if count > max {
				max = count
			}
		}
	}

	return max
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

	trees := readInput(file)
	fmt.Println("Part1:", part1(trees))
	fmt.Println("Part2:", part2(trees))
}
