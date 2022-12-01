package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func readInput(file *os.File) [3]int {
	scanner := bufio.NewScanner(file)
	var numbers [3]int
	current := 0

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			for i := range numbers {
				if current > numbers[i] {
					current, numbers[i] = numbers[i], current
				}
			}

			current = 0
			continue
		}

		if number, err := strconv.Atoi(line); err == nil {
			current += number
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
	fmt.Println("Part1:", numbers[0])
	fmt.Println("Part2:", numbers[0]+numbers[1]+numbers[2])
}
