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

const multiplier int = 5

func fromSnafu(text string) int {
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

func sum(numbers []int) int {
	sum := 0
	for i := range numbers {
		sum += numbers[i]
	}

	return sum
}

func getChar(number int) byte {
	switch number {
	case 2:
		return '2'
	case 1:
		return '1'
	case 0:
		return '0'
	case -1:
		return '-'
	case -2:
		return '='
	}

	return ' '
}

func abs(number int) int {
	if number < 0 {
		return 0 - number
	}

	return number
}

func reverse(bytes []byte) []byte {
	edge := len(bytes) - 1
	var reversed []byte

	for i := edge; i >= 0; i-- {
		reversed = append(reversed, bytes[i])
	}

	return reversed
}

func toSnafu(number int) string {
	var result []byte

	for {
		if number <= 0 {
			break
		}

		rem := number % multiplier
		number /= multiplier
		if rem == 3 {
			rem = -2
			number += 1
		} else if rem == 4 {
			rem = -1
			number++
		}

		result = append(result, getChar(rem))
	}

	return string(reverse(result))
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
	sum := sum(numbers)
	fmt.Println("Part1:", toSnafu(sum))
}
