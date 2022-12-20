package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type entry struct {
	value int
	id    int
}

func readInput(file *os.File) []entry {
	scanner := bufio.NewScanner(file)
	var numbers []entry
	count := 0

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		var current entry
		n, err := fmt.Sscanf(line, "%d", &current.value)
		if n != 1 || err != nil {
			log.Fatal("Can't parse:", line, err)
		}

		if current.value == 0 {
			current.id = -1
		} else {
			current.id = count
		}

		numbers = append(numbers, current)
		count++
	}

	return numbers
}

func indexOf(numbers []entry, id int) int {
	for i := range numbers {
		if numbers[i].id == id {
			return i
		}
	}

	return -1
}

func establishNewIndex(edge int, current int, value int) int {
	delta := current + value
	if delta <= 0 {
		delta = 0 - delta
		rest := delta % edge

		return edge - rest
	}

	return delta % edge
}

func removeAt(numbers []entry, index int) []entry {
	return append(numbers[:index], numbers[index+1:]...)
}

func addAt(numbers []entry, value entry, index int) []entry {
	if index >= len(numbers) {
		return append(numbers, value)
	}

	var temp []entry
	temp = append(temp, numbers[:index]...)
	temp = append(temp, value)

	return append(temp, numbers[index:]...)
}

func mix(numbers []entry, times int) []entry {
	size := len(numbers)
	edge := size - 1
	mixed := make([]entry, size)
	copy(mixed, numbers)

	for t := 0; t < times; t++ {
		for i := range numbers {
			if numbers[i].value == 0 {
				continue
			}

			currentIndex := indexOf(mixed, numbers[i].id)
			newIndex := establishNewIndex(edge, currentIndex, numbers[i].value)

			mixed = removeAt(mixed, currentIndex)
			mixed = addAt(mixed, numbers[i], newIndex)
		}
	}

	return mixed
}

func calculate(mixed []entry) int {
	zeroIndex := indexOf(mixed, -1)
	result := 0
	size := len(mixed)

	for i := 1; i < 4; i++ {
		index := establishNewIndex(size, zeroIndex, i*1000)
		result += mixed[index].value
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
	mixed1 := mix(numbers, 1)
	fmt.Println("Part1:", calculate(mixed1))
}
