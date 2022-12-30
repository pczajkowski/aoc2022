package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func getNumber(char byte) int {
	switch char {
	case '2':
		return 2
	case '1':
		return 1
	case '0':
		return 0
	case '-':
		return -1
	case '=':
		return -2
	}

	return 300
}

func fromSnafu(text string) int {
	multiplier := 5
	modifier := 1
	end := len(text) - 1
	result := 0

	for i := end; i >= 0; i-- {
		n := getNumber(text[i])
		result += modifier * n
		modifier *= multiplier
	}

	return result
}

func readInput(file *os.File) []int {
	scanner := bufio.NewScanner(file)
	var numbers []int

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		numbers = append(numbers, fromSnafu(line))
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
