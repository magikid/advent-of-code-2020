package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGroupCheck(t *testing.T) {
	groupAnswers := "abcx abcy abcz"
	customForm := makeCustomsForm(groupAnswers)

	assert.Equal(t, 6, len(customForm.AnyoneAnsweredYes()))
}

func TestExampleInput(t *testing.T) {
	rawPlaneAnswers := []string{
		"abc",
		"",
		"a",
		"b",
		"c",
		"",
		"ab",
		"ac",
		"",
		"a",
		"a",
		"a",
		"a",
		"",
		"b",
	}

	records := make(chan string)
	go fixInput(rawPlaneAnswers, records)

	group1 := makeCustomsForm(<-records)
	assert.Equal(t, []rune{'a', 'b', 'c'}, group1.AnyoneAnsweredYes())

	group2 := makeCustomsForm(<-records)
	assert.Equal(t, []rune{'a', 'b', 'c'}, group2.AnyoneAnsweredYes())

	group3 := makeCustomsForm(<-records)
	assert.Equal(t, []rune{'a', 'b', 'c'}, group3.AnyoneAnsweredYes())

	group4 := makeCustomsForm(<-records)
	assert.Equal(t, []rune{'a'}, group4.AnyoneAnsweredYes())

	group5 := makeCustomsForm(<-records)
	assert.Equal(t, []rune{'b'}, group5.AnyoneAnsweredYes())
}

func TestExampleInputPart2(t *testing.T) {
	rawPlaneAnswers := []string{
		"abc",
		"",
		"a",
		"b",
		"c",
		"",
		"ab",
		"ac",
		"",
		"a",
		"a",
		"a",
		"a",
		"",
		"b",
	}

	records := make(chan string)
	go fixInput(rawPlaneAnswers, records)

	group1 := makeCustomsForm(<-records)
	assert.Equal(t, []rune{'a', 'b', 'c'}, group1.EveryoneAnsweredYes())

	group2 := makeCustomsForm(<-records)
	assert.Equal(t, []rune{}, group2.EveryoneAnsweredYes())

	group3 := makeCustomsForm(<-records)
	assert.Equal(t, []rune{'a'}, group3.EveryoneAnsweredYes())

	group4 := makeCustomsForm(<-records)
	assert.Equal(t, []rune{'a'}, group4.EveryoneAnsweredYes())

	group5 := makeCustomsForm(<-records)
	assert.Equal(t, []rune{'b'}, group5.EveryoneAnsweredYes())
}

func TestMakeForm(t *testing.T) {
	groupAnswers := "abcx abcy abcz"

	customForm := makeCustomsForm(groupAnswers)
	assert.Equal(t, []rune{'a', 'b', 'c', 'x', 'y', 'z'}, customForm.AnyoneAnsweredYes())
}
