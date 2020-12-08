package main

import (
	"fmt"
	"strconv"
	"strings"
)

type instruction struct {
	operation string
	argument  int
	visited   bool
}

func buildInstructions(input *[]string) []*instruction {
	instructions := make([]*instruction, len(*input))
	for i, line := range *input {
		splitParts := strings.Split(line, " ")

		convertedArgument, err := strconv.Atoi(splitParts[1])
		check(err)

		newInt := instruction{visited: false}
		newInt.operation = splitParts[0]
		newInt.argument = convertedArgument
		instructions[i] = &newInt
	}

	return instructions
}

type cpu struct {
	instructions []*instruction
	accumulator  int
	currentIndex int
	halted       bool
	err          error
}

func buildCPU(input []string) *cpu {
	console := cpu{}
	console.accumulator = 0
	console.currentIndex = 0
	console.instructions = buildInstructions(&input)
	console.halted = false

	return &console
}

func (console *cpu) Next() {
	console.currentIndex++
}

func (console *cpu) Tick() {
	currentInstruction := console.instructions[console.currentIndex]
	if currentInstruction.visited || console.halted {
		console.halted = true
		console.err = fmt.Errorf("Infinite loop detected! Accumulator: %v", console.accumulator)
	}

	switch currentInstruction.operation {
	case "acc":
		console.accumulator += currentInstruction.argument
	case "jmp":
		console.currentIndex += currentInstruction.argument
		return
	case "nop":
	}
	currentInstruction.visited = true
	console.Next()
}

func (console *cpu) Boot() string {
	for !console.halted {
		console.Tick()
	}
	return console.err.Error()
}
