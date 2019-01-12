package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func day7() {
	steps := getSteps()
	var execOrder []string

	for i := 0; i < len(steps); i++ {
		execOrder = append(execOrder, findCandidate(execOrder, steps))
	}

	fmt.Printf("Final execution order:\n")
	for _, e := range execOrder {
		fmt.Printf("%v", e)
	}
}

func getSteps() map[string]([]string) {
	// instr is a map of step: steps it comes after
	var instr map[string]([]string)

	instr = make(map[string]([]string))

	file, err := os.Open("day7.input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		line := scanner.Text()
		if len(line) > 0 {
			var step, bef string

			if _, err := fmt.Sscanf(line, "Step %s must be finished before step %s can begin.", &step, &bef); err != nil {
				log.Fatal(err)
			}
			instr[bef] = append(instr[bef], step)
			// `step`might not have any requirements and thus never appear as bef; this ensure the key is nevertheless present.
			if elem, ok := instr[step]; ok == false {
				instr[step] = elem
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return instr
}

func findCandidate(executionOrder []string, steps map[string]([]string)) string {
	var options []string
	var skip bool

	for step, reqs := range steps {
		skip = false
		// if already executed, skip
		for _, exec := range executionOrder {
			if step == exec {
				skip = true
				break
			}
		}
		if skip == true {
			continue
		}

		// No requirements, so we're done
		if len(reqs) == 0 {
			options = append(options, step)
			continue
		}
		// Nothing has been executed and reqs are not 0. So this step can't possibly be allowed yet, skip
		if len(executionOrder) == 0 {
			continue
		}
		// we have executed some steps and have some reqs. Check if reqs have been executed
		if subset(reqs, executionOrder) {
			options = append(options, step)
		}
	}

	// pick alphabetically first
	var first string
	if len(options) > 0 {
		first = options[0]
		for _, o := range options {
			if []rune(o)[0] < []rune(first)[0] {
				first = o
			}
		}
	} else {
		log.Fatal("Couldn't find an option!")
	}

	fmt.Printf("Next step: %v\n", first)
	return first
}

// subset returns true if the first array is completely
// contained in the second array. There must be at least
// the same number of duplicate values in second as there
// are in first.
// Cf. https://stackoverflow.com/a/18879994
func subset(first, second []string) bool {
	set := make(map[string]int) // count occurences

	for _, val := range second {
		set[val]++
	}

	for _, e := range first {
		if count, ok := set[e]; !ok {
			return false
		} else if count < 1 {
			return false
		} else {
			set[e]--
		}
	}
	return true
}
