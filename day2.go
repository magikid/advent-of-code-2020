package main

import (
	"log"
	"strings"
)

// Day2Solution1 finds the number of valid passwords in the list
func Day2Solution1(puzzleInputs []string, completed chan bool) {
	validPasswords := 0

	for _, rawLine := range puzzleInputs {
		if policyCheck(rawLine, day2Part1PolicyChecker).valid {
			validPasswords++
		}
	}

	log.Printf("Number of valid passwords: %v", validPasswords)
	completed <- true
}

func day2Part1PolicyChecker(password string, policy passwordPolicy) bool {
	matchCounts := strings.Count(password, string(policy.letter))

	return policy.lowCount <= matchCounts && matchCounts <= policy.highCount
}

// Day2Solution2 finds the number of valid passwords with the new set of rules
func Day2Solution2(puzzleInputs []string, completed chan bool) {
	validPasswords := 0

	for _, rawLine := range puzzleInputs {
		if policyCheck(rawLine, day2Part2PolicyChecker).valid {
			validPasswords++
		}
	}

	log.Printf("Number of valid passwords: %v", validPasswords)
	completed <- true
}

func day2Part2PolicyChecker(password string, policy passwordPolicy) bool {
	firstLetter := rune(password[policy.lowCount-1])
	secondLetter := rune(password[policy.highCount-1])

	return (firstLetter == policy.letter) != (secondLetter == policy.letter)
}
