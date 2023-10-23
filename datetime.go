package rfc3339date

import (
	"time"

	"github.com/boundedinfinity/rfc3339date/internal"
)

type Rfc3339DateTime struct {
	time.Time
}

func (t Rfc3339DateTime) String() string {
	return t.Format(internal.FORMAT_DATETIME)
}

func (t Rfc3339DateTime) After(v Rfc3339DateTime) bool {
	return t.Time.After(v.Time)
}

func (t Rfc3339DateTime) Before(v Rfc3339DateTime) bool {
	return t.Time.Before(v.Time)
}

func ZeroDateTime() Rfc3339DateTime {
	var zero time.Time
	return NewDateTime(zero)
}

func NewDateTime(d time.Time) Rfc3339DateTime {
	return Rfc3339DateTime{d}
}

func ParseDateTime(s string) (Rfc3339DateTime, error) {
	v, err := time.Parse(internal.FORMAT_DATETIME, s)

	if err != nil {
		return Rfc3339DateTime{}, err
	}

	return Rfc3339DateTime{v}, nil
}
