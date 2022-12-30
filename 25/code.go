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

func toSnafu(number int) string {
	multiplier := 5
	modifier := 1
	count := 1
	var result []byte

	found := false
	toMatch := 0
	for {
		for i := 1; i < 3; i++ {
			if i*modifier >= number {
				found = true
				result = append(result, getChar(i))
				toMatch = modifier*i - number
				break
			}
		}

		if found {
			break
		}

		modifier *= multiplier
		count++
	}

	for i := 1; i < count; i++ {
		modifier /= multiplier

		for j := -2; j <= 2; j++ {
			p := j * modifier
			delta := toMatch + p
			if delta >= 0 || 0-delta < 2*modifier/multiplier {
				result = append(result, getChar(j))
				toMatch = delta
				break
			}
		}
	}

	return string(result)
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
	fmt.Println(sum)
	fmt.Println(toSnafu(sum))
}
