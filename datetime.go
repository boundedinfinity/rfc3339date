package rfc3339date

import (
	"time"

	"github.com/boundedinfinity/rfc3339date/internal"
)

var ZeroDateTime Rfc3339DateTime

func init() {
	var zero time.Time
	ZeroDateTime = NewDateTime(zero)
}

type Rfc3339DateTime struct {
	time.Time
}

func (t Rfc3339DateTime) IsZero() bool {
	return t == ZeroDateTime
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

func NewDateTime(d time.Time) Rfc3339DateTime {
	return Rfc3339DateTime{d}
}

func DateTimeNow() Rfc3339DateTime {
	return NewDateTime(time.Now())
}

func ParseDateTime(s string) (Rfc3339DateTime, error) {
	v, err := time.Parse(internal.FORMAT_DATETIME, s)

	if err != nil {
		return Rfc3339DateTime{}, err
	}

	return Rfc3339DateTime{v}, nil
}
