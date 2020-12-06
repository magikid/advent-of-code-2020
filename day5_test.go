package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvert(t *testing.T) {
	seat := makeSeatAssignment("FBFBBFFRLR")
	assert.Equal(t, 44, seat.row)
	assert.Equal(t, 5, seat.column)
	assert.Equal(t, 357, seat.id)

	seat2 := makeSeatAssignment("BFFFBBFRRR")
	assert.Equal(t, 70, seat2.row)
	assert.Equal(t, 7, seat2.column)
	assert.Equal(t, 567, seat2.id)

	seat3 := makeSeatAssignment("FFFBBBFRRR")
	assert.Equal(t, 14, seat3.row)
	assert.Equal(t, 7, seat3.column)
	assert.Equal(t, 119, seat3.id)

	seat4 := makeSeatAssignment("BBFFBBFRLL")
	assert.Equal(t, 102, seat4.row)
	assert.Equal(t, 4, seat4.column)
	assert.Equal(t, 820, seat4.id)
}

func TestBinary(t *testing.T) {
	input := "FF"
	assert.Equal(t, 0, makeBinary(input, 'F', 'B'))
}
