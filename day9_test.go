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
	cipher := makeXmasCipher(exampleInput1)
	cipher.Add(26)
	assert.False(t, cipher.Broken(), "cipher: %v", cipher)

	cipher = makeXmasCipher(exampleInput1)
	cipher.Add(49)
	assert.False(t, cipher.Broken(), "cipher: %v", cipher)

	cipher = makeXmasCipher(exampleInput1)
	cipher.Add(100)
	assert.True(t, cipher.Broken())
	assert.Equal(t, 100, cipher.weaknessTarget)

	cipher = makeXmasCipher(exampleInput1)
	cipher.Add(50)
	assert.True(t, cipher.Broken())
	assert.Equal(t, 50, cipher.weaknessTarget)
}

func TestExamplePart2(t *testing.T) {
	exampleInput2 := []string{
		"35",
		"20",
		"15",
		"25",
		"47",
		"40",
		"62",
		"55",
		"65",
		"95",
		"102",
		"117",
		"150",
		"182",
		"127",
		"219",
		"299",
		"277",
		"309",
		"576",
	}
	cipher2 := makeXmasCipherBase(exampleInput2, 5)
	assert.True(t, cipher2.Broken(), "cipher: %v", cipher2)
	assert.Equal(t, 127, cipher2.weaknessTarget)
	assert.Equal(t, 62, cipher2.FindWeakness())

	exampleInput1 := []string{"1", "3", "5", "7", "9", "15"}
	cipher := makeXmasCipherBase(exampleInput1, 5)
	assert.True(t, cipher.Broken())
	assert.Equal(t, 10, cipher.FindWeakness())
}
