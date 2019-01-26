package main

import (
	"fmt"
	"math"
)

const (
	day11Serial = 9424
	day11Grid   = 300
)

func day11() {
	maxPow := math.MinInt32
	var maxX, maxY, maxN int

	for i := 1; i <= day11Grid-3+1; i++ {
		for j := 1; j <= day11Grid-3+1; j++ {
			if p := squarePower(i, j, 3); p > maxPow {
				maxPow = p
				maxX = i
				maxY = j
			}
		}
	}
	fmt.Printf("Max power %v at coords (%v, %v)", maxPow, maxX, maxY)

	// Part 2
	// need to precalculate cell powers otherwise it's too slow
	var grid [day11Grid + 1][day11Grid + 1]int
	for i := 1; i <= day11Grid; i++ {
		for j := 1; j <= day11Grid; j++ {
			grid[i][j] = power(i, j)
		}
	}
	for n := 1; n <= day11Grid; n++ {
		fmt.Printf("\n Trying N=%v", n)
		for i := 1; i <= day11Grid-n+1; i++ {
			for j := 1; j <= day11Grid-n+1; j++ {

				if p := sumGrid(i, j, n, &grid); p > maxPow {
					maxPow = p
					maxN = n
					maxX = i
					maxY = j
				}
			}
		}
	}
	fmt.Println("----")
	fmt.Printf("Max power %v at coords (%v, %v) with grid size %v", maxPow, maxX, maxY, maxN)
}

// calculates the given fuel cell's power
// from the description:
// Find the fuel cell's rack ID, which is its X coordinate plus 10.
// Begin with a power level of the rack ID times the Y coordinate.
// Increase the power level by the value of the grid serial number (your puzzle input).
// Set the power level to itself multiplied by the rack ID.
// Keep only the hundreds digit of the power level (so 12345 becomes 3; numbers with no hundreds digit become 0).
// Subtract 5 from the power level.
func power(x, y int) (pow int) {
	rack := x + 10
	pow = rack * y
	pow += day11Serial
	pow *= rack
	pow = pow % 1000
	pow = pow / 100
	pow -= 5
	return
}

// calculate total power of n*n square whose upper
// left corner is at x,y
func squarePower(x, y, n int) (pow int) {
	for i := x; i < x+n; i++ {
		for j := y; j < y+n; j++ {
			pow += power(i, j)
		}
	}
	return
}

func sumGrid(x, y, n int, grid *[day11Grid + 1][day11Grid + 1]int) (pow int) {
	for i := x; i < x+n; i++ {
		for j := y; j < y+n; j++ {
			pow += grid[i][j]
		}
	}
	return
}
