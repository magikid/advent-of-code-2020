package main

import (
	"fmt"
	"strconv"
)

const cipherLength = 25

type xmasCipher struct {
	preamble       [cipherLength]int
	contents       map[int]bool
	weaknessTarget int
}

func (cipher xmasCipher) String() string {
	return fmt.Sprintf("target: %v, preamble: %v", cipher.weaknessTarget, cipher.preamble)
}

func (cipher *xmasCipher) UpdatePreamble(nextNumber int) {
	var tempCipher [25]int

	tempCipher[cipherLength-1] = nextNumber
	for i := 0; i < cipherLength-1; i++ {
		tempCipher[i] = cipher.preamble[i+1]
	}

	cipher.preamble = tempCipher
}

func (cipher *xmasCipher) Add(nextNumber int) {

	for i := 0; i < cipherLength; i++ {
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

func makeXmasCipher(input []string) (xmasCipher, error) {
	var err error
	var preamble [cipherLength]int
	contents := make(map[int]bool)
	cipher := xmasCipher{contents: contents}

	if len(input) < cipherLength {
		err = fmt.Errorf("Preable too small! %v", input)
		return cipher, err
	}

	for i := 0; i < cipherLength; i++ {
		preamble[i] = makeNumber(input[i])
	}
	cipher.preamble = preamble

	for i := 25; i < len(input); i++ {
		cipher.Add(makeNumber(input[i]))
		if cipher.Broken() {
			return cipher, err
		}
	}

	return cipher, err
}

func makeNumber(input string) int {
	number, err := strconv.Atoi(input)
	check(err)

	return number
}
