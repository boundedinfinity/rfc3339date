package rfc3339date

import (
	"fmt"
	"time"

	"github.com/boundedinfinity/rfc3339date/internal"
)

var Times = times{}

type times struct{}

func (t times) Native(d time.Time) Rfc3339Time {
	return NewTime(d)
}

func (t times) Now() Rfc3339Time {
	return NewTime(time.Now())
}

func (t times) Zero() Rfc3339Time {
	var zero time.Time
	return NewTime(zero)
}

func (t times) Parse(s string) (Rfc3339Time, error) {
	s2 := fmt.Sprintf("%v%v%v", internal.FAKE_DATE1, s, internal.FAKE_DATE2)
	d, err := time.Parse(internal.FORMAT_DATETIME, s2)

	if err != nil {
		return Rfc3339Time{}, err
	}

	return Rfc3339Time{d}, nil
}
