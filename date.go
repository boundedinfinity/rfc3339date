package rfc3339date

import (
	"fmt"
	"time"

	"github.com/boundedinfinity/rfc3339date/internal"
)

var ZeroDate Rfc3339Date

func init() {
	var zero time.Time
	ZeroDate = NewDate(zero)
}

type Rfc3339Date struct {
	time.Time
}

func (t Rfc3339Date) IsZero() bool {
	return t == ZeroDate
}

func (t Rfc3339Date) String() string {
	return t.Format(internal.FORMAT_DATE)
}

func (t Rfc3339Date) After(v Rfc3339Date) bool {
	return t.Time.After(v.Time)
}

func (t Rfc3339Date) Before(v Rfc3339Date) bool {
	return t.Time.Before(v.Time)
}

func (t Rfc3339Date) Ahead(v Rfc3339Duration) bool {
	x := time.Now().Add(v.Duration)
	return NewDate(x).Before(t)
}

func (t Rfc3339Date) Begin(v Rfc3339Duration) bool {
	x := time.Now().Add(-v.Duration)
	return NewDate(x).After(t)
}

func NewDate(d time.Time) Rfc3339Date {
	d2 := time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, time.UTC)
	return Rfc3339Date{d2}
}

func ParseDate(s string) (Rfc3339Date, error) {
	s2 := fmt.Sprintf("%v%v", s, internal.FAKE_TIME)
	d, err := time.Parse(internal.FORMAT_DATETIME, s2)

	if err != nil {
		return Rfc3339Date{}, err
	}

	return NewDate(d), nil
}
