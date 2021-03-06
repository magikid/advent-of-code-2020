package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"magikid/advent-of-code-2020/day3"
	"os"
)

func check(e error) {
	if e != nil {
		log.Fatalf("error: %v", e)
	}
}

func getUserSelection() userSelections {
	dayPtr := flag.Int("day", 0, "Which day's solution to check")

	flag.Parse()
	return userSelections{*dayPtr, allParts}
}

func runPart(day int, part part, puzzleInput []string) {
	log.SetPrefix(fmt.Sprintf("day%v ", day))
	functionMapping := map[string]func([]string, chan string){
		"day1part1": Day1Solution1,
		"day1part2": Day1Solution2,
		"day2part1": Day2Solution1,
		"day2part2": Day2Solution2,
		"day3part1": day3.Day3Solution1,
		"day3part2": day3.Day3Solution2,
		"day4part1": Day4Solution1,
		"day4part2": Day4Solution2,
		"day5part1": Day5Solution1,
		"day5part2": Day5Solution2,
		"day6part1": Day6Solution1,
		"day6part2": Day6Solution2,
		"day7part1": Day7Solution1,
		"day7part2": Day7Solution2,
		"day8part1": Day8Solution1,
		"day8part2": Day8Solution2,
		"day9part1": Day9Solution1,
		"day9part2": Day9Solution2,
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
		daysToRun = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	} else {
		daysToRun = []int{selection.day}
	}

	for _, day := range daysToRun {
		inputFileName := fmt.Sprintf("inputs/day%v_input.txt", day)
		file, _ := os.Open(inputFileName)
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)
		defer file.Close()
		var puzzleInput []string

		for scanner.Scan() {
			puzzleInput = append(puzzleInput, scanner.Text())
		}

		runPart(day, selection.part, puzzleInput)
	}
}
