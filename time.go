package rfc3339date

import (
	"database/sql/driver"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"strings"
	"time"

	"github.com/boundedinfinity/rfc3339date/internal"
	"gopkg.in/yaml.v3"
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

// /////////////////////////////////////////////////////////////////
//  unmarshal implementation
// /////////////////////////////////////////////////////////////////

func (t *Rfc3339Time) unmarshal(s string) error {
	v, err := Times.Parse(s)

	if err != nil {
		return err
	}

	*t = v

	return nil
}

// /////////////////////////////////////////////////////////////////
//  JSON marshal/unmarshal implementation
// /////////////////////////////////////////////////////////////////

func (t Rfc3339Time) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func (t *Rfc3339Time) UnmarshalJSON(data []byte) error {
	var s string

	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	return t.unmarshal(s)
}

// /////////////////////////////////////////////////////////////////
//  XML marshal/unmarshal implementation
// /////////////////////////////////////////////////////////////////

func (t Rfc3339Time) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(t.String(), start)
}

func (t *Rfc3339Time) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var s string

	if err := d.DecodeElement(&s, &start); err != nil {
		return err
	}

	return t.unmarshal(s)
}

// /////////////////////////////////////////////////////////////////
//  YAML marshal/unmarshal implementation
// /////////////////////////////////////////////////////////////////

func (t Rfc3339Time) MarshalYAML() (interface{}, error) {
	return t.String(), nil
}

func (t *Rfc3339Time) UnmarshalYAML(value *yaml.Node) error {
	var s string

	if err := value.Decode(&s); err != nil {
		return err
	}

	return t.unmarshal(s)
}

// /////////////////////////////////////////////////////////////////
//  SQL marshal/unmarshal implementation
// /////////////////////////////////////////////////////////////////

func (t Rfc3339Time) Value() (driver.Value, error) {
	return t.Date, nil
}

func (t *Rfc3339Time) Scan(value interface{}) error {
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
