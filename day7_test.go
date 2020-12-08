package main

import (
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

func TestRulesContaining(t *testing.T) {
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

	formattedRules := findRulesContaining(rawRules, "shiny gold")
	assert.Equal(t, 4, len(formattedRules.Values()))
}

func TestContains(t *testing.T) {
	testColorRule := colorRule{color: "shiny blue", subRules: []subRule{{color: "muted black", bagsInside: 1}}}
	assert.False(t, testColorRule.contains("shiny black"))

	testColorRule1 := colorRule{color: "dark orange", subRules: []subRule{{color: "faded blue", bagsInside: 1}}}
	assert.True(t, testColorRule1.contains("faded blue"))
}

func BenchmarkMakeRule(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		makeRule("light red bags contain 1 bright white bag, 2 muted yellow bags.")
		makeRule("faded blue bags contain no other bags.")
	}
}
