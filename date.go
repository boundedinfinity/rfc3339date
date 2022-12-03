package rfc3339date

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"time"

	"github.com/boundedinfinity/rfc3339date/internal"
	"gopkg.in/yaml.v3"
)

type Rfc3339Date struct {
	time.Time
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

func (t Rfc3339Date) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(t.String(), start)
}

func (t *Rfc3339Date) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var s string

	if err := d.DecodeElement(&s, &start); err != nil {
		return err
	}

	v, err := ParseDate(s)

	if err != nil {
		return err
	}

	*t = v

	return nil
}

func (t Rfc3339Date) MarshalYAML() (interface{}, error) {
	return t.String(), nil
}

func (t *Rfc3339Date) UnmarshalYAML(value *yaml.Node) error {
	var s string

	if err := value.Decode(&s); err != nil {
		return err
	}

	v, err := ParseDate(s)

	if err != nil {
		return err
	}

	*t = v

	return nil
}
