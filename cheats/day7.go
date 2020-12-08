package main

// I coulnd't figure out how to do this day's work in a performant manner.
// Once I solved part 1 (taking around 1.5s to do so), I went looking for a
// better solution. I found this in the AoC reddit:
// https://www.reddit.com/r/adventofcode/comments/k8a31f/2020_day_07_solutions/gex0m0p?context=3

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	lines := helperLib.ReadFileLines("input.txt")
	fmt.Println("Step 1:")
	contains, containedBy := bagMaps(lines)
	fmt.Println(step1(containedBy))

	fmt.Println("Step 2:")
	fmt.Println(step2(contains))
}

type bagCount struct {
	color string
	num   int
}

var (
	inputRegex = regexp.MustCompile("^([a-z ]+) bags contain ([a-z0-9, ]+)\\.$")
	bagRegex   = regexp.MustCompile("^(\\d+) ([a-z ]+) bag[s]?$")
)

func bagMaps(lines []string) (map[string][]bagCount, map[string][]string) {
	contains, containedBy := make(map[string][]bagCount), make(map[string][]string)
	for _, l := range lines {
		tokens := inputRegex.FindStringSubmatch(l)
		if tokens == nil {
			log.Fatalf("Failed to parse %q\n", l)
		}
		// container = "light red"
		container := tokens[1]
		// c = "1 bright white bag"
		for _, c := range strings.Split(tokens[2], ", ") {
			if c == "no other bags" {
				continue
			}
			// contents = ["...", "1", "bright white"]
			contents := bagRegex.FindStringSubmatch(c)
			if contents == nil {
				log.Fatalf("Failed to parse %q\n", c)
			}
			// bagCount == subRule
			bag := bagCount{}
			qty, _ := strconv.Atoi(contents[1])
			bag.num = qty
			bag.color = contents[2]
			contains[container] = append(contains[container], bag)
			containedBy[bag.color] = append(containedBy[bag.color], container)
		}
	}
	return contains, containedBy
}

const targetBag = "shiny gold"

func step1(containedBy map[string][]string) int {
	canContain := containedBy[targetBag]
	seen := make(map[string]bool)
	seen[targetBag] = true
	for len(canContain) > 0 {
		curr := canContain[0]
		canContain = canContain[1:]
		if seen[curr] {
			continue
		}
		seen[curr] = true
		canContain = append(canContain, containedBy[curr]...)
	}
	// Subtract one for shiny gold
	return len(seen) - 1
}

func step2(contains map[string][]bagCount) int64 {
	return countContents(targetBag, contains)
}

func countContents(target string, contains map[string][]bagCount) int64 {
	sum := int64(0)
	for _, c := range contains[target] {
		sum += int64(c.num)
		sum += int64(c.num) * countContents(c.color, contains)
	}
	return sum
}
