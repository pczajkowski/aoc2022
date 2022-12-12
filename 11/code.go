package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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
	items   []int
	op      operation
	test    int
	iftrue  int
	iffalse int
	counter int
}

func readItems(line string) []int {
	line = strings.Replace(line, "  Starting items: ", "", 1)
	parts := strings.Split(line, ", ")

	var result []int
	for i := range parts {
		n, err := strconv.Atoi(parts[i])
		if err != nil {
			log.Fatal("Can't parse", parts[i], err)
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
			log.Fatal("Can't parse", text, err)
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

func readInt(line string, format string) int {
	var result int
	n, err := fmt.Sscanf(line, format, &result)
	if n != 1 || err != nil {
		log.Fatal("Can't parse", line, err)
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
		case 2:
			currentMonkey.op = readOperation(line)
		case 3:
			currentMonkey.test = readInt(line, "  Test: divisible by %d")
		case 4:
			currentMonkey.iftrue = readInt(line, "    If true: throw to monkey %d")
		case 5:
			currentMonkey.iffalse = readInt(line, "    If false: throw to monkey %d")
		}
		counter++
	}

	monkeys = append(monkeys, currentMonkey)
	return monkeys
}

func performOperation(mon monkey, itemIndex int) int {
	var x int
	if mon.op.x.t == old {
		x = mon.items[itemIndex]
	} else {
		x = mon.op.x.val
	}

	var y int
	if mon.op.y.t == old {
		y = mon.items[itemIndex]
	} else {
		y = mon.op.y.val
	}

	if mon.op.action == '+' {
		return x + y
	}

	return x * y
}

func processMonkey(index int, monkeys []monkey, relief int, divider int) []monkey {
	for i := range monkeys[index].items {
		worryLevel := performOperation(monkeys[index], i) / relief % divider

		if worryLevel%monkeys[index].test == 0 {
			monkeys[monkeys[index].iftrue].items = append(monkeys[monkeys[index].iftrue].items, worryLevel)
		} else {
			monkeys[monkeys[index].iffalse].items = append(monkeys[monkeys[index].iffalse].items, worryLevel)
		}

		monkeys[index].counter++
	}

	monkeys[index].items = []int{}
	return monkeys
}

func process(monkeys []monkey, rounds int, relief int, divider int) int {
	for i := 0; i < rounds; i++ {
		for m := range monkeys {
			monkeys = processMonkey(m, monkeys, relief, divider)
		}
	}

	sort.Slice(monkeys, func(i, j int) bool { return monkeys[i].counter > monkeys[j].counter })
	return monkeys[0].counter * monkeys[1].counter
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
	originalMonkeys := make([]monkey, len(monkeys))
	for i := range monkeys {
		originalMonkeys[i] = monkeys[i]
	}

	divider := 1
	for i := range monkeys {
		divider *= monkeys[i].test
	}

	fmt.Println("Part1:", process(monkeys, 20, 3, divider))
	fmt.Println("Part2:", process(originalMonkeys, 10000, 1, divider))
}
