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
		"...",
	}
	treeMap := buildTreeMap(input)

	assert.Equal(t, 0, treeMap.currentX)
	assert.Equal(t, 0, treeMap.currentY)
	assert.Equal(t, tree(false), treeMap.grid[0][0])
	assert.Equal(t, tree(true), treeMap.grid[0][2])
	assert.Equal(t, tree(true), treeMap.grid[1][0])
}

func TestMoveX(t *testing.T) {
	input := []string{
		"..#",
		"#..",
		"...",
	}
	slope := buildTreeMap(input)

	moveX(slope, 3)
	assert.Equal(t, 0, slope.currentX)

	moveX(slope, 2)
	assert.Equal(t, 2, slope.currentX)

	moveX(slope, 1)
	assert.Equal(t, 0, slope.currentX)

	moveX(slope, 4)
	assert.Equal(t, 1, slope.currentX)
}

func TestMoveY(t *testing.T) {
	input := []string{
		"..#.....",
		"#.......",
		"........",
	}
	slope := buildTreeMap(input)

	moveY(slope, 5)
	assert.Equal(t, 2, slope.currentY)
}

func TestSledDown(t *testing.T) {
	input := []string{
		"..#.....",
		"#.......",
		"....#...",
	}
	slope := buildTreeMap(input)
	sledDown(slope, 2, 1)

	assert.Equal(t, 2, slope.currentX)
	assert.Equal(t, 1, slope.currentY)
}

func TestOnTree(t *testing.T) {
	input := []string{
		"..#",
		"#..",
		"...",
	}
	slope := buildTreeMap(input)
	sledDown(slope, 2, 0)

	assert.True(t, onTree(slope))
}

func TestHitBottom(t *testing.T) {
	input := []string{
		"..#",
		"#..",
		"...",
	}
	slope := buildTreeMap(input)
	sledDown(slope, 0, 2)

	assert.True(t, hitBottom(slope))
}

func TestPuzzleInput(t *testing.T) {
	input := []string{
		"..##.......",
		"#...#...#..",
		".#....#..#.",
		"..#.#...#.#",
		".#...##..#.",
		"..#.##.....",
		".#.#.#....#",
		".#........#",
		"#.##...#...",
		"#...##....#",
		".#..#...#.#",
	}
	slope := buildTreeMap(input)
	treesHit := findTreesOnPath(slope, 3, 1)

	assert.Equal(t, 7, treesHit)
}

func BenchmarkDay3(b *testing.B) {
	b.Run("part 1 building tree map", d3p1buildTreeMap)
	b.Run("part 1 going down slope", d3p1SledDown)
}

func d3p1buildTreeMap(b *testing.B) {
	input := []string{"..#", "#..", "..."}
	for i := 0; i < b.N; i++ {
		buildTreeMap(input)
	}
}

func d3p1SledDown(b *testing.B) {
	input := []string{"...", "..#", "#.."}
	slope := buildTreeMap(input)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		findTreesOnPath(slope, 1, 1)
	}
}
