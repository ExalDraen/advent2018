package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sync"
)

const day7workers = 5
const day7baseLength = 60
const day7TickLimit = 10000

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

// day 7 part two
func day7Part2() {
	var lock sync.Mutex
	tick := 0
	steps := getSteps()
	stepTotal := len(steps)
	var execOrder []string
	var actions []chan string
	// channel for sending ticks to workers
	var tickers []chan bool

	for i := 0; i < day7workers; i++ {
		action := make(chan string)
		actions = append(actions, action)
		ticker := make(chan bool) // true: go; false: stop
		tickers = append(tickers, ticker)
		go work(steps, &execOrder, &lock, action, ticker)
	}

	for {
		// 2 phases
		// 1) receive work update
		// 2) send tick to advance workers to next turn
		fmt.Printf("\nTick %5d: ", tick)
		for i := 0; i < day7workers; i++ {
			c := <-actions[i]
			fmt.Printf(" %v", c)
		}
		tick++
		// Send no work outstanding -> quit
		// otherwise -> send tick to have workers do work this tick
		// This synchronization is necessary to make sure work done by
		// workers is only counted on the next tick.
		if len(execOrder) == stepTotal {
			for i := 0; i < day7workers; i++ {
				tickers[i] <- false
			}
			break
		} else {
			for i := 0; i < day7workers; i++ {
				tickers[i] <- true
			}
		}
		if tick > day7TickLimit {
			log.Fatal("Exceeded tick limit. Infinite loop?")
		}
	}
	fmt.Printf("\nWork complete. Took %v ticks", tick)
	fmt.Printf("\nExecution order was: %v", execOrder)
}

func work(steps map[string]([]string), executionOrder *[]string, stepLock *sync.Mutex, action chan string, ticker chan bool) {
	for {
		stepLock.Lock()
		c := findCandidate(*executionOrder, steps)
		if c == "-1" {
			stepLock.Unlock()
			action <- "."
		} else {
			delete(steps, c)
			stepLock.Unlock()

			// Do all but last step of the work
			for j := 0; j < workLength(c)-1; j++ {
				action <- c
				<-ticker
			}
			// Do the last step of the work and update completed steps
			action <- c
			*executionOrder = append(*executionOrder, c)
		}
		// The ticker channel is used to synchronize the end of the workers.
		// This is to make sure any updates to executionOrder are only taken into
		// account on the next "tick"
		proceed := <-ticker
		if !proceed {
			return
		}
	}
}

// calculates the length of work in "seconds"
// A=1, B=2, ...
func workLength(step string) (length int) {
	r := []rune(step)[0]
	length = day7baseLength + int(r) - 'A' + 1
	return
}
