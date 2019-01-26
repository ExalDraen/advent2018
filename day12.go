package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	day12file          = "day12.input"
	day12MaxGen        = 100
	day12PartTwoMaxGen = 100
	day12FiftyBillion  = 50000000000
)

// Rule represents a generational surivival rule
type Rule struct {
	LeftLeft   bool
	Left       bool
	AppliesTo  bool // does this rule apply to live or dead pots
	Right      bool
	RightRight bool
}

func day12() {
	state, rules := getState()
	fmt.Printf("Rules: %v, initial state: \n", rules)
	printPotState(state)

	for i := 1; i <= day12MaxGen; i++ {
		state = nextPotGen(state, rules)
		printPotState(state)
	}

	// Sum alive pots. Note that we need to subtract the
	// left padding to get the true pot number
	sum := potSum(state)
	fmt.Printf("\nSum of alive pot values after %v gen: %v", day12MaxGen, sum)
}

func day12PartTwo() {
	state, rules := getState()
	//states := make(map[string]bool)

	var sum int
	var delta int
	for i := 1; i <= day12PartTwoMaxGen; i++ {
		state = nextPotGen(state, rules)
		val := potSum(state)
		delta = val - sum
		sum = val
		fmt.Printf("\n Sum after %v: %5d delta: %5d", i, sum, delta)
	}

	fmt.Printf("\nSum of alive pot values after %v gen: %v", day12PartTwoMaxGen, sum)
	ex := sum + (day12FiftyBillion-day12PartTwoMaxGen)*delta
	fmt.Printf("\nExtrapolated sum after %v gen: %v", day12FiftyBillion, ex)
}

func getState() ([]bool, map[Rule]bool) {
	var state []bool
	rules := make(map[Rule]bool)

	file, err := os.Open(day12file)
	if err != nil {
		log.Fatal("Failed to open file")
	}
	scanner := bufio.NewScanner(file)

	// First line has initial state
	scanner.Scan()
	state = parseInitial(scanner.Text())
	// then an empty line
	scanner.Scan()

	// then the rules
	for scanner.Scan() {
		r, out := lineToRule(scanner.Text())
		rules[r] = out
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return state, rules
}

// Initial line looks like:	initial state: #..#.#..##......###...###
func parseInitial(line string) []bool {
	// left and right pad initial state with MaxGenx2 null values to account
	// for potential growth
	state := make([]bool, day12MaxGen)

	comp := strings.Split(line, ": ")
	for _, r := range comp[1] {
		if r == '#' {
			state = append(state, true)
		} else if r == '.' {
			state = append(state, false)
		}
	}
	pad := make([]bool, day12MaxGen*2)
	state = append(state, pad...)
	return state
}

// Rules are like ...## => #
// return Rule and it's outcome separately so we can easily put it in a map
func lineToRule(line string) (Rule, bool) {
	var r Rule
	var out bool
	comp := strings.Split(line, " => ")
	left := []rune(comp[0])
	right := []rune(comp[1])

	if right[0] == '#' {
		out = true
	}
	if left[0] == '#' {
		r.LeftLeft = true
	}
	if left[1] == '#' {
		r.Left = true
	}
	if left[2] == '#' {
		r.AppliesTo = true
	}
	if left[3] == '#' {
		r.Right = true
	}
	if left[4] == '#' {
		r.RightRight = true
	}

	return r, out
}

func printPotState(state []bool) {
	for _, p := range state {
		if p == true {
			fmt.Print("#")
		} else {
			fmt.Print(".")
		}
	}
	fmt.Printf(" -- sum: %v\n", potSum(state))
}

func nextPotGen(cur []bool, rules map[Rule]bool) []bool {
	nextGen := make([]bool, len(cur))

	for i := 2; i < len(cur)-2; i++ {
		r := Rule{cur[i-2], cur[i-1], cur[i], cur[i+1], cur[i+2]}

		if val, ok := rules[r]; ok == true {
			nextGen[i] = val
		} else {
			nextGen[i] = false
		}
	}
	return nextGen
}

func potSum(state []bool) (sum int) {
	// Sum alive pots. Note that we need to subtract the
	// left padding to get the true pot number
	for i, v := range state {
		if v == true {
			sum += i - day12MaxGen
		}
	}
	return
}
