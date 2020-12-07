package main

import (
	"fmt"
)

// Day6Solution1 finding answers where anyone in group answered yes
func Day6Solution1(input []string, done chan string) {
	records := make(chan string)
	declarationForms := make(chan customsForm)

	go fixInput(input, records)
	go getForms(records, declarationForms)

	counter := 0
	for form := range declarationForms {
		counter += len(form.AnyoneAnsweredYes())
	}

	done <- fmt.Sprintf("part1 anyone in group answered yes sum %v", counter)
}

// Day6Solution2 finding answers where everyone in group answered yes
func Day6Solution2(input []string, done chan string) {
	records := make(chan string)
	declarationForms := make(chan customsForm)

	go fixInput(input, records)
	go getForms(records, declarationForms)

	counter := 0
	for form := range declarationForms {
		counter += len(form.EveryoneAnsweredYes())
	}

	done <- fmt.Sprintf("part2 everyone in group answered yes sum %v", counter)
}
