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

func NewDateTime(d time.Time) Rfc3339DateTime {
	return Rfc3339DateTime{d}
}

// /////////////////////////////////////////////////////////////////
//  unmarshal implemenation
// /////////////////////////////////////////////////////////////////

func (t *Rfc3339DateTime) unmarshal(s string) error {
	v, err := DateTimes.Parse(s)

	if err != nil {
		return err
	}

	*t = v

	return nil
}

// /////////////////////////////////////////////////////////////////
//  JSON marshal/unmarshal implemenation
// /////////////////////////////////////////////////////////////////

func (t Rfc3339DateTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func (t *Rfc3339DateTime) UnmarshalJSON(data []byte) error {
	var s string

	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	return t.unmarshal(s)
}

// /////////////////////////////////////////////////////////////////
//  XML marshal/unmarshal implemenation
// /////////////////////////////////////////////////////////////////

func (t Rfc3339DateTime) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(t.String(), start)
}

func (t *Rfc3339DateTime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var s string

	if err := d.DecodeElement(&s, &start); err != nil {
		return err
	}

	return t.unmarshal(s)
}

// /////////////////////////////////////////////////////////////////
//  YAML marshal/unmarshal implemenation
// /////////////////////////////////////////////////////////////////

func (t Rfc3339DateTime) MarshalYAML() (interface{}, error) {
	return t.String(), nil
}

func (t *Rfc3339DateTime) UnmarshalYAML(value *yaml.Node) error {
	var s string

	if err := value.Decode(&s); err != nil {
		return err
	}

	return t.unmarshal(s)
}

// /////////////////////////////////////////////////////////////////
//  SQL marshal/unmarshal implementation
// /////////////////////////////////////////////////////////////////

func (t Rfc3339DateTime) Value() (driver.Value, error) {
	return t.Date, nil
}

func (t *Rfc3339DateTime) Scan(value interface{}) error {
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
