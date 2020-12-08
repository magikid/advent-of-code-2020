package main

import "fmt"

// Day8Solution1 WIP
func Day8Solution1(input []string, done chan string) {
	console := buildCPU(input)
	output := console.Boot()
	done <- fmt.Sprintf("part 1 %v", output)
}

// Day8Solution2 WIP
func Day8Solution2(input []string, done chan string) {
	console := buildCPU(input)
	output := console.CorrectErrors()
	done <- "part 2 " + output
}
