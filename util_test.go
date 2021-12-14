package rfc3339date_test

import (
	"time"

	"github.com/boundedinfinity/rfc3339date"
	. "github.com/onsi/gomega"
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

func compareDate(a, e rfc3339date.Rfc3339Date) {
	Expect(a.Year()).WithOffset(1).To(Equal(e.Year()))
	Expect(a.Month()).WithOffset(2).To(Equal(e.Month()))
	Expect(a.Day()).WithOffset(3).To(Equal(e.Day()))
}

func compareDateTime(a, e rfc3339date.Rfc3339DateTime) {
	Expect(a.Year()).WithOffset(1).To(Equal(e.Year()))
	Expect(a.Month()).WithOffset(2).To(Equal(e.Month()))
	Expect(a.Day()).WithOffset(3).To(Equal(e.Day()))

	Expect(a.Hour()).WithOffset(4).To(Equal(e.Hour()))
	Expect(a.Minute()).WithOffset(5).To(Equal(e.Minute()))
	Expect(a.Second()).WithOffset(6).To(Equal(e.Second()))
}

func compareTime(a, e rfc3339date.Rfc3339Time) {
	Expect(a.Hour()).WithOffset(4).To(Equal(e.Hour()))
	Expect(a.Minute()).WithOffset(5).To(Equal(e.Minute()))
	Expect(a.Second()).WithOffset(6).To(Equal(e.Second()))
}
