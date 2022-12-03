package rfc3339date_test

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"testing"

	"github.com/boundedinfinity/rfc3339date"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func Test_DateTime_Parse(t *testing.T) {
	input := TEST_DATETIME
	expected := testDateTime()
	actual, err := rfc3339date.ParseDateTime(input)

	assert.Nil(t, err)
	compareDateTime(t, actual, expected)
}

func Test_DateTime_String(t *testing.T) {
	actual := testDateTime().String()
	expected := TEST_DATETIME

	assert.Equal(t, actual, expected)
}

func Test_DateTime_Unmarshal_Json(t *testing.T) {
	input := fmt.Sprintf(`"%v"`, TEST_DATETIME)
	expected := testDateTime()
	var actual rfc3339date.Rfc3339DateTime
	err := json.Unmarshal([]byte(input), &actual)

	assert.Nil(t, err)
	compareDateTime(t, actual, expected)
}

func Test_DateTime_Marshal_Json(t *testing.T) {
	input := testDateTime()
	expected := fmt.Sprintf(`"%v"`, TEST_DATETIME)
	bs, err := json.Marshal(input)
	actual := string(bs)

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func Test_DateTime_Unmarshal_Xml(t *testing.T) {
	input := fmt.Sprintf(`<Rfc3339DateTime>%v</Rfc3339DateTime>`, TEST_DATETIME)
	expected := testDateTime()
	var actual rfc3339date.Rfc3339DateTime
	err := xml.Unmarshal([]byte(input), &actual)

	assert.Nil(t, err)
	compareDateTime(t, actual, expected)
}

func Test_DateTime_Marshal_Xml(t *testing.T) {
	input := testDateTime()
	expected := fmt.Sprintf(`<Rfc3339DateTime>%v</Rfc3339DateTime>`, TEST_DATETIME)
	bs, err := xml.Marshal(input)
	actual := string(bs)

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func Test_DateTime_Unmarshal_Yaml(t *testing.T) {
	input := fmt.Sprintf(`%v`, TEST_DATETIME)
	expected := testDateTime()
	var actual rfc3339date.Rfc3339DateTime
	err := yaml.Unmarshal([]byte(input), &actual)

	assert.Nil(t, err)
	compareDateTime(t, actual, expected)
}

func Test_DateTime_Marshal_Yaml(t *testing.T) {
	input := testDateTime()
	expected := fmt.Sprintf("\"%v\"\n", TEST_DATETIME)
	bs, err := yaml.Marshal(input)
	actual := string(bs)

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}
