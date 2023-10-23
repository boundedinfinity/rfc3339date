package rfc3339date

import (
	"encoding/json"
	"encoding/xml"

	"gopkg.in/yaml.v3"
)

func (t Rfc3339Date) MarshalJSON() ([]byte, error) {
	bs, err := json.Marshal(t.String())

	if err != nil {
		return nil, err
	}

	return bs, nil
}

func (t *Rfc3339Date) UnmarshalJSON(data []byte) error {
	var s string

	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	v, err := ParseDate(s)

	if err != nil {
		return err
	}

	*t = v

	return nil
}

func (t Rfc3339Date) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(t.String(), start)
}

func (t *Rfc3339Date) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var s string

	if err := d.DecodeElement(&s, &start); err != nil {
		return err
	}

	v, err := ParseDate(s)

	if err != nil {
		return err
	}

	*t = v

	return nil
}

func (t Rfc3339Date) MarshalYAML() (interface{}, error) {
	return t.String(), nil
}

func (t *Rfc3339Date) UnmarshalYAML(value *yaml.Node) error {
	var s string

	if err := value.Decode(&s); err != nil {
		return err
	}

	v, err := ParseDate(s)

	if err != nil {
		return err
	}

	*t = v

	return nil
}
