package main

import (
	"fmt"
	"sort"
	"strconv"
)

func day1Cleaner(dirtyInput []string) []int {
	var puzzleInput []int

	for _, i := range dirtyInput {
		j, err := strconv.Atoi(i)
		check(err)
		puzzleInput = append(puzzleInput, j)
	}
	sort.Ints(puzzleInput)

	return puzzleInput
}

// Day1Solution1 finds two numbers in the list that equal 2020 when summed
func Day1Solution1(dirtyInput []string, results chan string) {
	puzzleInput := day1Cleaner(dirtyInput)
	for _, x := range puzzleInput {
		for _, y := range puzzleInput {
			if x+y == 2020 {
				results <- fmt.Sprintf("part1: %v * %v = %v; ", x, y, x*y)
				return
			}
		}
	}
}

// Day1Solution2 finds three numbers in the list that equal 2020 when summed
func Day1Solution2(dirtyInput []string, results chan string) {
	puzzleInput := day1Cleaner(dirtyInput)
	for _, x := range puzzleInput {
		for _, y := range puzzleInput {
			for _, z := range puzzleInput {
				if x+y+z == 2020 {
					results <- fmt.Sprintf("part2: %v * %v * %v = %v; ", x, y, z, x*y*z)
					return
				}
			}
		}
	}
}
