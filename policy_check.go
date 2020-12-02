package main

import (
	"strconv"
	"strings"
)

func policyCheck(line string, checker func(string, passwordPolicy) bool) passwordLine {
	explodedPassword := strings.Split(line, ":")
	passwordPolicyPart := strings.TrimSpace(explodedPassword[0])
	password := strings.TrimSpace(explodedPassword[1])

	policy := makePolicy(passwordPolicyPart)

	validity := checker(password, policy)

	return passwordLine{policy, password, validity}
}

func makePolicy(policyPart string) passwordPolicy {
	explodedPasswordPolicy := strings.Split(policyPart, " ")
	policyCharacter := strings.TrimSpace(explodedPasswordPolicy[1])
	policyRepeats := strings.TrimSpace(explodedPasswordPolicy[0])

	repitionCounts := strings.Split(policyRepeats, "-")
	lowCount, err := strconv.Atoi(repitionCounts[0])
	check(err)
	highCount, err := strconv.Atoi(repitionCounts[1])
	check(err)

	return passwordPolicy{lowCount, highCount, policyCharacter}
}
