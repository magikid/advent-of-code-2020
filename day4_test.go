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
	height3 := "60in"
	passport3 := buildPassport(fmt.Sprintf("hgt:%v", height3))
	assert.True(t, passport3.ValidHeight())
	height4 := "190cm"
	passport4 := buildPassport(fmt.Sprintf("hgt:%v", height4))
	assert.True(t, passport4.ValidHeight())
	height5 := "190in"
	passport5 := buildPassport(fmt.Sprintf("hgt:%v", height5))
	assert.False(t, passport5.ValidHeight())
	height6 := "190"
	passport6 := buildPassport(fmt.Sprintf("hgt:%v", height6))
	assert.False(t, passport6.ValidHeight())
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

	hairColor4 := "#123abc"
	passport4 := buildPassport(fmt.Sprintf("hcl:%v", hairColor4))
	assert.True(t, passport4.ValidHairColor())

	hairColor5 := "#123abz"
	passport5 := buildPassport(fmt.Sprintf("hcl:%v", hairColor5))
	assert.False(t, passport5.ValidHairColor())

	hairColor6 := "123abc"
	passport6 := buildPassport(fmt.Sprintf("hcl:%v", hairColor6))
	assert.False(t, passport6.ValidHairColor())

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

	pid4 := "000000001"
	passport4 := buildPassport(fmt.Sprintf("pid:%v", pid4))
	assert.True(t, passport4.ValidPassportID())

	pid5 := "0123456789"
	passport5 := buildPassport(fmt.Sprintf("pid:%v", pid5))
	assert.False(t, passport5.ValidPassportID())
}

func TestFullyValidPassports(t *testing.T) {
	passport1 := buildPassport("eyr:1972 cid:100 hcl:#18171d ecl:amb hgt:170 pid:186cm iyr:2018 byr:1926")
	assert.False(t, passport1.FullyValid())
	passport2 := buildPassport("iyr:2019 hcl:#602927 eyr:1967 hgt:170cm ecl:grn pid:012533040 byr:1946")
	assert.False(t, passport2.FullyValid())
	passport3 := buildPassport("hcl:dab227 iyr:2012 ecl:brn hgt:182cm pid:021572410 eyr:2020 byr:1992 cid:277")
	assert.False(t, passport3.FullyValid())
	passport4 := buildPassport("hgt:59cm ecl:zzz eyr:2038 hcl:74454a iyr:2023 pid:3556412378 byr:2007")
	assert.False(t, passport4.FullyValid())

	passport5 := buildPassport("pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980 hcl:#623a2f")
	assert.True(t, passport5.FullyValid())
	passport6 := buildPassport("eyr:2029 ecl:blu cid:129 byr:1989 iyr:2014 pid:896056539 hcl:#a97842 hgt:165cm")
	assert.True(t, passport6.FullyValid())
	passport7 := buildPassport("hcl:#888785 hgt:164cm byr:2001 iyr:2015 cid:88 pid:545766238 ecl:hzl eyr:2022")
	assert.True(t, passport7.FullyValid())
	passport8 := buildPassport("iyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719")
	assert.True(t, passport8.FullyValid())
}

func TestValidEyeColor(t *testing.T) {
	eyeColor1 := "amb"
	passport1 := buildPassport(fmt.Sprintf("ecl:%v", eyeColor1))
	assert.True(t, passport1.ValidEyeColor())

	eyeColor2 := "foo"
	passport2 := buildPassport(fmt.Sprintf("ecl:%v", eyeColor2))
	assert.False(t, passport2.ValidEyeColor())
}
