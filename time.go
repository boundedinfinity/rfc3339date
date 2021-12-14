package rfc3339date

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"strings"
	"time"

	"gopkg.in/yaml.v3"
)

type Rfc3339Time struct {
	time.Time
}

func (t Rfc3339Time) String() string {
	return strings.TrimSuffix(t.Format(_FORMAT_TIME), "Z")
}

func NewTime(d time.Time) Rfc3339Time {
	return Rfc3339Time{d}
}

func ParseTime(s string) (Rfc3339Time, error) {
	s2 := fmt.Sprintf("%v%v%v", _FAKE_DATE1, s, _FAKE_DATE2)
	d, err := time.Parse(_FORMAT_DATETIME, s2)

	if err != nil {
		return Rfc3339Time{}, err
	}

	return Rfc3339Time{d}, nil
}

func (t Rfc3339Time) MarshalJSON() ([]byte, error) {
	bs, err := json.Marshal(t.String())

	if err != nil {
		return nil, err
	}

	return bs, nil
}

func (t *Rfc3339Time) UnmarshalJSON(data []byte) error {
	var s string

	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	v, err := ParseTime(s)

	if err != nil {
		return err
	}

	*t = v

	return nil
}

func (t Rfc3339Time) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(t.String(), start)
}

func (t *Rfc3339Time) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var s string

	if err := d.DecodeElement(&s, &start); err != nil {
		return err
	}

	v, err := ParseTime(s)

	if err != nil {
		return err
	}

	*t = v

	return nil
}

func (t Rfc3339Time) MarshalYAML() (interface{}, error) {
	return t.String(), nil
}

func (t *Rfc3339Time) UnmarshalYAML(value *yaml.Node) error {
	var s string

	if err := value.Decode(&s); err != nil {
		return err
	}

	v, err := ParseTime(s)

	if err != nil {
		return err
	}

	*t = v

	return nil
}
