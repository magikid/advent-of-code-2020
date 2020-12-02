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

	go Solution1(puzzleInput, doneCheck)
	go Solution2(puzzleInput, doneCheck)
	<-doneCheck
	<-doneCheck
}
