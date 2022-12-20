package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readInput(file *os.File) []int {
	scanner := bufio.NewScanner(file)
	var numbers []int

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		var current int
		n, err := fmt.Sscanf(line, "%d", &current)
		if n != 1 || err != nil {
			log.Fatal("Can't parse:", line, err)
		}

		numbers = append(numbers, current)
	}

	return numbers
}

func indexOf(numbers []int, number int) int {
	for i := range numbers {
		if numbers[i] == number {
			return i
		}
	}

	return -1
}

func establishNewIndex(size int, current int, value int) int {
	delta := current + value
	if delta <= 0 {
		delta = 0 - delta
		rest := delta % size

		return size - rest
	}

	return delta % size
}

func removeAt(numbers []int, index int) []int {
	return append(numbers[:index], numbers[index+1:]...)
}

func addAt(numbers []int, value int, index int) []int {
	if index >= len(numbers) {
		return append(numbers, value)
	}

	var temp []int
	temp = append(temp, numbers[:index]...)
	temp = append(temp, value)

	return append(temp, numbers[index:]...)
}

func mix(numbers []int) []int {
	size := len(numbers)
	edge := size - 1
	mixed := make([]int, size)
	copy(mixed, numbers)

	for i := range numbers {
		if numbers[i] == 0 {
			continue
		}

		currentIndex := indexOf(mixed, numbers[i])
		newIndex := establishNewIndex(edge, currentIndex, numbers[i])

		mixed = removeAt(mixed, currentIndex)
		mixed = addAt(mixed, numbers[i], newIndex)
	}

	return mixed
}

func part1(mixed []int) int {
	zeroIndex := indexOf(mixed, 0)
	result := 0
	size := len(mixed)

	for i := 1; i < 4; i++ {
		index := establishNewIndex(size, zeroIndex, i*1000)
		result += mixed[index]
	}

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

	numbers := readInput(file)
	mixed := mix(numbers)
	fmt.Println("Part1:", part1(mixed))
}
