package main

import "fmt"

const day9Players = 9
const day9MaxMarble = 25

// Insert value x into slice v at position i
func Insert(v []int, i, x int) []int {
	v = append(v, 0)
	copy(v[i+1:], v[i:])
	v[i] = x
	return v
}

func day9() {
	// create slice len = maxMarble+1
	circle := []int{0}
	cur := 0
	for i := 0; i <= day9MaxMarble; i++ {
		// work out target position where marble will be inserted
		tar := (cur + 1) % len(circle)
		circle = Insert(circle, tar, i)
	}
	fmt.Printf("Circle: %v", circle)
}
