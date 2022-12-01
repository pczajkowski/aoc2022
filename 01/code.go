package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func readInput(file *os.File) []int {
	scanner := bufio.NewScanner(file)
	numbers := []int{0}
	index := 0

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			index++
			numbers = append(numbers, 0)
			continue
		}

		if number, err := strconv.Atoi(line); err == nil {
			numbers[index] += number
		} else {
			log.Fatal("Numbers: ", err)
		}
	}

	return numbers
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
	fmt.Println(numbers)
}
