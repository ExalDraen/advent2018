package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

// GUARD = "Guard starts shift"
const GUARD = 0

// WAKE = "guard wakes"
const WAKE = 1

// SLEEP = "guard sleeps"
const SLEEP = 2

// Event represents a single event in the guard log
type Event struct {
	EventTime time.Time
	EventType uint8 // 0 = guard start, 1 = wake up, 2 = sleep
	GuardID   int   // 0 means no guard
}

// Shift represents one guards' shift
type Shift struct {
	GuardID     int
	StartTime   time.Time
	sleepRanges []time.Duration
}

func day4() {
	e := getEvents()
	for _, v := range e {
		fmt.Printf("%q\n", v)
	}
	// parse event lines into a map of guard id: sleepRanges
	// - loop over events
	// - if new guard; if sleeping, add sleep period for previous guard. switch cur -> new guard. mark curTime = time
	// - if wake up; sleeping = false. add sleep range: curTime until time. Update curTime = time
	// - if sleep. sleeping = true. update curTime

	// find guard with most sleepMins:
	// - loop over guards.
	// - sum Mins for each sleep duration
	// - if cur > max, update max

	// find max sleep minute:
	// slice[0-60]: int
	// for each period:
	// add being-> end to slice
	// finally find max

}

func getEvents() []Event {
	const filename = "day4.example"
	var events []Event

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		events = append(events, lineToEvent(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// sort
	sort.Slice(events, func(i, j int) bool {
		return events[i].EventTime.Before(events[j].EventTime)
	})

	return events
}

func lineToEvent(text string) (e Event) {
	// example lines:
	// [1518-11-01 00:00] Guard #10 begins shift
	// [1518-11-03 00:29] wakes up
	components := strings.Split(text, "]")

	// timeForm shows by example how the reference time would be represented
	// ref time is: Mon Jan 2 15:04:05 -0700 MST 2006
	const timeForm = "[2006-01-02 15:04"
	t, err := time.Parse(timeForm, components[0])
	if err != nil {
		log.Fatal(err)
	}
	e.EventTime = t

	rg := regexp.MustCompile("Guard #([[:digit:]]+)")
	rw := regexp.MustCompile("wakes up")

	if rwMatch := rw.FindStringIndex(components[1]); rwMatch != nil {
		e.EventType = WAKE
		return
	}

	if rgMatch := rg.FindAllStringSubmatch(components[1], -1); rgMatch != nil {
		e.EventType = GUARD
		if val, err := strconv.Atoi(rgMatch[0][1]); err == nil {
			e.GuardID = val
		}
		return
	}
	// Fall-through must be a sleep event
	e.EventType = SLEEP
	return
}

func eventsToShifts(events []Event) []Shift {
	var cur Shift
	//
	return nil
}
