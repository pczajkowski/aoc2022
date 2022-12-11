package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type monkey struct {
	items []int
}

func readItems(line string) []int {
	line = strings.Replace(line, "  Starting items: ", "", 1)
	parts := strings.Split(line, ", ")

	var result []int
	for i := range parts {
		n, err := strconv.Atoi(parts[i])
		if err != nil {
			log.Fatal("Can't pasrse", parts[i], err)
		}

		result = append(result, n)
	}

	return result
}

func readInput(file *os.File) []monkey {
	scanner := bufio.NewScanner(file)
	counter := 0
	var monkeys []monkey
	var currentMonkey monkey

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			monkeys = append(monkeys, currentMonkey)
			counter = 0
			currentMonkey = monkey{}
			continue
		}

		switch counter {
		case 1:
			currentMonkey.items = readItems(line)
		}
		counter++
	}

	monkeys = append(monkeys, currentMonkey)
	return monkeys
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

	monkeys := readInput(file)
	fmt.Println(monkeys)
}
