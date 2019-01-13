package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
)

const day7workers = 2
const day7baseLength = 0

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

	file, err := os.Open("day7.example")
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
		// log.Print("Couldn't find an option!")
		return "-1"
	}

	// fmt.Printf("Next step: %v\n", first)
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

func day7Part2() {

	// make channel for ticks
	quit := make(chan bool)

	var lock sync.Mutex
	tick := 0
	steps := getSteps()
	stepTotal := len(steps)
	var execOrder []string
	var actions []chan string

	for i := 0; i < day7workers; i++ {
		action := make(chan string)
		actions = append(actions, action)
		go work(steps, &execOrder, &lock, action, quit)
	}

	for {
		// Send no work outstanding -> quit
		// otherwise -> send tick
		if len(execOrder) == stepTotal {
			for i := 0; i < day7workers; i++ {
				quit <- true
			}
			break
		}
		var finished []string // Don't need this if workers marks jobs done

		fmt.Printf("\nTick %v: ", tick)
		for i := 0; i < day7workers; i++ {
			c := <-actions[i]
			// Uppercase letters mean job is finishing
			if c != "." && strings.ToUpper(c) == c {
				finished = append(finished, c)
			}
			fmt.Printf(" %v", c)
		}
		// mark finished things as finished <<< do this in work stage
		execOrder = append(execOrder, finished...)
		tick++
		if tick > 100 {
			log.Fatal("Infinite loop?")
		}
	}
	fmt.Printf("\nWork complete. Took %v ticks", tick)
}

// day 7 part two
//
// action status: "." = idle
//                "X" = working on X
//
// map(step: status) stepStatus
// mapLock Mutex
//
// get steps
//
// make channel(string) for action status
// for i up to worker limit:
// - go work(AllSteps, stepStatus, mapLock, channel)
//
// tick := 0
// while true:
// tick++
// -  for i up to worker limit
// - v, ok := <- ch
// - if ok := print value
// - if closed, break

func work(steps map[string]([]string), executionOrder *[]string, stepLock *sync.Mutex, action chan string, quit chan bool) {
	for {
		// wait to receive tick

		// wait to send data
		select {
		case <-quit:
			return
		default:
			stepLock.Lock()
			c := findCandidate(*executionOrder, steps)
			if c == "-1" {
				stepLock.Unlock()
				action <- "."
			} else {
				delete(steps, c)
				stepLock.Unlock()
				// Work on step:
				// Len -1: send "work in progress" -> lowercase letter
				// Len: send "work completing" -> uppercase
				for j := 0; j < workLength(c)-1; j++ {
					action <- strings.ToLower(c)
				}
				action <- c
				// update execOrder
			}
		}

	}
}

// func work(AllSteps, stepStatus, mapLock, channel)
// while true:
// - if all steps done:
// -- close(channel).
// -- Return.
// - lock mapLock
// - findCandidate(stepStatus, AllSteps)
// - if no candidate: unlock mapLock. channel <- idle
// - if candidate:
// -- mark as in progress
// -- unlock mapLock
// -- calculate work length
// -- for i -> length of work
// --- channel <- "working on stuff"
//

// calculates the length of work in "seconds"
// A=1, B=2, ...
func workLength(step string) (length int) {
	r := []rune(step)[0]
	length = day7baseLength + int(r) - 'A' + 1
	return
}
