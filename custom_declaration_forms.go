package main

import (
	"fmt"
	"sort"
	"strings"
)

type customsForm struct {
	anyYes        []rune
	allYes        []rune
	peopleAnswers []string
}

func (form *customsForm) String() string {
	return fmt.Sprintf("form: peoples' answers: %v, anyYes: %v, allYes: %v", form.peopleAnswers, form.anyYes, form.allYes)
}

func (form *customsForm) AnyoneAnsweredYes() []rune {
	runes := make([]rune, len(form.anyYes))

	for i, value := range form.anyYes {
		runes[i] = rune(value)
	}
	sort.SliceStable(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return runes
}

func (form *customsForm) EveryoneAnsweredYes() []rune {
	runes := make([]rune, len(form.allYes))

	for i, value := range form.allYes {
		runes[i] = rune(value)
	}
	sort.SliceStable(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})

	return runes
}

func makeCustomsForm(input string) customsForm {
	anyFoundTracker := make(map[rune]bool)
	allFoundTracker := make(map[rune]int)
	var anyYes, allYes []rune
	newForm := customsForm{anyYes: anyYes, allYes: allYes}
	numberOfPeopleInGroup := len(strings.Fields(input))

	for _, peoplesAnswers := range strings.Fields(input) {
		for _, letter := range peoplesAnswers {
			anyFoundTracker[letter] = true
			allFoundTracker[letter]++
		}
		newForm.peopleAnswers = append(newForm.peopleAnswers, peoplesAnswers)
	}

	for key, value := range allFoundTracker {
		if value == numberOfPeopleInGroup {
			newForm.allYes = append(newForm.allYes, rune(key))
		}
	}

	for key := range anyFoundTracker {
		newForm.anyYes = append(newForm.anyYes, key)
	}

	return newForm
}

func getForms(records chan string, forms chan customsForm) {
	for record := range records {
		forms <- makeCustomsForm(record)
	}
	close(forms)
}
