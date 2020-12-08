package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type subRule struct {
	color      string
	bagsInside int
}

func (r subRule) String() string {
	return fmt.Sprintf("rule color: %v, bags inside: %v;", r.color, r.bagsInside)
}

func makeSubRule(color string, quantityFromRegexp string) subRule {
	quantity, _ := strconv.Atoi(quantityFromRegexp)
	return subRule{color: color, bagsInside: quantity}
}

type colorRule struct {
	subRules []subRule
	color    string
}

func (c *colorRule) SubColors() []string {
	colors := make([]string, len(c.subRules))
	for i, rule := range c.subRules {
		colors[i] = rule.color
	}

	return colors
}

func (c *colorRule) String() string {
	return fmt.Sprintf("rule color: %v, subRules: %v;", c.color, c.subRules)
}

func (c *colorRule) contains(needle string) bool {
	for _, subRule := range c.subRules {
		if subRule.color == needle {
			return true
		}
	}

	return false
}

func makeRule(ruleString string) *colorRule {
	bags := strings.Split(ruleString, " bags contain ")

	otherBagString := bags[1]

	otherBagStrings := strings.Split(otherBagString, ",")
	otherBags := make([]subRule, len(otherBagStrings))
	var otherBagQuantity int
	var otherBagModifier, otherBagColor string

	for i, bag := range otherBagStrings {
		if bag == "no other bags." {
			otherBags[i] = subRule{color: "none", bagsInside: 0}
			continue
		}

		fmt.Sscanf(bag, "%d %s %s bag", &otherBagQuantity, &otherBagModifier, &otherBagColor)
		if otherBagQuantity == 0 || otherBagModifier == "" || otherBagColor == "" {
			log.Fatalf("Couldn't parse bag line: %v", bag)
		}

		otherBags[i] = subRule{color: otherBagModifier + " " + otherBagColor, bagsInside: otherBagQuantity}
	}

	return &colorRule{color: bags[0], subRules: otherBags}
}

func findAllRules(input []string) map[string][]subRule {
	rules := make(map[string][]subRule)
	for _, row := range input {
		rule := makeRule(row)
		rules[rule.color] = rule.subRules
	}

	return rules
}

func findRules(input []string) (map[string][]subRule, map[string][]string) {
	mainRuleRegexp := regexp.MustCompile(`^([a-z ]+) bags contain ([a-z0-9, ]+)\.$`)
	subRuleRegexp := regexp.MustCompile(`^(\d) ([a-z ]+) bags?$`)

	contains := make(map[string][]subRule)
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
			bag := makeSubRule(innerBagContents[2], innerBagContents[1])
			contains[outerBag] = append(contains[outerBag], bag)
			containedBy[bag.color] = append(containedBy[bag.color], outerBag)
		}
	}

	return contains, containedBy
}

func findRulesContaining(input []string, needle string) []string {
	_, containedBy := findRules(input)
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
