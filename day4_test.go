package main

import (
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
