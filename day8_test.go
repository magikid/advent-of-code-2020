package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func BenchmarkDay8(b *testing.B) {
	b.Run("build instructions", benchBuildInstructions)
	b.Run("build cpu", benchBuildCPU)
}

func benchBuildInstructions(b *testing.B) {
	inputInstructions := []string{
		"nop +0",
		"acc +1",
		"jmp -4",
	}

	for i := 0; i < b.N; i++ {
		buildInstructions(&inputInstructions)
	}
}

func benchBuildCPU(b *testing.B) {
	inputInstructions := []string{
		"nop +0",
		"acc +1",
		"jmp +4",
		"acc +3",
		"jmp -3",
		"acc -99",
		"acc +1",
		"jmp -4",
		"acc +6",
	}

	for i := 0; i < b.N; i++ {
		buildCPU(inputInstructions)
	}
}

func TestBuildInstructions(t *testing.T) {
	inputInstructions := []string{
		"nop +0",
		"acc +1",
		"jmp -4",
	}
	instructions := buildInstructions(&inputInstructions)
	assert.Equal(t, 0, instructions[0].argument)
	assert.Equal(t, "nop", instructions[0].operation)
	assert.Equal(t, 1, instructions[1].argument)
	assert.Equal(t, "acc", instructions[1].operation)
	assert.Equal(t, -4, instructions[2].argument)
	assert.Equal(t, "jmp", instructions[2].operation)
}

func TestBuildConsole(t *testing.T) {
	inputInstructions := []string{
		"nop +0",
	}
	console := buildCPU(inputInstructions)

	assert.Equal(t, 0, console.accumulator)
	assert.Equal(t, 0, console.currentIndex)
}

func TestBootConsole(t *testing.T) {
	inputInstructions := []string{
		"nop +0",
		"acc +1",
		"jmp +4",
		"acc +3",
		"jmp -3",
		"acc -99",
		"acc +1",
		"jmp -4",
		"acc +6",
	}
	console := buildCPU(inputInstructions)
	output := console.Boot()

	assert.Equal(t, "Infinite loop detected! Accumulator: 5", output)
}
