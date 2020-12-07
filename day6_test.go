package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGroupCheck(t *testing.T) {
	groupAnswers := []string{
		"abcx",
		"abcy",
		"abcz",
	}
	customForm := makeCustomsForm(groupAnswers)

	assert.Equal(t, 6, customForm.CountYes())
}
