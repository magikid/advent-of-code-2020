package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
)

func check(e error) {
	if e != nil {
		log.Fatalf("error: %v", e)
	}
}

func part1(puzzleInput []int, completed chan bool) {
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

func part2(puzzleInput []int, completed chan bool) {
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

func main() {
	file, err := os.Open("inputs/day1_input.txt")
	check(err)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var puzzleInput []int
	doneCheck := make(chan bool)

	for scanner.Scan() {
		number, err := strconv.Atoi(scanner.Text())
		check(err)
		puzzleInput = append(puzzleInput, number)
	}

	file.Close()
	sort.Ints(puzzleInput)

	go part1(puzzleInput, doneCheck)
	go part2(puzzleInput, doneCheck)
	<-doneCheck
	<-doneCheck
}
