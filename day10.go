package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

const (
	day10file  = "day10.input"
	day10Limit = 1000000
)

// Star Represents a "star" / point of light
type Star struct {
	X  int
	Y  int
	vX int
	vY int
}

// Pos represents an X/Y position
type Pos struct {
	X int
	Y int
}

func day10() {
	var stars []Star

	file, err := os.Open(day10file)
	if err != nil {
		log.Fatal("Failed to open file")
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := lineToStar(scanner.Text())
		stars = append(stars, s)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// progress star's position by 1s at a time and print the result if it's message-like.
	// The star's positions will start out being non message-like, then be message-like, then stop being so
	// again.
	// Thus, once the star's position stops being message like, stop.
	var stopFlag bool
	for i := 0; ; i++ {
		msg := messageLike(stars)
		if msg == true {
			fmt.Printf("Time: %6d", i)
			printStars(stars)
			stopFlag = true
		} else if stopFlag == true {
			break
		}
		stars = progress(stars, 1)
		if i > day10Limit {
			log.Fatal("Exceeded max iterations")
		}
	}
}

func lineToStar(line string) Star {
	var s Star
	//inputs are like: position=< 9,  1> velocity=< 0,  2>
	fmt.Sscanf(line, "position=<%d,%d> velocity=<%d,%d>", &s.X, &s.Y, &s.vX, &s.vY)
	return s
}

// progress advances each star's position by n seconds,
// according to its velocity, then returns the new set of stars
func progress(stars []Star, n int) []Star {
	for i := range stars {
		stars[i].X = stars[i].X + stars[i].vX*n
		stars[i].Y = stars[i].Y + stars[i].vY*n
	}
	return stars
}

// test if a given set of stars represents a message
// heuristic:
// * width (rightmost-leftmost X) must be < #stars
//   each character consumes at least it's width + 2 in stars (L)
//   and a sentence is a sequence of characters + 1 space.
// * height (topmost - bottommost Y) must be < #stars /2
//   this is because the message is > 2 characters and each character consumes at least it's height in stars
func messageLike(stars []Star) bool {
	var minX, maxX = math.MaxInt32, math.MinInt32
	var minY, maxY = math.MaxInt32, math.MinInt32

	for _, s := range stars {
		if s.X < minX {
			minX = s.X
		}
		if s.X > maxX {
			maxX = s.X
		}
		if s.Y < minY {
			minY = s.Y
		}
		if s.Y > maxY {
			maxY = s.Y
		}
	}
	width := Abs(maxX - minX)
	height := Abs(maxY - minY)
	if width > len(stars) {
		return false
	}
	if height > len(stars)/2 {
		return false
	}
	return true
}

func printStars(stars []Star) {
	starMap := make(map[Pos]bool)
	var minX, maxX, minY, maxY int

	for _, s := range stars {
		if s.X < minX {
			minX = s.X
		}
		if s.X > maxX {
			maxX = s.X
		}
		if s.Y < minY {
			minY = s.Y
		}
		if s.Y > maxY {
			maxY = s.Y
		}
		starMap[Pos{s.X, s.Y}] = true
	}

	for i := minY; i <= maxY; i++ {
		fmt.Print("\n")
		for j := minX; j <= maxX; j++ {
			p := Pos{j, i}
			if starMap[p] == true {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
	}
}
