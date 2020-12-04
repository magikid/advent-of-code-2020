package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

type year struct {
	year time.Time
}

func buildYear(inputYear string) *year {
	if inputYear == "" {
		return nil
	}

	date, err := time.Parse("2006", inputYear)
	check(err)

	return &year{date}
}

func (y *year) Year() int {
	return y.year.Year()
}

type passport struct {
	birthYear      *year
	issueYear      *year
	expirationYear *year
	height         int
	hairColor      string
	eyeColor       string
	passpordID     int
	countryID      int
}

func (y *year) String() string {
	return fmt.Sprint(y.Year())
}

func (p *passport) MostlyValid() bool {
	return p.birthYear != nil &&
		p.issueYear != nil &&
		p.expirationYear != nil &&
		p.height != 0 &&
		p.hairColor != "" &&
		p.eyeColor != "" &&
		p.passpordID != 0
}

func (p *passport) String() string {
	return fmt.Sprintf("passport: byr %v iyr %v eyr %v hgt %v hcl %v ecl %v pid %v", p.birthYear, p.issueYear, p.expirationYear, p.height, p.hairColor, p.eyeColor, p.passpordID)
}

func buildPassport(record string) *passport {
	var err error
	var newPassport passport
	fields := strings.Fields(record)
	for _, field := range fields {
		var fieldName, fieldValue string
		fmt.Sscanf(field, "%3s:%s", &fieldName, &fieldValue)
		switch fieldName {
		case "byr":
			newPassport.birthYear = buildYear(fieldValue)
		case "iyr":
			newPassport.issueYear = buildYear(fieldValue)
		case "eyr":
			newPassport.expirationYear = buildYear(fieldValue)
		case "hgt":
			fieldValue = strings.TrimSuffix(fieldValue, "cm")
			fieldValue = strings.TrimSuffix(fieldValue, "in")
			newPassport.height, err = strconv.Atoi(fieldValue)
			check(err)
		case "hcl":
			newPassport.hairColor = fieldValue
		case "ecl":
			newPassport.eyeColor = fieldValue
		case "pid":
			newPassport.passpordID, err = strconv.Atoi(fieldValue)
			check(err)
		case "cid":
			newPassport.countryID, err = strconv.Atoi(fieldValue)
			check(err)
		default:
			log.Fatal("Invalid field in record")
		}
	}

	return &newPassport
}
