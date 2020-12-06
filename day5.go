package main

import (
	"fmt"
	"sort"
)

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
	var seats []seatAssignment

	for _, input := range input {
		seats = append(seats, makeSeatAssignment(input))
	}
	sort.SliceStable(seats, func(i int, j int) bool {
		return seats[i].id < seats[j].id
	})

	lowestID := seats[0].id
	sumOfIDsUpToLowest := (lowestID * (lowestID + 1)) / 2
	highestID := seats[len(seats)-1].id
	sumOfIDsUpToHighest := (highestID * (highestID + 1)) / 2
	sumOfIdsShouldBe := sumOfIDsUpToHighest - sumOfIDsUpToLowest
	seatIDActualSum := 0
	for _, seat := range seats {
		seatIDActualSum += seat.id
	}
	indexOfSeatWithIDOneHigherThanMissingSeat := sumOfIdsShouldBe - seatIDActualSum

	done <- fmt.Sprintf("part2: missing seat id %v", seats[indexOfSeatWithIDOneHigherThanMissingSeat].id-1)
}
