package rfc3339date

import (
	"encoding/json"
	"encoding/xml"
	"time"

	"gopkg.in/yaml.v3"
)

type Rfc3339DateTime struct {
	time.Time
}

func (t Rfc3339DateTime) String() string {
	return t.Format(_FORMAT_DATETIME)
}

func ZeroDateTime() Rfc3339DateTime {
	var zero time.Time
	return NewDateTime(zero)
}

func NewDateTime(d time.Time) Rfc3339DateTime {
	return Rfc3339DateTime{d}
}

func ParseDateTime(s string) (Rfc3339DateTime, error) {
	v, err := time.Parse(_FORMAT_DATETIME, s)

	if err != nil {
		return Rfc3339DateTime{}, err
	}

	return Rfc3339DateTime{v}, nil
}

func (t Rfc3339DateTime) MarshalJSON() ([]byte, error) {
	bs, err := json.Marshal(t.String())

	if err != nil {
		return nil, err
	}

	return bs, nil
}

func (t *Rfc3339DateTime) UnmarshalJSON(data []byte) error {
	var s string

	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	v, err := ParseDateTime(s)

	if err != nil {
		return err
	}

	*t = v

	return nil
}

func (t Rfc3339DateTime) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(t.String(), start)
}

func (t *Rfc3339DateTime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var s string

	if err := d.DecodeElement(&s, &start); err != nil {
		return err
	}

	v, err := ParseDateTime(s)

	if err != nil {
		return err
	}

	*t = v

	return nil
}

func (t Rfc3339DateTime) MarshalYAML() (interface{}, error) {
	return t.String(), nil
}

func (t *Rfc3339DateTime) UnmarshalYAML(value *yaml.Node) error {
	var s string

	if err := value.Decode(&s); err != nil {
		return err
	}

	v, err := ParseDateTime(s)

	if err != nil {
		return err
	}

	*t = v

	return nil
}
