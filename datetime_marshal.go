package rfc3339date

import (
	"encoding/json"
	"encoding/xml"

	"gopkg.in/yaml.v3"
)

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
