package main

import (
	"strconv"
	"strings"
)

type seatAssignment struct {
	row                  int
	column               int
	id                   int
	binarySpacePartition string
}

func makeBinary(input string, zero rune, one rune) int {
	input = strings.ReplaceAll(input, string(zero), "0")
	input = strings.ReplaceAll(input, string(one), "1")

	newInt, err := strconv.ParseInt(input, 2, 32)
	check(err)

	return int(newInt)
}

func makeSeatAssignment(binarySpacePartition string) seatAssignment {
	seat := seatAssignment{binarySpacePartition: binarySpacePartition}
	if len(binarySpacePartition) < 10 {
		return seat
	}

	seat.row = makeBinary(binarySpacePartition[0:7], 'F', 'B')
	seat.column = makeBinary(binarySpacePartition[7:10], 'L', 'R')
	seat.id = (seat.row * 8) + seat.column

	return seat
}

func quickSeatAssignment(inputs <-chan string, seats chan<- seatAssignment) {
	for input := range inputs {
		seats <- makeSeatAssignment(input)
	}
	close(seats)
}
