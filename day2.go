package main

import (
	"fmt"
	"strings"
)

func fastPolicyChecker(lines <-chan string, passwords chan<- passwordLine) {
	for rawLine := range lines {
		passwords <- policyCheck(rawLine, day2Part1PolicyChecker)
	}
	close(passwords)
}

// Day2Solution1 finds the number of valid passwords in the list
func Day2Solution1(puzzleInputs []string, results chan string) {
	validPasswords := 0
	lines := make(chan string, len(puzzleInputs))
	passwords := make(chan passwordLine)

	for _, rawLine := range puzzleInputs {
		lines <- rawLine
	}
	close(lines)

	go fastPolicyChecker(lines, passwords)

	for password := range passwords {
		if password.valid {
			validPasswords++
		}
	}

	results <- fmt.Sprintf("part1: Number of valid passwords: %v", validPasswords)
}

func day2Part1PolicyChecker(password string, policy passwordPolicy) bool {
	matchCounts := strings.Count(password, string(policy.letter))

	return policy.lowCount <= matchCounts && matchCounts <= policy.highCount
}

// Day2Solution2 finds the number of valid passwords with the new set of rules
func Day2Solution2(puzzleInputs []string, results chan string) {
	validPasswords := 0

	for _, rawLine := range puzzleInputs {
		if policyCheck(rawLine, day2Part2PolicyChecker).valid {
			validPasswords++
		}
	}

	results <- fmt.Sprintf("part2: Number of valid passwords: %v", validPasswords)
}

func day2Part2PolicyChecker(password string, policy passwordPolicy) bool {
	firstLetter := rune(password[policy.lowCount-1])
	secondLetter := rune(password[policy.highCount-1])

	return (firstLetter == policy.letter) != (secondLetter == policy.letter)
}
