package main

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildTree(t *testing.T) {
	input := []string{"blue", "red", "green"}
	rules := buildRules(input)
	log.Print(rules)

	assert.Equal(t, 0, len(rules))
}
