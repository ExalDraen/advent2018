package main

import "fmt"

const day9Players = 9
const day9MaxMarble = 25

// Insert value x into slice v at position i
// shifting everything up by one (i.e. it's inserted "before" i)
func Insert(v []int, i, x int) []int {
	v = append(v, 0)
	copy(v[i+1:], v[i:])
	v[i] = x
	return v
}

// Delete elements at i from slice v
func Delete(v []int, i int) []int {
	return append(v[:i], v[i+1:]...)
}

// Convert a circle index into a regular array index,
// including -ve indices and wrapping.
func getCircleIndex(i, n int) int {
	switch {
	case i <= n && i > 0:
		return i
	case i > n:
		return i % n
	case i < 0:
		return n - (-i % n)
	}
	return i
}

func day9() {
	for j := 0; j < 50; j++ {
		fmt.Printf("%v -> %v\n", j, getCircleIndex(j, 9))
		fmt.Printf("%v -> %v\n", -j, getCircleIndex(-j, 9))
	}
	// create slice len = maxMarble+1
	scores := make(map[int]int)
	circle := []int{0, 1}
	cur := 1
	tar := 0
	player := 0
	for i := 2; i <= day9MaxMarble; i++ {
		// players take one turn each
		player = (i % day9Players)

		// Divisible by 23 -> special case
		// add 23 to score
		// 7 counterclockwise -> remove and add to score
		// current = 1 clockwise from removed marble
		if i%23 == 0 {
			scores[player] += i
			scores[player] += circle[i-7]
			circle = Delete(circle, i-7)
			cur = i - 7 + 1
		}
		// marble needs to be inserted at position 1 and 2 clockwise
		// from current marble, i.e. before position 2
		tar = (cur + 2) % len(circle)
		circle = Insert(circle, tar, i)
		cur = tar
		fmt.Printf("[%v] : %v\n", player, circle)
	}
	fmt.Printf("Circle: %v\n", circle)
}
