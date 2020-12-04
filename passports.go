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
	// if y == nil {
	// 	return 0
	// }

	return y.year.Year()
}

type passport struct {
	birthYear      *year
	issueYear      *year
	expirationYear *year
	height         string
	hairColor      string
	eyeColor       string
	passpordID     string
	countryID      int
}

func (y *year) String() string {
	return fmt.Sprint(y.Year())
}

func (y *year) ValidBirthYear() bool {
	return 1920 <= y.Year() && y.Year() <= 2002
}

func (p *passport) ValidBirthYear() bool {
	if p.birthYear == nil {
		return false
	}

	return p.birthYear.ValidBirthYear()
}

func (y *year) ValidIssueYear() bool {
	return 2010 <= y.Year() && y.Year() <= 2020
}

func (p *passport) ValidIssueYear() bool {
	if p.issueYear == nil {
		return false
	}

	return p.issueYear.ValidIssueYear()
}

func (y *year) ValidExpirationYear() bool {
	return 2020 <= y.Year() && y.Year() <= 2030
}

func (p *passport) ValidExpirationYear() bool {
	if p.expirationYear == nil {
		return false
	}

	return p.expirationYear.ValidExpirationYear()
}

func (p *passport) ValidHeight() bool {
	var height int
	var err error
	if strings.HasSuffix(p.height, "in") {
		height, err = strconv.Atoi(strings.TrimSuffix(p.height, "in"))
		check(err)
		return 59 <= height && height <= 76
	}

	if strings.HasSuffix(p.height, "cm") {
		height, err = strconv.Atoi(strings.TrimSuffix(p.height, "cm"))
		check(err)
		return 150 <= height && height <= 193
	}

	return false
}

func (p *passport) ValidHairColor() bool {
	var hex1, hex2, hex3 int
	var err error

	if p.hairColor[0] != '#' {
		return false
	}

	_, err = fmt.Sscanf(p.hairColor, "#%02x%02x%02x", &hex1, &hex2, &hex3)
	check(err)
	return &hex1 != nil && &hex2 != nil && &hex3 != nil
}

func (p *passport) ValidEyeColor() bool {
	switch p.eyeColor {
	case "amb":
	case "blu":
	case "brn":
	case "gry":
	case "grn":
	case "hzl":
	case "oth":
		return true
	default:
		return false
	}
	return false
}

func (p *passport) ValidPassportID() bool {
	if len(p.passpordID) != 9 {
		return false
	}

	var passportId int
	_, err := fmt.Sscanf(p.passpordID, "%09d", &passportId)
	if err != nil {
		return false
	}

	return passportId != 0
}

func (p *passport) MostlyValid() bool {
	return p.birthYear != nil &&
		p.issueYear != nil &&
		p.expirationYear != nil &&
		p.height != "" &&
		p.hairColor != "" &&
		p.eyeColor != "" &&
		p.passpordID != ""
}

func (p *passport) FullyValid() bool {
	return p.MostlyValid() &&
		p.ValidBirthYear() &&
		p.ValidIssueYear() &&
		p.ValidExpirationYear() &&
		p.ValidHeight() &&
		p.ValidHeight() &&
		p.ValidEyeColor() &&
		p.ValidPassportID()
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
			newPassport.height = fieldValue
		case "hcl":
			newPassport.hairColor = fieldValue
		case "ecl":
			newPassport.eyeColor = fieldValue
		case "pid":
			newPassport.passpordID = fieldValue
		case "cid":
			newPassport.countryID, err = strconv.Atoi(fieldValue)
			check(err)
		default:
			log.Fatal("Invalid field in record")
		}
	}

	return &newPassport
}

func fixInput(input []string, records chan<- string) {
	firstLineOfRecord := 0
	lastLineOfRecord := 0
	for i, line := range input {
		if len(line) <= 0 {
			records <- strings.Join(input[firstLineOfRecord:lastLineOfRecord], " ")
			firstLineOfRecord = i
			lastLineOfRecord = i
		}
		lastLineOfRecord++
	}
	records <- strings.Join(input[firstLineOfRecord:lastLineOfRecord], " ")
	close(records)
}

func buildPassports(records <-chan string, passports chan<- *passport) {
	for record := range records {
		passports <- buildPassport(record)
	}
	close(passports)
}
