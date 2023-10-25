package rfc3339date

import (
	"encoding/json"
	"encoding/xml"

	"gopkg.in/yaml.v3"
)

func (t *Rfc3339Time) unmarshalTime(fn func(*string) error) error {
	var s string
	var v Rfc3339Time
	var err error

	err = fn(&s)

	if err != nil {
		return err
	}

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

func (t *Rfc3339Time) UnmarshalJSON(data []byte) error {
	return t.unmarshalTime(func(ptr *string) error {
		return json.Unmarshal(data, ptr)
	})
}

func (t Rfc3339Time) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if t == ZeroTime {
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

func (t Rfc3339Time) MarshalYAML() (interface{}, error) {
	if t == ZeroTime {
		return nil, nil
	}

	return t.String(), nil
}

func (t *Rfc3339Time) UnmarshalYAML(value *yaml.Node) error {
	return t.unmarshalTime(func(ptr *string) error {
		return value.Decode(ptr)
	})
}
