package rfc3339date

import (
	"fmt"
	"time"

	"github.com/boundedinfinity/rfc3339date/internal"
)

var Dates = dates{}

type dates struct{}

func (t dates) Native(d time.Time) Rfc3339Date {
	return NewDate(d)
}

func (t dates) Now() Rfc3339Date {
	return NewDate(time.Now())
}

func (t dates) Zero() Rfc3339Date {
	var zero time.Time
	return NewDate(zero)
}

func (t dates) Parse(s string) (Rfc3339Date, error) {
	s2 := fmt.Sprintf("%v%v", s, internal.FAKE_TIME)
	d, err := time.Parse(internal.FORMAT_DATETIME, s2)

	if err != nil {
		return Rfc3339Date{}, err
	}

	return NewDate(d), nil
}
