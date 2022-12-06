package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func readInput(path string) string {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	return string(data)
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("You need to specify a file!")
	}

	filePath := os.Args[1]
	text := readInput(filePath)
	fmt.Println(text)
}
