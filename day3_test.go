package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeTree(t *testing.T) {
	realTree := makeTree('#')
	assert.Equal(t, tree(true), realTree)

	openSquare := makeTree('.')
	assert.Equal(t, tree(false), openSquare)
}

func TestMakeTreeMap(t *testing.T) {
	input := []string{
		"..#",
		"#..",
		"..."}
	treeMap := makeTreeMap(input)

	assert.Equal(t, 0, treeMap.currentX)
	assert.Equal(t, 0, treeMap.currentY)
	assert.Equal(t, tree(false), treeMap.grid[0][0])
	assert.Equal(t, tree(true), treeMap.grid[0][2])
	assert.Equal(t, tree(true), treeMap.grid[1][0])
}

func BenchmarkDay3(b *testing.B) {
	b.Run("part 1", d3p1)
}

func d3p1(b *testing.B) {
	input := []string{"..#", "#..", "..."}
	for i := 0; i < b.N; i++ {
		makeTreeMap(input)
	}
}
