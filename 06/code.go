package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func readInput(path string) []byte {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	return data
}

func check(duplicates map[byte]int) bool {
	for _, value := range duplicates {
		if value > 1 {
			return false
		}
	}

	return true
}

func part1(text []byte) int {
	count := 0
	toDelete := 0
	duplicates := make(map[byte]int)

	for i := range text {
		count++

		duplicates[text[i]]++
		if count-toDelete > 3 {
			if check(duplicates) {
				break
			}

			duplicates[text[toDelete]]--
			toDelete++
		}
	}

	return count
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("You need to specify a file!")
	}

	filePath := os.Args[1]
	text := readInput(filePath)
	fmt.Println("Part1:", part1(text))
}
