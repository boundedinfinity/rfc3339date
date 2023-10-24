package rfc3339date

import (
	"encoding/json"
	"encoding/xml"

	"gopkg.in/yaml.v3"
)

func (t Rfc3339DateTime) MarshalJSON() ([]byte, error) {
	var bs []byte
	var err error

	if t == ZeroDateTime {
		bs, err = json.Marshal(nil)
	} else {
		bs, err = json.Marshal(t.String())
	}

	return bs, err
}

func (t *Rfc3339DateTime) unmarshalJSON(s string) error {
	var v Rfc3339DateTime
	var err error

	if s == "" {
		v = NewDateTime(ZeroDateTime.Time)
	} else {
		v, err = ParseDateTime(s)
	}

	if err != nil {
		return err
	}

	*t = v

	return nil
}

func (t *Rfc3339DateTime) UnmarshalJSON(data []byte) error {
	var s string

	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	return t.unmarshalJSON(s)
}

func (t Rfc3339DateTime) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if t == ZeroDateTime {
		return e.EncodeElement("", start)
	} else {
		return e.EncodeElement(t.String(), start)
	}
}

func (t *Rfc3339DateTime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var s string

	if err := d.DecodeElement(&s, &start); err != nil {
		return err
	}

	return t.unmarshalJSON(s)
}

func (t Rfc3339DateTime) MarshalYAML() (interface{}, error) {
	if t == ZeroDateTime {
		return nil, nil
	}

	return t.String(), nil
}

func (t *Rfc3339DateTime) UnmarshalYAML(value *yaml.Node) error {
	var s string

	if err := value.Decode(&s); err != nil {
		return err
	}

	return t.unmarshalJSON(s)
}
