package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"unicode"
)

// INPUT filename
const INPUT = "day5.input"

// ALPHABET represents all possible polymers to delete
const ALPHABET = "abcdefghijklmnopqrstuvwxyz"

func day5() {
	// read input
	content, err := ioutil.ReadFile(INPUT)
	if err != nil {
		log.Fatal(err)
	}
	runes := []rune(string(content))
	fmt.Printf("Initial elements:   %v\n", string(runes))
	fmt.Printf("Initial length:     %v\n", len(runes))

	// iterate over runes:
	// if current does not destroy next; copy rune to output
	// if current destroys next; continue
	result, reactions := react(runes)

	fmt.Printf("Reacted %v times in total\n", reactions)
	fmt.Printf("Remaining elements: %v\n", string(result))
	fmt.Printf("#Remaining elements: %v\n", len(result))
}

func day5Part2() {
	// read input
	content, err := ioutil.ReadFile(INPUT)
	if err != nil {
		log.Fatal(err)
	}

	candidates := []rune(ALPHABET)
	minElem := candidates[0]
	min := len(string(content))

	// for each candidate rune:
	// - strip candidate from input
	// - react fully
	// - check if shortest
	for _, c := range candidates {
		cleaned := strings.Map(func(r rune) rune {
			if r == c || r == unicode.ToUpper(c) {
				return -1
			}
			return r
		}, string(content))
		runes := []rune(cleaned)

		result, _ := react(runes)
		if len(result) < min {
			min = len(result)
			minElem = c
		}
	}

	fmt.Printf("Removing %v is best; remaining length: %v\n", string(minElem), min)
}

// IsAntirune returns whether or not two runes are the same kind
// but opposite "polarity" such that they "destroy" each other
func IsAntirune(l rune, r rune) bool {
	// aA, Aa, but not aa and AA
	return (unicode.ToUpper(l) == unicode.ToUpper(r) && (unicode.IsUpper(l) && unicode.IsLower(r) || unicode.IsLower(l) && unicode.IsUpper(r)))
}

func react(input []rune) ([]rune, int) {
	var output []rune
	var reactions int

	for i := 0; i < len(input); i++ {
		if i < len(input)-1 && IsAntirune(input[i], input[i+1]) {
			i++
			reactions++
		} else {
			output = append(output, input[i])
		}
	}
	// fmt.Printf("Reacted %5d times\n", reactions)
	if reactions > 0 {
		var r int
		output, r = react(output)
		reactions += r
	}
	return output, reactions
}
