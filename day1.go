package main

import (
	"log"
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
func Day1Solution1(dirtyInput []string, completed chan bool) {
	puzzleInput := day1Cleaner(dirtyInput)
	for _, x := range puzzleInput {
		for _, y := range puzzleInput {
			if x+y == 2020 {
				log.Printf("Found part 1! %v * %v = %v; ", x, y, x*y)
				completed <- true
				return
			}
		}
	}
}

// Day1Solution2 finds three numbers in the list that equal 2020 when summed
func Day1Solution2(dirtyInput []string, completed chan bool) {
	puzzleInput := day1Cleaner(dirtyInput)
	for _, x := range puzzleInput {
		for _, y := range puzzleInput {
			for _, z := range puzzleInput {
				if x+y+z == 2020 {
					log.Printf("Found part 2! %v * %v * %v = %v; ", x, y, z, x*y*z)
					completed <- true
					return
				}
			}
		}
	}
}
