package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay2Part1PolicyChecker(t *testing.T) {
	checker := day2Part1PolicyChecker

	policy := passwordPolicy{1, 3, 'a'}
	firstLine := checker("abcde", policy)
	assert.True(t, firstLine)

	policy = passwordPolicy{1, 3, 'b'}
	secondLine := checker("cdefg", policy)
	assert.False(t, secondLine)

	policy = passwordPolicy{2, 9, 'c'}
	thirdLine := checker("ccccccccc", policy)
	assert.True(t, thirdLine)
}

func TestDay2Part2PolicyChecker(t *testing.T) {
	checker := day2Part2PolicyChecker

	policy := passwordPolicy{1, 3, 'a'}
	firstPassword := checker("abcde", policy)
	assert.True(t, firstPassword)

	policy = passwordPolicy{1, 3, 'b'}
	secondPassword := checker("cdefg", policy)
	assert.False(t, secondPassword)

	policy = passwordPolicy{2, 9, 'c'}
	thirdPassword := checker("ccccccccc", policy)
	assert.False(t, thirdPassword)
}

func BenchmarkDay2(b *testing.B) {
	b.Run("part1", d2p1)
	b.Run("part2", d2p2)
}

func d2p1(b *testing.B) {
	checker := day2Part1PolicyChecker
	policy1 := passwordPolicy{1, 3, 'a'}
	policy2 := passwordPolicy{1, 3, 'b'}
	policy3 := passwordPolicy{2, 9, 'c'}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		checker("abcde", policy1)
		checker("cdefg", policy2)
		checker("ccccccccc", policy3)
	}
}

func d2p2(b *testing.B) {
	checker := day2Part2PolicyChecker
	policy1 := passwordPolicy{1, 3, 'a'}
	policy2 := passwordPolicy{1, 3, 'b'}
	policy3 := passwordPolicy{2, 9, 'c'}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		checker("abcde", policy1)
		checker("cdefg", policy2)
		checker("ccccccccc", policy3)
	}
}
