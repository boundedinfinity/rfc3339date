package rfc3339date

import (
	"encoding/json"
	"encoding/xml"

	"gopkg.in/yaml.v3"
)

func (t *Rfc3339DateTime) unmarshalDateTime(fn func(*string) error) error {
	var s string
	var v Rfc3339DateTime
	var err error

	err = fn(&s)

	if err != nil {
		return err
	}

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

func (t *Rfc3339DateTime) UnmarshalJSON(data []byte) error {
	return t.unmarshalDateTime(func(ptr *string) error {
		return json.Unmarshal(data, ptr)
	})
}

func (t Rfc3339DateTime) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if t == ZeroDateTime {
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

func (t Rfc3339DateTime) MarshalYAML() (interface{}, error) {
	if t == ZeroDateTime {
		return nil, nil
	}

	return t.String(), nil
}

func (t *Rfc3339DateTime) UnmarshalYAML(value *yaml.Node) error {
	return t.unmarshalDateTime(func(ptr *string) error {
		return value.Decode(ptr)
	})
}
