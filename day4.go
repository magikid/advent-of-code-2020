package main

import (
	"fmt"
)

// Day4Solution1 validates passports
func Day4Solution1(input []string, done chan string) {
	records := make(chan string)
	passports := make(chan *passport)

	go fixInput(input, records)
	go buildPassports(records, passports)

	validPassports := 0
	for passport := range passports {
		if passport.MostlyValid() {
			validPassports++
		}
	}

	done <- fmt.Sprintf("part1: found mostly valid %v  passports", validPassports)
}

// Day4Solution2 validates passports
func Day4Solution2(input []string, done chan string) {
	records := make(chan string)
	passports := make(chan *passport)

	go fixInput(input, records)
	go buildPassports(records, passports)

	validPassports := 0
	for passport := range passports {
		if passport.FullyValid() {
			validPassports++
		}
	}

	done <- fmt.Sprintf("part2: found fully valid %v passport", validPassports)
}
