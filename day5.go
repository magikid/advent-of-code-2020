package main

import "fmt"

// Day5Solution1 finding the max seat id
func Day5Solution1(input []string, done chan string) {
	seatAssignments := make(chan seatAssignment)
	inputChannel := make(chan string, len(input))

	for _, input := range input {
		inputChannel <- input
	}
	close(inputChannel)

	go quickSeatAssignment(inputChannel, seatAssignments)

	var maxSeat seatAssignment
	for seat := range seatAssignments {
		if seat.id > maxSeat.id {
			maxSeat = seat
		}
	}
	done <- fmt.Sprintf("part1: max seat id seen %v", maxSeat.id)
}

// Day5Solution2 WIP
func Day5Solution2(input []string, done chan string) {
	done <- "part2"
}
