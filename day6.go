package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// DAY6INPUT is the filename for our input
const DAY6INPUT = "day6.input"

// DAY6LIMIT is the max distance a region is allowed to be to pass
const DAY6LIMIT = 10000

// Coord models a "Chronal" coordinate
type Coord struct {
	X int
	Y int
}

func day6Part2() {
	var grid [][]int
	var distSum int
	var regionSize int

	// Read & parse coords; track max x,y
	coords, maxX, maxY := getCoords()

	// create 2d-array(int) for grid:
	// Grid contains (maxY+1) vectors of length (maxX+1)
	// This means that that the first index is the y coordinate,
	// the second one is the x coordinate.
	for i := 0; i <= maxX; i++ {
		grid = append(grid, make([]int, maxY+1))
	}

	// Iterate through each point in array and find distance to all cordinates
	// for col in grid (y coord)
	// - for x in row (x coord)
	// -- for c in coords:
	// --- val(x) = distance(x,y, c)
	// --- Sum of distances += distance
	// -- if (sum of distances) < limit => add to region
	for y, col := range grid {
		for x := range col {
			distSum = 0
			for _, c := range coords {
				d := manhattanDistance(y, x, c)
				// set grid point value to distance
				col[x] = d
				distSum += d
			}
			if distSum < DAY6LIMIT {
				regionSize++
			}
		}
	}

	fmt.Printf("Size of region of points with dist < %v, %v \n", DAY6LIMIT, regionSize)
}

func day6() {
	var area map[int](int) // maps Coord ID to its area
	var grid [][]int
	var minDist int
	var minCoordID int

	area = make(map[int](int))
	// Read & parse coords; track max x,y
	coords, maxX, maxY := getCoords()

	// create 2d-array(int) for grid:
	// Grid contains (maxY+1) vectors of length (maxX+1)
	// This means that that the first index is the y coordinate,
	// the second one is the x coordinate.
	for i := 0; i <= maxX; i++ {
		grid = append(grid, make([]int, maxY+1))
	}

	// Iterate through each point in array and assign to the closest coord based on manhattan distance:
	// for col in grid (y coord)
	// - for row in row (x coord)
	// -- min_dist = distance from (x,y) -> coords[0]
	// -- min_c = coords[1]
	// -- for c in coords:
	// --- if distance (x,y) -> c < min_dist: set min_c = c
	// --- if distance (x,y) = c: set (x,y) to 0 and continue to next x,y
	// -- (x,y) = min_c.id
	// -- area[min_c.id] ++
	for y, col := range grid {
		for x := range col {
			// Start by assuming the first coordinate is closest
			minCoordID = 0
			minDist = -1

			// Skip first coord, since we've already used it above.
			for i := 0; i < len(coords); i++ {
				c := coords[i]
				d := manhattanDistance(y, x, c)
				switch {
				case d < minDist || minDist == -1:
					// closer coord found; flag gridpoint as belonging to coord
					minCoordID = i
					minDist = d
				case d == minDist:
					// gridpoint equidistant with another one; set the grid coordinate to -1 (invalid). Don't need to look at other coordinates
					minCoordID = -1
					break
				}
			}
			col[x] = minCoordID
			area[minCoordID]++
		}
	}
	fmt.Printf("Areas of each coordinate before purge: %+v\n", area)
	// drawGrid(grid, coords)

	// "Infinite" areas do not count. That is, areas that touch the edge
	// of the grid don't count.
	// Set such areas to 0.
	for _, col := range grid {
		// First row: (x, 0) column
		area[col[0]] = 0
		// Last row: (x, maxY) row
		area[col[maxY]] = 0
	}
	for _, eY := range grid[0] {
		// First column: (0, y)
		area[eY] = 0
	}
	for _, eY := range grid[maxX] {
		// Last column: (maxX, y)
		area[eY] = 0
	}

	fmt.Printf("Areas of each coordinate after purge: %+v\n", area)

	// Iterate through area map to find largest area:
	// for k,v in area:
	//   if k.x/k.y on edge: skip
	//   if v> max_v : update max_v, max_c
	// print max_c
	var maxArea, maxCoordID int
	for k, v := range area {
		if v > maxArea {
			maxArea = v
			maxCoordID = k
		}
	}

	fmt.Printf("Max area: %v; Coord: %v\n", maxArea, maxCoordID)
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

func drawGrid(grid [][]int, coords []Coord) {
	var val rune
	tr := Transpose(grid)
	for x, eX := range tr {
		for y, eY := range eX {
			if eY == -1 {
				val = rune('.')
			} else {
				val = rune('a') + rune(eY)
			}
			for ic, c := range coords {
				if x == c.Y && y == c.X {
					val = rune('a') + rune(ic) - rune(32) // -32 means uppercase
				}
			}
			fmt.Printf("%v", string(val))
		}
		fmt.Printf("\n")
	}
}

// Transpose transpoes a rectangular 2d array
func Transpose(a [][]int) [][]int {
	nCol := len(a)
	nRow := len(a[0])

	tr := make([][]int, nRow)
	for x := 0; x < nRow; x++ {
		tr[x] = make([]int, nCol)
	}
	for i := 0; i < nCol; i++ {
		for j := 0; j < nRow; j++ {
			tr[j][i] = a[i][j]
		}
	}
	return tr
}
