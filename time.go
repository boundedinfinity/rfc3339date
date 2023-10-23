package rfc3339date

import (
	"fmt"
	"strings"
	"time"

	"github.com/boundedinfinity/rfc3339date/internal"
)

type Rfc3339Time struct {
	time.Time
}

func (t Rfc3339Time) String() string {
	return strings.TrimSuffix(t.Format(internal.FORMAT_TIME), "Z")
}

func (t Rfc3339Time) After(v Rfc3339Time) bool {
	return t.Time.After(v.Time)
}

func (t Rfc3339Time) Before(v Rfc3339Time) bool {
	return t.Time.Before(v.Time)
}

func NewTime(d time.Time) Rfc3339Time {
	return Rfc3339Time{d}
}

func ParseTime(s string) (Rfc3339Time, error) {
	s2 := fmt.Sprintf("%v%v%v", internal.FAKE_DATE1, s, internal.FAKE_DATE2)
	d, err := time.Parse(internal.FORMAT_DATETIME, s2)

	if err != nil {
		return Rfc3339Time{}, err
	}

	return Rfc3339Time{d}, nil
}
