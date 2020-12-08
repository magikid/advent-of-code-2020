package main

import (
	"fmt"
)

// Day7Solution1 WIP
func Day7Solution1(input []string, done chan string) {
	formattedRules := findRulesContaining(input, "shiny gold")
	done <- fmt.Sprintf("part1 found %v bags that can hold a shiny gold", formattedRules.Size())
}

// Day7Solution2 WIP
func Day7Solution2(input []string, done chan string) {
	done <- fmt.Sprintf("part2")
}
