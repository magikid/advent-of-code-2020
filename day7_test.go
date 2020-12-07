package main

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeRule(t *testing.T) {
	testRule := makeRule("light red bags contain 1 bright white bag, 2 muted yellow bags.")
	expectedSubRules := []subRule{{color: "bright white", bagsInside: 1}, {color: "muted yellow", bagsInside: 2}}

	assert.Equal(t, expectedSubRules, testRule.subRules)
	assert.Equal(t, "light red", testRule.color)

	testRule2 := makeRule("bright white bags contain 1 shiny gold bag.")
	expectedSubRule := []subRule{{color: "shiny gold", bagsInside: 1}}
	assert.Equal(t, expectedSubRule, testRule2.subRules)
	assert.Equal(t, "bright white", testRule2.color)
}

func TestMakeManyRules(t *testing.T) {
	rawRules := []string{
		"light red bags contain 1 bright white bag, 2 muted yellow bags.",
		"dark orange bags contain 3 bright white bags, 4 muted yellow bags.",
		"bright white bags contain 1 shiny gold bag.",
		"muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.",
		"shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.",
		"dark olive bags contain 3 faded blue bags, 4 dotted black bags.",
		"vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.",
		"faded blue bags contain no other bags.",
		"dotted black bags contain no other bags.",
	}

	formattedRules := findAllRules(rawRules)
	assert.Equal(t, 9, len(formattedRules))
	assert.Equal(t, []subRule{{"none", 0}}, formattedRules["faded blue"])
}

func TestMakeManyRules2(t *testing.T) {
	rawRules := []string{
		"light red bags contain 1 bright white bag, 2 muted yellow bags.",
		"dark orange bags contain 3 bright white bags, 4 muted yellow bags.",
		"bright white bags contain 1 shiny gold bag.",
		"muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.",
		"shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.",
		"dark olive bags contain 3 faded blue bags, 4 dotted black bags.",
		"vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.",
		"faded blue bags contain no other bags.",
		"dotted black bags contain no other bags.",
	}

	formattedRules := findAllRules2(rawRules)
	log.Print(formattedRules)
	assert.Equal(t, 9, formattedRules.Size())
	actualRules, _ := formattedRules.Get("faded blue")
	assert.Equal(t, []subRule{{"none", 0}}, actualRules)
}

func BenchmarkMakeRule(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		makeRule("light red bags contain 1 bright white bag, 2 muted yellow bags.")
		makeRule("faded blue bags contain no other bags.")
	}
}
