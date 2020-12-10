package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
)

type xmasCipher struct {
	rawInput       []int
	preambleLength int
	preamble       []int
	contents       map[int]bool
	weaknessTarget int
}

func (cipher xmasCipher) String() string {
	return fmt.Sprintf("target: %v, preambleLength: %v, preamble: %v, contents: %v, input: %v", cipher.weaknessTarget, cipher.preambleLength, cipher.preamble, cipher.contents, cipher.rawInput)
}

func (cipher *xmasCipher) UpdatePreamble(nextNumber int) {
	tempCipher := make([]int, cipher.preambleLength)

	tempCipher[cipher.preambleLength-1] = nextNumber
	for i := 0; i < cipher.preambleLength-1; i++ {
		tempCipher[i] = cipher.preamble[i+1]
	}

	cipher.preamble = tempCipher
}

func (cipher *xmasCipher) Add(nextNumber int) {
	cipher.contents = make(map[int]bool)

	for i := 0; i < cipher.preambleLength; i++ {
		tempSum := nextNumber - cipher.preamble[i]
		_, ok := cipher.contents[tempSum]
		if ok {
			cipher.contents[cipher.preamble[i]] = true
			cipher.UpdatePreamble(nextNumber)
			return
		}
		cipher.contents[cipher.preamble[i]] = true
	}

	cipher.weaknessTarget = nextNumber
}

func (cipher *xmasCipher) Broken() bool {
	if cipher.weaknessTarget != 0 {
		return true
	}

	return false
}

func (cipher *xmasCipher) FindWeakness() int {
	if cipher.weaknessTarget == 0 {
		return 0
	}

	targetIndex := 0
	for i, val := range cipher.rawInput {
		if val == cipher.weaknessTarget {
			targetIndex = i
			break
		}
	}

	for i := 0; i <= targetIndex; i++ {
		for j := i; j < targetIndex; j++ {
			testRange := cipher.rawInput[i:j]
			if sum(testRange) == cipher.weaknessTarget {
				sort.Ints(testRange)
				return testRange[0] + testRange[len(testRange)-1]
			}
		}
	}

	return 0
}

func sum(numbers []int) int {
	result := 0
	for _, number := range numbers {
		result += number
	}

	return result
}

func makeXmasCipherBase(input []string, preambleLength int) xmasCipher {
	contents := make(map[int]bool)
	rawInput := make([]int, len(input))
	cipher := xmasCipher{contents: contents, preambleLength: preambleLength}

	for i, value := range input {
		rawInput[i] = makeNumber(value)
	}
	cipher.rawInput = rawInput

	preamble := make([]int, cipher.preambleLength)
	cipher.preamble = preamble

	if len(cipher.rawInput) < cipher.preambleLength {
		log.Printf("Preable too small! %v", cipher.rawInput)
		return cipher
	}

	for i := 0; i < cipher.preambleLength; i++ {
		cipher.UpdatePreamble(cipher.rawInput[i])
	}

	for i := cipher.preambleLength; i < len(cipher.rawInput); i++ {
		cipher.Add(cipher.rawInput[i])
		if cipher.Broken() {
			return cipher
		}
	}

	return cipher
}

func makeXmasCipher(input []string) xmasCipher {
	return makeXmasCipherBase(input, 25)
}

func makeNumber(input string) int {
	number, err := strconv.Atoi(input)
	check(err)

	return number
}
