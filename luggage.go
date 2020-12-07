package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/emirpasic/gods/trees/avltree"
)

type subRule struct {
	color      string
	bagsInside int
}

func (r subRule) String() string {
	return fmt.Sprintf("rule color: %v, bags inside: %v;", r.color, r.bagsInside)
}

type colorRule struct {
	subRules []subRule
	color    string
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

func findAllRules2(input []string) *avltree.Tree {
	rules := avltree.NewWithStringComparator()
	for _, row := range input {
		rule := makeRule(row)
		for _, subRule := range rule.subRules {
			rules.Put(subRule.color, subRule.bagsInside)
		}
		// if value, found := rules.Get(rule.color) {
		// 	if found {

		// 	}
		// }
	}

	return rules
}
