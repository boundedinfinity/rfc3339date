package rfc3339date

import (
	"encoding/json"
	"encoding/xml"

	"gopkg.in/yaml.v3"
)

func (t Rfc3339Time) MarshalJSON() ([]byte, error) {
	var bs []byte
	var err error

	if t == ZeroTime {
		bs, err = json.Marshal(nil)
	} else {
		bs, err = json.Marshal(t.String())
	}

	return bs, err
}

func (t *Rfc3339Time) unmarshalJSON(s string) error {
	var v Rfc3339Time
	var err error

	if s == "" {
		v = NewTime(ZeroTime.Time)
	} else {
		v, err = ParseTime(s)
	}

	if err != nil {
		return err
	}

	*t = v

	return nil
}

func (t *Rfc3339Time) UnmarshalJSON(data []byte) error {
	var s string

	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	return t.unmarshalJSON(s)
}

func (t Rfc3339Time) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if t == ZeroTime {
		return e.EncodeElement("", start)
	} else {
		return e.EncodeElement(t.String(), start)
	}
}

func (t *Rfc3339Time) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var s string

	if err := d.DecodeElement(&s, &start); err != nil {
		return err
	}

	return t.unmarshalJSON(s)
}

func (t Rfc3339Time) MarshalYAML() (interface{}, error) {
	if t == ZeroTime {
		return nil, nil
	}

	return t.String(), nil
}

func (t *Rfc3339Time) UnmarshalYAML(value *yaml.Node) error {
	var s string

	if err := value.Decode(&s); err != nil {
		return err
	}

	return t.unmarshalJSON(s)
}
