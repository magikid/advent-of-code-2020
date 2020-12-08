package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type innerBag struct {
	color      string
	bagsInside int
}

func (r innerBag) String() string {
	return fmt.Sprintf("color: %v, bags inside: %v;", r.color, r.bagsInside)
}

func makeInnerBag(color string, quantityFromRegexp string) innerBag {
	quantity, _ := strconv.Atoi(quantityFromRegexp)
	return innerBag{color: color, bagsInside: quantity}
}

func findBags(input []string) (map[string][]innerBag, map[string][]string) {
	mainRuleRegexp := regexp.MustCompile(`^([a-z ]+) bags contain ([a-z0-9, ]+)\.$`)
	subRuleRegexp := regexp.MustCompile(`^(\d) ([a-z ]+) bags?$`)

	contains := make(map[string][]innerBag)
	containedBy := make(map[string][]string)

	for _, line := range input {
		mainRule := mainRuleRegexp.FindStringSubmatch(line)
		if mainRule == nil {
			log.Panicf("Failed to parse outer bag line: %v", line)
		}
		outerBag := mainRule[1]
		innerBags := strings.Split(mainRule[2], ", ")
		for _, innerBag := range innerBags {
			if innerBag == "no other bags" {
				continue
			}
			innerBagContents := subRuleRegexp.FindStringSubmatch(innerBag)
			if innerBagContents == nil {
				log.Panicf("Failed to parse inner bag: %v", innerBag)
			}
			bag := makeInnerBag(innerBagContents[2], innerBagContents[1])
			contains[outerBag] = append(contains[outerBag], bag)
			containedBy[bag.color] = append(containedBy[bag.color], outerBag)
		}
	}

	return contains, containedBy
}

func findBagsContaining(input []string, needle string) []string {
	_, containedBy := findBags(input)
	canContain := containedBy[needle]
	seen := make(map[string]bool)
	seen[needle] = true
	for len(canContain) > 0 {
		curr := canContain[0]
		canContain = canContain[1:]
		if seen[curr] {
			continue
		}
		seen[curr] = true
		canContain = append(canContain, containedBy[curr]...)
	}

	i := 0
	bagColors := make([]string, len(seen)-1)
	for bagColor := range seen {
		if bagColor == needle {
			continue
		}

		bagColors[i] = bagColor
		i++
	}

	return bagColors
}
