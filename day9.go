package main

import (
	"fmt"
)

// Day9Solution1 WIP
func Day9Solution1(input []string, output chan string) {
	cipher := makeXmasCipher(input)
	if cipher.Broken() {
		output <- fmt.Sprintf("part1 error found %v", cipher.weaknessTarget)
		return
	}
	output <- "part1 no solution found"
}

// Day9Solution2 WIP
func Day9Solution2(input []string, output chan string) {
	cipher := makeXmasCipher(input)
	weakness := cipher.FindWeakness()
	output <- fmt.Sprintf("part2 weakness found! %v", weakness)
}
