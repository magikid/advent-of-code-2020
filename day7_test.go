package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBagsContaining(t *testing.T) {
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

	formattedRules := findBagsContaining(rawRules, "shiny gold")
	assert.Equal(t, 4, len(formattedRules))
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
	assert.Equal(t, 4, len(formattedRules))
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
