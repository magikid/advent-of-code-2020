package main

type part int

const (
	allParts part = iota
	part1
	part2
)

type parter interface {
	Part() part
}

func (p part) Part() part {
	return p
}

type userSelections struct {
	day  int
	part part
}

type passwordPolicy struct {
	lowCount  int
	highCount int
	letter    string
}

type passwordLine struct {
	policy   passwordPolicy
	password string
	valid    bool
}
