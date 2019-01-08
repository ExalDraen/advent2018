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
	var area map[int](int) // maps Coord ID to its area
	var grid [][]int
	var minDist int
	var minCoordID int

	area = make(map[int](int))
	coords, maxX, maxY := getCoords()

	// create grid of size x+1 * y+1
	for i := 0; i <= maxX; i++ {
		grid = append(grid, make([]int, maxY))
	}
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
	for x, eX := range grid {
		for y := range eX {
			// prime minimal values
			minCoordID = 0
			minDist = manhattanDistance(x, y, coords[0])

			// find closest coord to gridpoint
			for i, c := range coords {
				d := manhattanDistance(x, y, c)
				switch {
				case d < minDist:
					// closer coord found; flag gridpoint as belonging to coord
					minCoordID = i
					minDist = d
					eX[y] = i
					continue
				case d == minDist:
					// gridpoint equidistant; discount
					eX[y] = -1
					continue
				}
			}
			area[minCoordID]++
		}
	}

	// Iterate through area map to find largest area:
	// for k,v in area:
	//   if k.x/k.y on edge: skip
	//   if v> max_v : update max_v, max_c
	// print max_c
	var maxArea, maxCoordID int
	for k, v := range area {
		c := coords[k]
		if v > maxArea {
			v = maxArea
			maxCoordID = k
		}
	}
}

func getCoords() ([]Coord, int, int) {
	var coords []Coord
	var maxX, maxY int

	file, err := os.Open(DAY6INPUT)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		c := lineToCoord(scanner.Text())
		coords = append(coords, c)
		if c.X > maxX {
			maxX = c.X
		}
		if c.Y > maxY {
			maxY = c.Y
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return coords, maxX, maxY
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

// manhattanDistance calculates the Manhattan distance between (x,y) and a Coord
func manhattanDistance(x int, y int, c Coord) int {
	return Abs(x-c.X) + Abs(y-c.Y)
}

// Abs computes the absolute integer value
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
