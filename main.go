package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func check(e error) {
	if e != nil {
		log.Fatalf("error: %v", e)
	}
}

func getUserSelection() userSelections {
	dayPtr := flag.Int("day", 0, "Which day's solution to check")
	partPtr := flag.Int("part", 0, "Which part of the day's solution to run")

	flag.Parse()
	return userSelections{*dayPtr, part(*partPtr)}
}

func runPart(day int, part part, puzzleInput []string) {
	log.SetPrefix(fmt.Sprintf("day%v ", day))
	functionMapping := map[string]func([]string, chan string){
		"day1part1": Day1Solution1,
		"day1part2": Day1Solution2,
		"day2part1": Day2Solution1,
		"day2part2": Day2Solution2,
	}
	var results = make(chan string)

	if part == allParts {
		go functionMapping[fmt.Sprintf("day%vpart1", day)](puzzleInput, results)
		go functionMapping[fmt.Sprintf("day%vpart2", day)](puzzleInput, results)
		log.Print(<-results)
		log.Print(<-results)

		return
	}

	go functionMapping[fmt.Sprintf("day%vpart%v", day, part)](puzzleInput, results)
	log.Print(<-results)
}

func main() {
	var daysToRun []int
	selection := getUserSelection()

	if selection.day == 0 {
		daysToRun = []int{1, 2}
	} else {
		daysToRun = []int{selection.day}
	}

	for _, day := range daysToRun {
		inputFileName := fmt.Sprintf("inputs/day%v_input.txt", day)
		file, err := os.Open(inputFileName)
		check(err)
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)
		var puzzleInput []string

		for scanner.Scan() {
			puzzleInput = append(puzzleInput, scanner.Text())
		}

		file.Close()

		runPart(day, selection.part, puzzleInput)
	}
}
