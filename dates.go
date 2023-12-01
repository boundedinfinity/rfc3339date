package rfc3339date

import (
	"fmt"
	"time"

	"github.com/boundedinfinity/rfc3339date/internal"
)

var Dates = dates{}

func init() {
	var zero time.Time
	Dates.Zero = NewDate(zero)
}

type dates struct {
	Zero Rfc3339Date
}

func (t dates) Native(d time.Time) Rfc3339Date {
	return NewDate(d)
}

func (t dates) Now() Rfc3339Date {
	return NewDate(time.Now())
}

func (t dates) IsZero(d Rfc3339Date) bool {
	return d == Dates.Zero
}

func (t dates) Parse(s string) (Rfc3339Date, error) {
	s2 := fmt.Sprintf("%v%v", s, internal.FAKE_TIME)
	d, err := time.Parse(internal.FORMAT_DATETIME, s2)

	if err != nil {
		return Rfc3339Date{}, err
	}

	return NewDate(d), nil
}
