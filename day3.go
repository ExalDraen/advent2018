package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// LEN is the length of the fabric square
const LEN = 1250

// Claim represents an elves' claim to fabric
type Claim struct {
	ID     int
	X      int // from left
	Y      int // from top
	Width  int
	Height int
}

func day3() {
	// the full set of fabric
	// 0s are unused areas of fabric
	var fabric [LEN][LEN]uint8

	// parse claims into slice of claim structs
	claims := getClaims()

	// Mark occupied areas by incrementing value
	for _, c := range claims {
		for i := c.X; i <= c.X+c.Width; i++ {
			for j := c.Y; i <= c.Y+c.Height; j++ {
				fabric[i][j]++
			}
		}
	}

	// count number of conflicted square inches
	conflicts := 0
	for i := 0; i < LEN; i++ {
		for j := 0; j < LEN; j++ {
			if fabric[i][j] > 1 {
				conflicts++
			}
		}
	}
	fmt.Printf("Conflicted square inches: %d\n", conflicts)
}

func getClaims() []Claim {
	const filename = "day3.example"
	var claims []Claim

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		claims = append(claims, lineToClaim(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return claims
}

func lineToClaim(line string) Claim {
	var c Claim
	// TODO: parse line into a Claim
	return c
}
