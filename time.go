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

func (t Rfc3339Time) IsZero() bool {
	return t == Times.Zero
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

func (t *Rfc3339Time) unmarshalTime(fn func(*string) error) error {
	var s string
	var v Rfc3339Time
	var err error

	err = fn(&s)

	if err != nil {
		return err
	}

	if s == "" {
		v = NewTime(Times.Zero.Time)
	} else {
		v, err = Times.Parse(s)
	}

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
	var bs []byte
	var err error

	if t == Times.Zero {
		bs, err = json.Marshal(nil)
	} else {
		bs, err = json.Marshal(t.String())
	}

	return bs, err
}

func (t *Rfc3339Time) UnmarshalJSON(data []byte) error {
	return t.unmarshalTime(func(ptr *string) error {
		return json.Unmarshal(data, ptr)
	})
}

// /////////////////////////////////////////////////////////////////
//  XML marshal/unmarshal implementation
// /////////////////////////////////////////////////////////////////

func (t Rfc3339Time) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if t == Times.Zero {
		return e.EncodeElement("", start)
	} else {
		return e.EncodeElement(t.String(), start)
	}
}

func (t *Rfc3339Time) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	return t.unmarshalTime(func(ptr *string) error {
		return d.DecodeElement(ptr, &start)
	})
}

// /////////////////////////////////////////////////////////////////
//  YAML marshal/unmarshal implementation
// /////////////////////////////////////////////////////////////////

func (t Rfc3339Time) MarshalYAML() (interface{}, error) {
	if t == Times.Zero {
		return nil, nil
	}

	return t.String(), nil
}

func (t *Rfc3339Time) UnmarshalYAML(value *yaml.Node) error {
	return t.unmarshalTime(func(ptr *string) error {
		return value.Decode(ptr)
	})
}

// /////////////////////////////////////////////////////////////////
//  SQL marshal/unmarshal implementation
// /////////////////////////////////////////////////////////////////

func (t Rfc3339Time) Value() (driver.Value, error) {
	return t.Date, nil
}

func (t *Rfc3339Time) Scan(value interface{}) error {
	return t.unmarshalTime(func(ptr *string) error {
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
