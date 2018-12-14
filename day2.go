package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func day2() error {
	const filename = "day2.input"
	var twos int
	var threes int

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		isTwo, isThree := day2ParseLine(scanner.Text())
		if isTwo {
			twos++
		}
		if isThree {
			threes++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	log.Printf("#twos: %d, #three: %d", twos, threes)
	checksum := twos * threes
	fmt.Printf("Checksum is: %d\n", checksum)
	return nil
}

func day2Part2() error {
	const filename = "day2.input"
	var ids []string

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ids = append(ids, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var commonRunes []rune
	box1, box2 := day2FindBoxes(ids)

	log.Printf("Boxes with our stuff: %q, %q", box1, box2)
	for idx, c := range box1 {
		if c != box2[idx] {
			continue
		}
		commonRunes = append(commonRunes, c)
	}
	log.Printf("Common characters: '%s'", string(commonRunes))

	return nil
}

// day2FindBoxes finds the two boxes that differ by one character and returns
// their ids
func day2FindBoxes(ids []string) ([]rune, []rune) {
	var idRunes [][]rune

	// Turn lines into slices of runes
	for _, i := range ids {
		idRunes = append(idRunes, []rune(i))
	}

	for _, cur := range idRunes {
		for _, othr := range idRunes {
			if diffByOne(cur, othr) {
				return cur, othr
			}
		}
	}
	return nil, nil
}

// diffByOne returns true if and only if the two rune slices differ by one character
// Slices of different lengths are treated as false.
func diffByOne(x []rune, y []rune) bool {
	var diffs int

	if len(x) != len(y) {
		return false
	}
	for idx, c := range x {
		if c != y[idx] {
			diffs++
		}
	}
	return diffs == 1
}

// day2ParseLine checks a given input line to see if it contains EXACTLY two of the same
// and/or three of the same character
func day2ParseLine(line string) (isTwo bool, isThree bool) {
	var runeCounts map[rune]int
	runeCounts = make(map[rune]int)

	chars := []rune(line)
	for _, c := range chars {
		runeCounts[c] = runeCounts[c] + 1
	}

	for _, v := range runeCounts {
		if v == 2 {
			isTwo = true
		}
		if v == 3 {
			isThree = true
		}
	}

	//log.Printf("2? %t; 3? %t; Rune counts: %q", isTwo, isThree, runeCounts)
	log.Printf("Processed line '%s'  Two? %t.  Three? %t", line, isTwo, isThree)
	return
}
