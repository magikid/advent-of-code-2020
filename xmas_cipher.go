package main

import (
	"fmt"
	"log"
	"strconv"
)

type xmasCipher struct {
	preambleLength int
	preamble       []int
	contents       map[int]bool
	weaknessTarget int
}

func (cipher xmasCipher) String() string {
	return fmt.Sprintf("target: %v, preamble: %v", cipher.weaknessTarget, cipher.preamble)
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

	return 0
}

func makeXmasCipherBase(input []string, preambleLength int) xmasCipher {
	preamble := make([]int, preambleLength)
	contents := make(map[int]bool)
	cipher := xmasCipher{contents: contents, preambleLength: preambleLength, preamble: preamble}

	if len(input) < cipher.preambleLength {
		log.Printf("Preable too small! %v", input)
		return cipher
	}

	for i := 0; i < cipher.preambleLength; i++ {
		cipher.UpdatePreamble(makeNumber(input[i]))
	}

	for i := 25; i < len(input); i++ {
		cipher.Add(makeNumber(input[i]))
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
