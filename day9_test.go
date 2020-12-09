package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestExamplePart1(t *testing.T) {
	exampleInput1 := make([]string, 25)
	for i := 0; i < 25; i++ {
		exampleInput1[i] = fmt.Sprint(i + 1)
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(exampleInput1), func(i, j int) { exampleInput1[i], exampleInput1[j] = exampleInput1[j], exampleInput1[i] })
	cipher, _ := makeXmasCipher(exampleInput1)
	cipher.Add(26)
	assert.False(t, cipher.Broken(), "cipher: %v", cipher)

	cipher, _ = makeXmasCipher(exampleInput1)
	cipher.Add(49)
	assert.False(t, cipher.Broken(), "cipher: %v", cipher)

	cipher, _ = makeXmasCipher(exampleInput1)
	cipher.Add(100)
	assert.True(t, cipher.Broken())
	assert.Equal(t, 100, cipher.weaknessTarget)

	cipher, _ = makeXmasCipher(exampleInput1)
	cipher.Add(50)
	assert.True(t, cipher.Broken())
	assert.Equal(t, 50, cipher.weaknessTarget)
}

func TestMakeXmasCipherTooSmall(t *testing.T) {
	input := []string{
		"123",
		"456",
	}

	_, err := makeXmasCipher(input)
	assert.Contains(t, err.Error(), "too small")
}

func TestExamplePart2(t *testing.T) {
	exampleInput1 := []string{"1", "3", "5", "7", "9"}
	cipher, _ := makeXmasCipher(exampleInput1)
	cipher.Add(10)
}
