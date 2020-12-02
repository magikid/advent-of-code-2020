package main

type part int

const (
	Part1 part = iota
	Part2
)

type Parter interface {
	Part() part
}

func (p part) Part() part {
	return p
}

type Puzzle struct {
	PuzzleInput []string
	completed   bool
	part        part
}
