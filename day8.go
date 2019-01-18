package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type node struct {
	children []node
	meta     []int
}

func day8() {
	var data []int
	file, err := os.Open("day8.input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// read using bufio.Scanner with word delimiter
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		t := scanner.Text()
		if v, err := strconv.Atoi(t); err == nil {
			data = append(data, v)
		}
	}

	fmt.Printf("Values are: %v\n", data)
	head, size := parseNodes(data)
	fmt.Printf("Head node is %+v, size: %v\n", head, size)

	sum := sumMeta(head)
	fmt.Printf("Sum of metadata is: %v\n", sum)

	val := findNodeValue(head)
	fmt.Printf("Value of root: %v\n", val)

	return
}

// It would be more elegant to parse the string on the fly her
// and cunningly advance the pointer inside the string;
// instead we just chop down slices more and more
func parseNodes(inp []int) (node, int) {
	var n node
	var size int
	nChildren := inp[0]
	nMeta := inp[1]

	size += 2

	// Each successive child gets passed a slice of the input
	// that is shortened by the size of what came before.
	for i := 0; i < nChildren; i++ {
		c, cSize := parseNodes(inp[size:])
		n.children = append(n.children, c)
		size += cSize
	}

	// after the children comes the metadata
	for i := 0; i < nMeta; i++ {
		n.meta = append(n.meta, inp[size])
		size++
	}

	return n, size
}

func sumMeta(start node) (sum int) {

	for _, c := range start.children {
		sum += sumMeta(c)
	}

	for _, m := range start.meta {
		sum += m
	}
	return
}

func findNodeValue(n node) int {
	var val int

	// if no children, value is sum of metadata
	if len(n.children) == 0 {
		for _, m := range n.meta {
			val += m
		}
		return val
	}

	// if children, value is sum of children
	// indicated by metadata.
	for _, m := range n.meta {
		m-- // metadata is 1-index, but arrays are 0-indexed
		if m < len(n.children) {
			val += findNodeValue(n.children[m])
		}
	}
	return val
}
