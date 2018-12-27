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

// Interval represents a time interval
type Interval struct {
	StartTime time.Time
	EndTime   time.Time
}

// Elapsed returns the duration between the end and start of the interval
func (i Interval) Elapsed() time.Duration {
	return i.EndTime.Sub(i.StartTime)
}

func day4() {
	var laziestGuard int
	var laziestMinute int
	var sum, max int64
	var sleepMins [60]int

	e := getEvents()
	// for _, v := range e {
	// 	fmt.Printf("%+v\n", v)
	// }

	// find guard with most sleepMins:
	// - loop over guards.
	// - sum Mins for each sleep duration
	// - if cur > max, update max
	work := eventsToWorkLog(e)
	for k, v := range work {
		// fmt.Printf("%v: %+v\n", k, v)

		sum = 0
		for _, elem := range v {
			sum += elem.Elapsed().Nanoseconds()
		}
		if sum > max {
			max = sum
			laziestGuard = k
		}
	}
	fmt.Printf("Laziest guard is: #%v\n", laziestGuard)

	// find max sleep minute:
	// slice[0-60]: int
	// for each period:
	// add being-> end to slice
	// finally find max
	for _, r := range work[laziestGuard] {
		for i := r.StartTime.Minute(); i < r.EndTime.Minute(); i++ {
			sleepMins[i]++
		}
	}
	fmt.Printf("Amount of times slept on each minute: %v\n", sleepMins)
	for idx, val := range sleepMins {
		// inefficient, extra lookup. But I'm being lazy :)
		if val > sleepMins[laziestMinute] {
			laziestMinute = idx
		}
	}
	fmt.Printf("Laziest minute is: %v\n", laziestMinute)
	fmt.Printf("Laziest minute x Laziest Guard: %v\n", laziestGuard*laziestMinute)
}

func day4Part2() {
	fmt.Println("Part Two")
	var laziestGuard int
	var laziestMinute int
	var max int

	e := getEvents()
	// for _, v := range e {
	// 	fmt.Printf("%+v\n", v)
	// }

	// find guard most asleep on one particular minute
	// - loop over guards.
	// - calculate "sleep spectrum" (#times asleep on given minute)
	// - check if sleep spectrum max > current max.
	// - if spectrum max > cur max, update cur max, update guard id
	work := eventsToWorkLog(e)
	for k, v := range work {
		// fmt.Printf("%v: %+v\n", k, v)
		var sleepMins [60]int

		for _, elem := range v {
			for i := elem.StartTime.Minute(); i < elem.EndTime.Minute(); i++ {
				sleepMins[i]++
			}
		}

		for idx, val := range sleepMins {
			if val > max {
				max = val
				laziestMinute = idx
				laziestGuard = k
			}
		}
	}
	fmt.Printf("Laziest guard is: #%v\n", laziestGuard)
	fmt.Printf("Laziest minute is: %v\n", laziestMinute)
	fmt.Printf("Laziest minute x Laziest Guard: %v\n", laziestGuard*laziestMinute)
}

func getEvents() []Event {
	const filename = "day4.input"
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

// eventsToWorkLog transforms a list of events into a map of guard id: sleep ranges
func eventsToWorkLog(events []Event) map[int]([]Interval) {
	// parse event lines into a map of guard id: sleepRanges
	// - loop over events
	// - if new guard; if sleeping, add sleep period for previous guard. switch cur -> new guard. mark curTime = time
	// - if wake up; sleeping = false. add sleep range: curTime until time. Update curTime = time
	// - if sleep. sleeping = true. update curTime
	var workLog map[int]([]Interval)
	workLog = make(map[int][]Interval)
	var sleeping bool
	var curGuard int
	var lastTime time.Time

	lastTime = events[0].EventTime
	for _, e := range events {
		switch e.EventType {
		case SLEEP:
			sleeping = true
		case WAKE:
			if sleeping == true {
				workLog[curGuard] = append(workLog[curGuard], Interval{StartTime: lastTime, EndTime: e.EventTime})
			}
			sleeping = false
		case GUARD:
			if sleeping == true {
				workLog[curGuard] = append(workLog[curGuard], Interval{StartTime: lastTime, EndTime: e.EventTime})
			}
			curGuard = e.GuardID
		}
		lastTime = e.EventTime
	}

	return workLog
}
