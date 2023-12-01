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

func Test_Date_Parse(t *testing.T) {
	input := TEST_DATE
	expected := testDate()
	actual, err := rfc3339date.Dates.Parse(input)

	assert.Nil(t, err)
	compareDate(t, actual, expected)
}

func Test_Date_String(t *testing.T) {
	actual := testDate().String()
	expected := TEST_DATE

	assert.Equal(t, actual, expected)
}

func Test_Date_Unmarshal_Json(t *testing.T) {
	input := fmt.Sprintf(`"%v"`, TEST_DATE)
	expected := testDate()
	var actual rfc3339date.Rfc3339Date
	err := json.Unmarshal([]byte(input), &actual)

	assert.Nil(t, err)
	compareDate(t, actual, expected)
}

func Test_Date_Marshal_Json(t *testing.T) {
	input := testDate()
	expected := fmt.Sprintf(`"%v"`, TEST_DATE)
	bs, err := json.Marshal(input)
	actual := string(bs)

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func Test_Date_Unmarshal_Xml(t *testing.T) {
	input := fmt.Sprintf(`<Rfc3339Date>%v</Rfc3339Date>`, TEST_DATE)
	expected := testDate()
	var actual rfc3339date.Rfc3339Date
	err := xml.Unmarshal([]byte(input), &actual)

	assert.Nil(t, err)
	compareDate(t, actual, expected)
}

func Test_Date_Marshal_Xml(t *testing.T) {
	input := testDate()
	expected := fmt.Sprintf(`<Rfc3339Date>%v</Rfc3339Date>`, TEST_DATE)
	bs, err := xml.Marshal(input)
	actual := string(bs)

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func Test_Date_Unmarshal_Yaml(t *testing.T) {
	input := fmt.Sprintf(`%v`, TEST_DATE)
	expected := testDate()
	var actual rfc3339date.Rfc3339Date
	err := yaml.Unmarshal([]byte(input), &actual)

	assert.Nil(t, err)
	compareDate(t, actual, expected)
}

func Test_Date_Marshal_Yaml(t *testing.T) {
	input := testDate()
	expected := fmt.Sprintf("\"%v\"\n", TEST_DATE)
	bs, err := yaml.Marshal(input)
	actual := string(bs)

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}
