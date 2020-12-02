package main

type part int

const (
	allParts part = iota
	part1
	part2
)

type userSelections struct {
	day  int
	part part
}

type passwordPolicy struct {
	lowCount  int
	highCount int
	letter    rune
}

type passwordLine struct {
	policy   passwordPolicy
	password string
	valid    bool
}
