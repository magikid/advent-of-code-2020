package main

import (
	"fmt"
)

// Day7Solution1 WIP
func Day7Solution1(input []string, done chan string) {
	formattedRules := findBagsContaining(input, "shiny gold")
	done <- fmt.Sprintf("part1 found %v bags that can hold a shiny gold", len(formattedRules))
}

// Day7Solution2 WIP
func Day7Solution2(input []string, done chan string) {
	formattedRules := findBagsContaining(input, "shiny gold")
	done <- fmt.Sprintf("part2 %v", len(formattedRules))
}
