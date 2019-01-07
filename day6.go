package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

// DAY6INPUT is the filename for our input
const DAY6INPUT = "day6.example"

// Coord models a "Chronal" coordinate
type Coord struct {
	X int
	Y int
}

func day6() {

	// Outline:
	// map( coord id -> area).
	// Read & parse coords; track max x,y
	// create 2d-array(int) for grid using max x, y; size= x+1, y+1
	// Iterate through each point in array and assign to a coord based on manhattan distance:
	// for x in row
	// - for y in col
	// -- min_dist = distance(x,y) -> coords[1]
	// -- min_c = coords[1]
	// -- for c in coords:
	// --- if distance (x,y) -> c < min_dist: set min_c = c
	// --- if distance (x,y) = c: set (x,y) to 0 and continue to next x,y
	// -- (x,y) = min_c.id
	// -- area[min_c.id] ++
	//
	// Iterate through area map to find largest area:
	// for k,v in area:
	//   if k.x/k.y on edge: skip
	//   if v> max_v : update max_v, max_c
	// print max_c
}

func getCoords() []Coord {
	var coords []Coord

	file, err := os.Open(DAY6INPUT)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		coords = append(coords, lineToCoord(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return coords
}

func lineToCoord(line string) Coord {
	var c Coord

	components := strings.Split(line, ", ")
	if val, err := strconv.Atoi(components[0]); err == nil {
		c.X = val
	}
	if val, err := strconv.Atoi(components[1]); err == nil {
		c.Y = val
	}
	return c
}
