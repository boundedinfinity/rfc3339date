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

func (t Rfc3339DateTime) IsZero() bool {
	return t == DateTimes.Zero
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

func (t *Rfc3339DateTime) unmarshalDateTime(fn func(*string) error) error {
	var s string
	var v Rfc3339DateTime
	var err error

	err = fn(&s)

	if err != nil {
		return err
	}

	if s == "" {
		v = NewDateTime(DateTimes.Zero.Time)
	} else {
		v, err = DateTimes.Parse(s)
	}

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
	var bs []byte
	var err error

	if t == DateTimes.Zero {
		bs, err = json.Marshal(nil)
	} else {
		bs, err = json.Marshal(t.String())
	}

	return bs, err
}

func (t *Rfc3339DateTime) UnmarshalJSON(data []byte) error {
	return t.unmarshalDateTime(func(ptr *string) error {
		return json.Unmarshal(data, ptr)
	})
}

// /////////////////////////////////////////////////////////////////
//  XML marshal/unmarshal implemenation
// /////////////////////////////////////////////////////////////////

func (t Rfc3339DateTime) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if t == DateTimes.Zero {
		return e.EncodeElement("", start)
	} else {
		return e.EncodeElement(t.String(), start)
	}
}

func (t *Rfc3339DateTime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	return t.unmarshalDateTime(func(ptr *string) error {
		return d.DecodeElement(ptr, &start)
	})
}

// /////////////////////////////////////////////////////////////////
//  YAML marshal/unmarshal implemenation
// /////////////////////////////////////////////////////////////////

func (t Rfc3339DateTime) MarshalYAML() (interface{}, error) {
	if t == DateTimes.Zero {
		return nil, nil
	}

	return t.String(), nil
}

func (t *Rfc3339DateTime) UnmarshalYAML(value *yaml.Node) error {
	return t.unmarshalDateTime(func(ptr *string) error {
		return value.Decode(ptr)
	})
}

// /////////////////////////////////////////////////////////////////
//  SQL marshal/unmarshal implementation
// /////////////////////////////////////////////////////////////////

func (t Rfc3339DateTime) Value() (driver.Value, error) {
	return t.Date, nil
}

func (t *Rfc3339DateTime) Scan(value interface{}) error {
	return t.unmarshalDateTime(func(ptr *string) error {
		dv, err := driver.String.ConvertValue(value)

		if err != nil {
			return err
		}

		s, ok := dv.(string)

		if !ok {
			return fmt.Errorf("not a string")
		}

		*ptr = s
		return nil
	})
}
