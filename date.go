package rfc3339date

import (
	"database/sql/driver"
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

func (t Rfc3339Date) Within(v Rfc3339Duration) bool {
	half := v.Duration / 2
	before := time.Now().Add(half)
	after := time.Now().Add(-half)
	return NewDate(before).After(t) && NewDate(after).After(t)
}

func NewDate(d time.Time) Rfc3339Date {
	year := d.Year()
	month := d.Month()
	day := d.Day()
	d2 := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
	return Rfc3339Date{d2}
}

// /////////////////////////////////////////////////////////////////
//  unmarshal implemenation
// /////////////////////////////////////////////////////////////////

func (t *Rfc3339Date) unmarshal(s string) error {
	v, err := Dates.Parse(s)

	if err != nil {
		return err
	}

	*t = v

	return nil
}

// /////////////////////////////////////////////////////////////////
//  JSON marshal/unmarshal implemenation
// /////////////////////////////////////////////////////////////////

func (t Rfc3339Date) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func (t *Rfc3339Date) UnmarshalJSON(data []byte) error {
	var s string

	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	return t.unmarshal(s)
}

// /////////////////////////////////////////////////////////////////
//  XML marshal/unmarshal implemenation
// /////////////////////////////////////////////////////////////////

func (t Rfc3339Date) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(t.String(), start)
}

func (t *Rfc3339Date) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var s string

	if err := d.DecodeElement(&s, &start); err != nil {
		return err
	}

	return t.unmarshal(s)
}

// /////////////////////////////////////////////////////////////////
//  YAML marshal/unmarshal implemenation
// /////////////////////////////////////////////////////////////////

func (t Rfc3339Date) MarshalYAML() (interface{}, error) {
	return t.String(), nil
}

func (t *Rfc3339Date) UnmarshalYAML(value *yaml.Node) error {
	var s string

	if err := value.Decode(&s); err != nil {
		return err
	}

	return t.unmarshal(s)
}

// /////////////////////////////////////////////////////////////////
//  SQL marshal/unmarshal implementation
// /////////////////////////////////////////////////////////////////

func (t Rfc3339Date) Value() (driver.Value, error) {
	return t.Date, nil
}

func (t *Rfc3339Date) Scan(value interface{}) error {
	dv, err := driver.String.ConvertValue(value)

	if err != nil {
		return err
	}

	s, ok := dv.(string)

	if !ok {
		return fmt.Errorf("not a string")
	}

	return t.unmarshal(s)
}
