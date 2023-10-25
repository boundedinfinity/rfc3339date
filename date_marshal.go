package rfc3339date

import (
	"encoding/json"
	"encoding/xml"

	"gopkg.in/yaml.v3"
)

func (t *Rfc3339Date) unmarshalDate(fn func(*string) error) error {
	var s string
	var v Rfc3339Date
	var err error

	err = fn(&s)

	if err != nil {
		return err
	}

	if s == "" {
		v = NewDate(ZeroDate.Time)
	} else {
		v, err = ParseDate(s)
	}

	if err != nil {
		return err
	}

	*t = v

	return nil
}

func (t Rfc3339Date) MarshalJSON() ([]byte, error) {
	var bs []byte
	var err error

	if t == ZeroDate {
		bs, err = json.Marshal(nil)
	} else {
		bs, err = json.Marshal(t.String())
	}

	return bs, err
}

func (t *Rfc3339Date) UnmarshalJSON(data []byte) error {
	return t.unmarshalDate(func(ptr *string) error {
		return json.Unmarshal(data, ptr)
	})
}

func (t Rfc3339Date) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if t == ZeroDate {
		return e.EncodeElement("", start)
	} else {
		return e.EncodeElement(t.String(), start)
	}
}

func (t *Rfc3339Date) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	return t.unmarshalDate(func(ptr *string) error {
		return d.DecodeElement(ptr, &start)
	})
}

func (t Rfc3339Date) MarshalYAML() (interface{}, error) {
	if t == ZeroDate {
		return nil, nil
	}

	return t.String(), nil
}

func (t *Rfc3339Date) UnmarshalYAML(value *yaml.Node) error {
	return t.unmarshalDate(func(ptr *string) error {
		return value.Decode(ptr)
	})
}
