package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type kind byte

const (
	op kind = iota
	val
)

type monkey struct {
	name  string
	spec  kind
	left  string
	right string
	op    byte
	value int
}

func readInput(file *os.File) map[string]monkey {
	scanner := bufio.NewScanner(file)
	monkeys := make(map[string]monkey)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		var current monkey
		if strings.ContainsAny(line, "+-/*") {
			current.spec = op
			n, err := fmt.Sscanf(line, "%s %s %c %s", &current.name, &current.left, &current.op, &current.right)
			if n != 4 || err != nil {
				log.Fatal("Can't parse (op): ", line, err)
			}
		} else {
			current.spec = val
			n, err := fmt.Sscanf(line, "%s %d", &current.name, &current.value)
			if n != 2 || err != nil {
				log.Fatal("Can't parse: ", line, err)
			}
		}

		current.name = strings.TrimRight(current.name, ":")
		monkeys[current.name] = current
	}

	return monkeys
}

func processMonkey(being monkey, monkeys map[string]monkey) int {
	if being.spec == val {
		return being.value
	}

	switch being.op {
	case '+':
		return processMonkey(monkeys[being.left], monkeys) + processMonkey(monkeys[being.right], monkeys)
	case '-':
		return processMonkey(monkeys[being.left], monkeys) - processMonkey(monkeys[being.right], monkeys)
	case '*':
		return processMonkey(monkeys[being.left], monkeys) * processMonkey(monkeys[being.right], monkeys)
	case '/':
		return processMonkey(monkeys[being.left], monkeys) / processMonkey(monkeys[being.right], monkeys)
	}

	return 0
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
	fmt.Println("Part1:", processMonkey(monkeys["root"], monkeys))
}
