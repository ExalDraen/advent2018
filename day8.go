package main

import (
	"log"
	"os"
)

type node struct {
	children []node
	meta     []int
}

func day8() {
	file, err := os.Open("day8.example")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// read using bufio.Scanner (word delimeter)

	return
}

func parseNodes(inp *file) []node {

}
