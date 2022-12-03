package rfc3339date_test

import (
	"testing"
	"time"

	"github.com/boundedinfinity/rfc3339date"
	"github.com/stretchr/testify/assert"
)

const (
	TEST_DATETIME = `2020-06-13T07:08:09Z`
	TEST_DATE     = `2020-06-13`
	TEST_TIME     = `07:08:09`
)

func testTime() rfc3339date.Rfc3339Time {
	d, err := time.Parse(time.RFC3339, TEST_DATETIME)

	if err != nil {
		panic(err)
	}

	return rfc3339date.NewTime(d)
}

func testDate() rfc3339date.Rfc3339Date {
	d, err := time.Parse(time.RFC3339, TEST_DATETIME)

	if err != nil {
		panic(err)
	}

	return rfc3339date.NewDate(d)
}

func testDateTime() rfc3339date.Rfc3339DateTime {
	d, err := time.Parse(time.RFC3339, TEST_DATETIME)

	if err != nil {
		panic(err)
	}

	return rfc3339date.NewDateTime(d)
}

func compareDate(t *testing.T, a, e rfc3339date.Rfc3339Date) {
	assert.Equal(t, e.Year(), a.Year())
	assert.Equal(t, e.Month(), a.Month())
	assert.Equal(t, e.Day(), a.Day())
}

func compareDateTime(t *testing.T, a, e rfc3339date.Rfc3339DateTime) {
	assert.Equal(t, e.Year(), a.Year())
	assert.Equal(t, e.Month(), a.Month())
	assert.Equal(t, e.Day(), a.Day())

	assert.Equal(t, e.Hour(), a.Hour())
	assert.Equal(t, e.Minute(), a.Minute())
	assert.Equal(t, e.Second(), a.Second())
}

func compareTime(t *testing.T, a, e rfc3339date.Rfc3339Time) {
	assert.Equal(t, e.Hour(), a.Hour())
	assert.Equal(t, e.Minute(), a.Minute())
	assert.Equal(t, e.Second(), a.Second())
}
