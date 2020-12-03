package main

import "fmt"

// Day3Solution1 finds the number of tress in a path on the slope right 3, down 1
func Day3Solution1(input []string, results chan string) {
	slope := buildTreeMap(input)
	treesHit := findTreesOnPath(slope, 3, 1)
	results <- fmt.Sprintf("part1: hit %v trees", treesHit)
}

// Day3Solution2 WIP
func Day3Solution2(input []string, results chan string) {
	slope := buildTreeMap(input)
	r1d1 := findTreesOnPath(slope, 1, 1)

	resetSlope(slope)
	r3d1 := findTreesOnPath(slope, 3, 1)

	resetSlope(slope)
	r5d1 := findTreesOnPath(slope, 5, 1)

	resetSlope(slope)
	r7d1 := findTreesOnPath(slope, 7, 1)

	resetSlope(slope)
	r1d2 := findTreesOnPath(slope, 1, 2)

	results <- fmt.Sprintf("part2: tress hit: %v * %v * %v * %v * %v = %v", r1d1, r3d1, r5d1, r7d1, r1d2, r1d1*r3d1*r5d1*r7d1*r1d2)
}
