package rfc3339date

import (
	"encoding/json"
	"fmt"
	"time"
)

type Rfc3339Date struct {
	time.Time
}

func (t Rfc3339Date) String() string {
	return t.Format(_FORMAT_DATE)
}

func NewDate(d time.Time) Rfc3339Date {
	d2 := time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, time.UTC)
	return Rfc3339Date{d2}
}

func ParseDate(s string) (Rfc3339Date, error) {
	s2 := fmt.Sprintf("%v%v", s, _FAKE_TIME)
	d, err := time.Parse(_FORMAT_DATETIME, s2)

	if err != nil {
		return Rfc3339Date{}, err
	}

	return NewDate(d), nil
}

func (t Rfc3339Date) MarshalJSON() ([]byte, error) {
	bs, err := json.Marshal(t.String())

	if err != nil {
		return nil, err
	}

	return bs, nil
}

func (t *Rfc3339Date) UnmarshalJSON(data []byte) error {
	var s string

	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	v, err := ParseDate(s)

	if err != nil {
		return err
	}

	*t = v

	return nil
}
