package rfc3339date

import (
	"time"

	"github.com/boundedinfinity/rfc3339date/internal"
)

var DateTimes = dateTimes{}

func init() {
	var zero time.Time
	DateTimes.Zero = NewDateTime(zero)
}

type dateTimes struct {
	Zero Rfc3339DateTime
}

func (t dateTimes) Native(d time.Time) Rfc3339DateTime {
	return NewDateTime(d)
}

func (t dateTimes) Now() Rfc3339DateTime {
	return NewDateTime(time.Now())
}

func (t dateTimes) IsZero(d Rfc3339DateTime) bool {
	return d == DateTimes.Zero
}

func (t dateTimes) Parse(s string) (Rfc3339DateTime, error) {
	d, err := time.Parse(internal.FORMAT_DATETIME, s)

	if err != nil {
		return Rfc3339DateTime{}, err
	}

	return Rfc3339DateTime{d}, nil
}
