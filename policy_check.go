package main

import (
	"fmt"
)

func policyCheck(line string, checker func(string, passwordPolicy) bool) passwordLine {
	var (
		lowCount, highCount int
		letter              rune
		password            string
	)
	fmt.Sscanf(line, "%d-%d %c: %s", &lowCount, &highCount, &letter, &password)
	policy := passwordPolicy{lowCount, highCount, letter}
	validity := checker(password, policy)

	return passwordLine{policy, password, validity}
}
