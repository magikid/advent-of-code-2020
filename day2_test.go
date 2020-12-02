package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay2Part1PolicyChecker(t *testing.T) {
	checker := day2Part1PolicyChecker

	policy := passwordPolicy{1, 3, "a"}
	firstLine := checker("abcde", policy)
	assert.True(t, firstLine)

	policy = passwordPolicy{1, 3, "b"}
	secondLine := checker("cdefg", policy)
	assert.False(t, secondLine)

	policy = passwordPolicy{2, 9, "c"}
	thirdLine := checker("ccccccccc", policy)
	assert.True(t, thirdLine)
}

func TestDay2Part2PolicyChecker(t *testing.T) {
	checker := day2Part2PolicyChecker

	policy := passwordPolicy{1, 3, "a"}
	firstPassword := checker("abcde", policy)
	assert.True(t, firstPassword)

	policy = passwordPolicy{1, 3, "b"}
	secondPassword := checker("cdefg", policy)
	assert.False(t, secondPassword)

	policy = passwordPolicy{2, 9, "c"}
	thirdPassword := checker("ccccccccc", policy)
	assert.False(t, thirdPassword)
}
