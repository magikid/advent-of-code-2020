package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildYear(t *testing.T) {
	testYear := buildYear("1929")
	testYear2 := buildYear("2020")

	assert.Equal(t, 1929, testYear.Year())
	assert.Equal(t, 2020, testYear2.Year())
}

func TestMakePassport(t *testing.T) {
	testRecord1 := buildPassport(`ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
	byr:1937 iyr:2017 cid:147 hgt:183cm`)
	assert.True(t, testRecord1.MostlyValid())

	testRecord2 := buildPassport(`iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
	hcl:#cfa07d byr:1929`)
	assert.False(t, testRecord2.MostlyValid())

	testRecord3 := buildPassport(`hcl:#ae17e1 iyr:2013
	eyr:2024
	ecl:brn pid:760753108 byr:1931
	hgt:179cm`)
	assert.True(t, testRecord3.MostlyValid())

	testRecord4 := buildPassport(`hcl:#cfa07d eyr:2025 pid:166559648
	iyr:2011 ecl:brn hgt:59in`)
	assert.False(t, testRecord4.MostlyValid())

	testRecord5 := buildPassport(`pid:#1bb4d8`)
	assert.Equal(t, "#1bb4d8", testRecord5.passpordID)

	testRecord6 := buildPassport(`pid:937877382`)
	assert.Equal(t, "937877382", testRecord6.passpordID)
}

func BenchmarkDay4(b *testing.B) {
	b.Run("passport builder", benchmarkBuildPassport)
}

func benchmarkBuildPassport(b *testing.B) {
	for i := 0; i < b.N; i++ {
		buildPassport(`ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
	byr:1937 iyr:2017 cid:147 hgt:183cm`)
	}
}

func TestValidBirthYear(t *testing.T) {
	year1 := buildYear("1919")
	assert.False(t, year1.ValidBirthYear())
	year2 := buildYear("1920")
	assert.True(t, year2.ValidBirthYear())
	year3 := buildYear("1921")
	assert.True(t, year3.ValidBirthYear())

	year4 := buildYear("2001")
	assert.True(t, year4.ValidBirthYear())
	year5 := buildYear("2002")
	assert.True(t, year5.ValidBirthYear())
	year6 := buildYear("2003")
	assert.False(t, year6.ValidBirthYear())
}

func TestValidIssueYear(t *testing.T) {
	year1 := buildYear("2009")
	assert.False(t, year1.ValidIssueYear())
	year2 := buildYear("2010")
	assert.True(t, year2.ValidIssueYear())
	year3 := buildYear("2011")
	assert.True(t, year3.ValidIssueYear())
	year4 := buildYear("2019")
	assert.True(t, year4.ValidIssueYear())
	year5 := buildYear("2020")
	assert.True(t, year5.ValidIssueYear())
	year6 := buildYear("2021")
	assert.False(t, year6.ValidIssueYear())
}

func TestValidExpirationYear(t *testing.T) {
	year1 := buildYear("2019")
	assert.False(t, year1.ValidExpirationYear())
	year2 := buildYear("2020")
	assert.True(t, year2.ValidExpirationYear())
	year3 := buildYear("2021")
	assert.True(t, year3.ValidExpirationYear())
	year4 := buildYear("2029")
	assert.True(t, year4.ValidExpirationYear())
	year5 := buildYear("2030")
	assert.True(t, year5.ValidExpirationYear())
	year6 := buildYear("2031")
	assert.False(t, year6.ValidExpirationYear())
}

func TestValidHeight(t *testing.T) {
	height1 := "59in"
	passport1 := buildPassport(fmt.Sprintf("hgt:%v", height1))
	assert.True(t, passport1.ValidHeight())
	height2 := "77in"
	passport2 := buildPassport(fmt.Sprintf("hgt:%v", height2))
	assert.False(t, passport2.ValidHeight())
}

func TestValidHairColor(t *testing.T) {
	hairColor1 := "#101010"
	passport1 := buildPassport(fmt.Sprintf("hcl:%v", hairColor1))
	assert.True(t, passport1.ValidHairColor())

	hairColor2 := "1111gg"
	passport2 := buildPassport(fmt.Sprintf("hcl:%v", hairColor2))
	assert.False(t, passport2.ValidHairColor())

	hairColor3 := "#000000"
	passport3 := buildPassport(fmt.Sprintf("hcl:%v", hairColor3))
	assert.True(t, passport3.ValidHairColor())
}

func TestValidPassportId(t *testing.T) {
	pid1 := "000111222"
	passport1 := buildPassport(fmt.Sprintf("pid:%v", pid1))
	assert.True(t, passport1.ValidPassportID())

	pid2 := "123456"
	passport2 := buildPassport(fmt.Sprintf("pid:%v", pid2))
	assert.False(t, passport2.ValidPassportID())

	pid3 := "1234567890"
	passport3 := buildPassport(fmt.Sprintf("pid:%v", pid3))
	assert.False(t, passport3.ValidPassportID())
}
