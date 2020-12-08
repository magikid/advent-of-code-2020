package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/emirpasic/gods/sets/treeset"
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

func findRulesContaining(input []string, needle string) *treeset.Set {

	bags := make(map[string]*treeset.Set)

	for _, row := range input {
		rule := makeRule(row)
		for _, subRule := range rule.subRules {
			currentRule, ok := bags[subRule.color]
			if !ok {
				currentRule = treeset.NewWithStringComparator()
			}
			currentRule.Add(rule.color)
			bags[subRule.color] = currentRule

			otherColors, ok := bags[subRule.color]
			if ok {
				for _, color := range otherColors.Values() {
					otherColors, ok := bags[color.(string)]
					if ok {
						for _, newColor := range otherColors.Values() {
							nextColor := bags[subRule.color]
							nextColor.Add(newColor)
						}

					}
				}
			}
		}
	}

	for _, row := range input {
		rule := makeRule(row)
		for _, subRule := range rule.subRules {
			currentRule, ok := bags[subRule.color]
			if !ok {
				currentRule = treeset.NewWithStringComparator()
			}
			currentRule.Add(rule.color)
			bags[subRule.color] = currentRule

			otherColors, ok := bags[subRule.color]
			if ok {
				for _, color := range otherColors.Values() {
					otherColors, ok := bags[color.(string)]
					if ok {
						for _, newColor := range otherColors.Values() {
							nextColor := bags[subRule.color]
							nextColor.Add(newColor)
						}

					}
				}
			}
		}
	}

	return bags[needle]
}
