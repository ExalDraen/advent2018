package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func day1() error {
	const filename = "day1.input"
	var freq int

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		freq += day1ParseLine(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Final frequency is: %d\n", freq)
	return nil
}

func day1PartTwo() error {
	const filename = "day1.input"
	// Set-like object. See https://stackoverflow.com/a/10486196
	var freqChanges []int
	var freqs map[int]struct{}
	freqs = make(map[int]struct{})

	// First build up a list of frequency changes
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		change := day1ParseLine(scanner.Text())
		freqChanges = append(freqChanges, change)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Loop until one frequency repeats
	var freq int
	for {
		for _, c := range freqChanges {
			freq += c
			if _, ok := freqs[freq]; ok == true {
				log.Printf("Found duplicate frequency: %d", freq)
				return nil
			}
			freqs[freq] = struct{}{}
		}
	}
}

// day1ParseLine takes an input line and returns the corresponding frequency change
func day1ParseLine(line string) int {
	log.Println(line)

	val, err := strconv.Atoi(line)
	if err != nil {
		log.Fatal(err)
	}
	return val
}
