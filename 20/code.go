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
	if delta < 0 {
		delta = 0 - delta
		rest := delta % size

		return size - 1 - rest
	}

	return delta % size
}

func mix(numbers []int) []int {
	size := len(numbers)
	mixed := make([]int, size)
	copy(mixed, numbers)

	for i := range numbers {
		currentIndex := indexOf(mixed, numbers[i])
		newIndex := establishNewIndex(size, i, numbers[i])

		fmt.Println(currentIndex, newIndex, numbers[i])
	}

	return mixed
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
	fmt.Println(mix(numbers))
}
