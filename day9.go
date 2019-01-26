package main

import "fmt"

const (
	day9MaxMarble        = 71240
	day9PartTwoMaxMarble = day9MaxMarble * 100
	day9Players          = 478
)

// Node implements a node in a doubly linked list
type Node struct {
	Val  int
	Prev *Node
	Next *Node
}

// Insert inserts a new node after the current one
func (n *Node) Insert(newNode *Node) {
	next := n.Next
	newNode.Next = next
	newNode.Prev = n
	n.Next = newNode
	next.Prev = newNode
}

// Pop removes the node from the linked list and returns it
func (n *Node) Pop() *Node {
	prev := n.Prev
	next := n.Next
	prev.Next = next
	next.Prev = prev
	return n
}

// NewNode instantiates a new isolated node with value val
func NewNode(val int) *Node {
	n := &Node{
		Val: val,
	}
	n.Next = n
	n.Prev = n
	return n
}

func printList(n *Node) {
	cur := n
	for {
		fmt.Printf("%v ", cur.Val)
		cur = cur.Next
		if cur == n {
			return
		}
	}
}

func day9PartTwo() {
	scores := make(map[int]int)
	circle := NewNode(0)
	cur := circle
	player := 0
	for i := 1; i <= day9PartTwoMaxMarble; i++ {
		// fmt.Print("\n")
		// printList(cur)

		// players take one turn each
		player = (i % day9Players)

		if i%23 == 0 {
			scores[player] += i
			// spool back 7 steps, remove marble, make marble
			// immediately clockiwse of removed one new current
			for j := 0; j < 7; j++ {
				cur = cur.Prev
			}
			n := cur.Next
			victim := cur.Pop()
			scores[player] += victim.Val
			cur = n
		} else {
			n := NewNode(i)
			cur = cur.Next
			cur.Insert(n)
			cur = n
		}
	}
	fmt.Printf("\nScores: %+v\n", scores)
	max := 0
	for _, s := range scores {
		if s > max {
			max = s
		}
	}
	fmt.Printf("Max score: %+v\n", max)
}

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

func day9() {
	scores := make(map[int]int)
	circle := []int{0, 1}
	cur := 1
	tar := 0
	player := 1
	for i := 2; i <= day9MaxMarble; i++ {
		// fmt.Printf("[%v] : %v -- cur: %v\n", player, circle, cur)
		// players take one turn each
		player = (i % day9Players)

		// Divisible by 23 -> special case
		// - add what's about to be placed to score
		// - 7 counterclockwise -> remove and add to score
		// current = 1 clockwise from removed marble
		// Note that we need to deal with wrapping here as well, if cur < 7
		if i%23 == 0 {
			scores[player] += i
			tar = cur - 7
			// Deal with wrapping around 0
			if tar < 0 {
				tar = tar + len(circle)
			}
			scores[player] += circle[tar]
			circle = Delete(circle, tar)
			cur = tar
			continue
		}

		// marble needs to be inserted at position 1 and 2 clockwise
		// from current marble, i.e. before position 2.
		// - if the new position wraps over the end of the slice
		// - special case when the new position would be 1 beyond the current slice
		if tar = cur + 2; tar == len(circle) {
			circle = append(circle, i)
		} else {
			tar = tar % len(circle)
			circle = Insert(circle, tar, i)
		}
		cur = tar
	}
	//fmt.Printf("Circle: %v\n", circle)
	fmt.Printf("Scores: %+v\n", scores)
	max := 0
	for _, s := range scores {
		if s > max {
			max = s
		}
	}
	fmt.Printf("Max score: %+v\n", max)
}
