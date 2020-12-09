package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type instruction struct {
	operation string
	argument  int
	visited   bool
}

func (i *instruction) String() string {
	return fmt.Sprintf("%v %v", i.operation, i.argument)
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
	instructions            []*instruction
	accumulator             int
	currentIndex            int
	lastInstructionExecuted *instruction
	halted                  bool
	output                  string
	input                   []string
	debug                   bool
}

func buildCPU(input []string) *cpu {
	console := cpu{}
	console.accumulator = 0
	console.currentIndex = 0
	console.input = input
	console.instructions = buildInstructions(&console.input)
	console.halted = false

	return &console
}

func (console *cpu) Next() {
	console.currentIndex++
}

func (console *cpu) Tick() {
	if console.currentIndex >= len(console.instructions) {
		console.halted = true
		console.output = fmt.Sprintf("OUT Program terminated successfully, Accumulator: %v", console.accumulator)
		return
	}

	currentInstruction := console.instructions[console.currentIndex]

	if currentInstruction.visited {
		console.halted = true
		console.output = fmt.Sprintf("ERR Infinite loop detected! Last Instruction: %v, Accumulator: %v", console.lastInstructionExecuted, console.accumulator)
		return
	}

	console.lastInstructionExecuted = currentInstruction
	switch currentInstruction.operation {
	case "acc":
		console.accumulator += currentInstruction.argument
		currentInstruction.visited = true
	case "jmp":
		console.currentIndex += currentInstruction.argument
		currentInstruction.visited = true
		return
	case "nop":
		currentInstruction.visited = true
	}
	console.Next()
}

func (console *cpu) Boot() string {
	for !console.halted {
		if console.debug && console.currentIndex < len(console.instructions) {
			log.Printf("%v %v", console.currentIndex, console.instructions[console.currentIndex])
		}
		console.Tick()
	}
	return console.output
}

func (console *cpu) Reset() {
	console.accumulator = 0
	console.currentIndex = 0
	console.halted = false
	console.output = ""
	console.lastInstructionExecuted = nil
	console.instructions = buildInstructions(&console.input)
}

func (console *cpu) CorrectErrors() string {
	var nopIndexes, jmpIndexes []int
	for i, inst := range console.instructions {
		switch inst.operation {
		case "nop":
			nopIndexes = append(nopIndexes, i)
		case "jmp":
			jmpIndexes = append(jmpIndexes, i)
		default:
			continue
		}
	}

	for _, nopIndex := range nopIndexes {
		console.Reset()
		oldInstruction := console.instructions[nopIndex]
		newInstruction := instruction{operation: "jmp", argument: oldInstruction.argument, visited: false}
		console.instructions[nopIndex] = &newInstruction
		output := console.Boot()
		if strings.Contains(output, "success") {
			return output
		}
	}

	for _, jmpIndex := range jmpIndexes {
		console.Reset()
		oldInstruction := console.instructions[jmpIndex]
		newInstruction := instruction{operation: "nop", argument: oldInstruction.argument, visited: false}
		console.instructions[jmpIndex] = &newInstruction
		output := console.Boot()
		if strings.Contains(output, "success") {
			return output
		}
	}

	return ""
}
