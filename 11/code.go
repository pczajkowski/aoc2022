package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type kind byte

const (
	old kind = iota
	val
)

type variable struct {
	t   kind
	val int
}

type operation struct {
	x      variable
	y      variable
	action byte
}

type monkey struct {
	items []int
	op    operation
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

func readVariable(text string) variable {
	var result variable
	if text == "old" {
		result.t = old
	} else {
		result.t = val
		n, err := strconv.Atoi(text)
		if err != nil {
			log.Fatal("Can't pasrse", text, err)
		}

		result.val = n
	}

	return result
}

func readOperation(line string) operation {
	line = strings.Replace(line, "  Operation: new = ", "", 1)
	parts := strings.Split(line, " ")
	if len(parts) != 3 {
		log.Fatal("Bad operation input:", line)
	}

	var result operation
	result.x = readVariable(parts[0])
	result.y = readVariable(parts[2])
	result.action = parts[1][0]

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
		case 2:
			currentMonkey.op = readOperation(line)
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
