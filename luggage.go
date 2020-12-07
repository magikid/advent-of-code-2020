package main

import (
	"log"
)

type ruleTree struct {
	left     *ruleTree
	maxLeft  int
	color    string
	right    *ruleTree
	maxRight int
}

type rules []*ruleTree

func (r rules) String() {
	for _, rt := range r {
		channel := walker(rt)
		for value := range channel {
			log.Print(value)
		}
	}
}

func walk(t *ruleTree, ch chan string) {
	if t == nil {
		return
	}
	walk(t.left, ch)
	ch <- t.color
	walk(t.right, ch)
}

func walker(t *ruleTree) <-chan string {
	ch := make(chan string)
	go func() {
		walk(t, ch)
		close(ch)
	}()

	return ch
}

func compare(t1, t2 *ruleTree) bool {
	c1, c2 := walker(t1), walker(t2)
	for {
		v1, ok1 := <-c1
		v2, ok2 := <-c2
		if !ok1 || !ok2 {
			return ok1 == ok2
		}
		if v1 != v2 {
			break
		}
	}

	return false
}

func insert(t *ruleTree, v string) *ruleTree {
	if t == nil {
		return &ruleTree{nil, 0, v, nil, 0}
	}
	if v < t.color {
		t.left = insert(t.left, v)
		return t
	}
	t.right = insert(t.right, v)
	return t
}

func buildRules(input []string) rules {
	var rules rules
	var ruleTree *ruleTree

	if len(input) < 1 {
		return rules
	}

	ruleTree = insert(ruleTree, input[0])
	for i := 1; i < len(input); i++ {
		ruleTree = insert(ruleTree, input[i])
		rules = append(rules, ruleTree)
	}

	return rules
}
