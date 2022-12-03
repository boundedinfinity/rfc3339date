package rfc3339date

import (
	"encoding/json"
	"encoding/xml"
	"time"

	"gopkg.in/yaml.v3"
)

type Rfc3339Duration struct {
	time.Duration
}

func NewDuration(d time.Duration) Rfc3339Duration {
	return Rfc3339Duration{d}
}

func ParseDuration(f string) Rfc3339Duration {
	var d int64

	return Rfc3339Duration{time.Duration(d)}
}

func (t Rfc3339Duration) String() string {
	return t.Duration.String()
}

func (t Rfc3339Duration) MarshalJSON() ([]byte, error) {
	bs, err := json.Marshal(int64(t.Duration))

	if err != nil {
		return nil, err
	}

	return bs, nil
}

func (t *Rfc3339Duration) UnmarshalJSON(data []byte) error {
	var i int64

	if err := json.Unmarshal(data, &i); err != nil {
		return err
	}

	t.Duration = time.Duration(i)

	return nil
}

func (t Rfc3339Duration) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(int64(t.Duration), start)
}

func (t *Rfc3339Duration) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var i int64

	if err := d.DecodeElement(&i, &start); err != nil {
		return err
	}

	t.Duration = time.Duration(i)

	return nil
}

func (t Rfc3339Duration) MarshalYAML() (interface{}, error) {
	return int64(t.Duration), nil
}

func (t *Rfc3339Duration) UnmarshalYAML(value *yaml.Node) error {
	var i int64

	if err := value.Decode(&i); err != nil {
		return err
	}

	t.Duration = time.Duration(i)

	return nil
}
