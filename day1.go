package main

import "log"

// Solution1 finds two numbers in the list that equal 2020 when summed
func Solution1(puzzleInput []int, completed chan bool) {
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

// Solution2 finds three numbers in the list that equal 2020 when summed
func Solution2(puzzleInput []int, completed chan bool) {
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
