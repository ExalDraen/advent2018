package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
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
	// the full set of fabric. Fabric[i][j] is the ith column, jth row
	// 0s are unused areas of fabric
	var fabric [LEN][LEN]uint8

	// parse claims into slice of claim structs
	claims := getClaims()

	// Mark occupied areas by incrementing value
	for _, c := range claims {
		for i := c.X; i < c.X+c.Width; i++ {
			for j := c.Y; j < c.Y+c.Height; j++ {
				fabric[i][j]++
			}
		}
	}

	drawFabric(fabric, 80, 12)

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

func day3Part2() {
	// Represents the fabric. fabric[i][j] is ith column, jth row
	var fabric [LEN][LEN]int
	claims := getClaims()

	// Keep track of conflicting claims with a simple slice
	var conflicted []bool
	conflicted = make([]bool, len(claims)+1)

	// Mark used areas of fabric with ID of claim
	// In case of conflict, record both claims as conflicting
	for _, c := range claims {
		for i := c.X; i < c.X+c.Width; i++ {
			for j := c.Y; j < c.Y+c.Height; j++ {
				if v := fabric[i][j]; v != 0 {
					conflicted[v] = true
					conflicted[c.ID] = true
				}
				fabric[i][j] = c.ID
			}
		}
	}

	// Now find the non-conflicting claims (should only be one!)
	for idx, v := range conflicted {
		if v == false {
			fmt.Printf("Claim %d does not conflict!\n", idx)
		}
	}
}

// drawFabric "draws" the fabric up to width and height
func drawFabric(fabric [LEN][LEN]uint8, width int, height int) {
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if v := fabric[j][i]; v == 0 {
				fmt.Print(".")
			} else {
				fmt.Print(fabric[j][i])
			}
		}
		fmt.Print("\n")
	}
}

func getClaims() []Claim {
	const filename = "day3.input"
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
	// Claims are of the form #3 @ 5,5: 2x2
	// #<claim ID> @ <dist-left>,<dist-right>: <width>x<height>
	r := regexp.MustCompile("#(?P<ID>[[:digit:]]+) @ (?P<X>[[:digit:]]+),(?P<Y>[[:digit:]]+): (?P<Width>[[:digit:]]+)x(?P<Height>[[:digit:]]+)")
	n1 := r.SubexpNames()
	r2 := r.FindAllStringSubmatch(line, -1)[0]

	md := map[string]string{}
	for i, n := range r2 {
		md[n1[i]] = n
	}

	if val, err := strconv.Atoi(md["ID"]); err == nil {
		c.ID = val
	}
	if val, err := strconv.Atoi(md["X"]); err == nil {
		c.X = val
	}
	if val, err := strconv.Atoi(md["Y"]); err == nil {
		c.Y = val
	}
	if val, err := strconv.Atoi(md["Height"]); err == nil {
		c.Height = val
	}
	if val, err := strconv.Atoi(md["Width"]); err == nil {
		c.Width = val
	}
	return c
}
