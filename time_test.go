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

func Test_Time_Parse(t *testing.T) {
	input := TEST_TIME
	expected := testTime()
	actual, err := rfc3339date.Times.Parse(input)

	assert.Nil(t, err)
	compareTime(t, actual, expected)
}

func Test_Time_String(t *testing.T) {
	actual := testTime().String()
	expected := TEST_TIME

	assert.Equal(t, actual, expected)
}

func Test_Time_Unmarshal_Json(t *testing.T) {
	input := fmt.Sprintf(`"%v"`, TEST_TIME)
	expected := testTime()
	var actual rfc3339date.Rfc3339Time
	err := json.Unmarshal([]byte(input), &actual)

	assert.Nil(t, err)
	compareTime(t, actual, expected)
}

func Test_Time_Marshal_Json(t *testing.T) {
	input := testTime()
	expected := fmt.Sprintf(`"%v"`, TEST_TIME)
	bs, err := json.Marshal(input)
	actual := string(bs)

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func Test_Time_Unmarshal_Xml(t *testing.T) {
	input := fmt.Sprintf(`<Rfc3339Time>%v</Rfc3339Time>`, TEST_TIME)
	expected := testTime()
	var actual rfc3339date.Rfc3339Time
	err := xml.Unmarshal([]byte(input), &actual)

	assert.Nil(t, err)
	compareTime(t, actual, expected)
}

func Test_Time_Marshal_Xml(t *testing.T) {
	input := testTime()
	expected := fmt.Sprintf(`<Rfc3339Time>%v</Rfc3339Time>`, TEST_TIME)
	bs, err := xml.Marshal(input)
	actual := string(bs)

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func Test_Time_Unmarshal_Yaml(t *testing.T) {
	input := fmt.Sprintf(`%v`, TEST_TIME)
	expected := testTime()
	var actual rfc3339date.Rfc3339Time
	err := yaml.Unmarshal([]byte(input), &actual)

	assert.Nil(t, err)
	compareTime(t, actual, expected)
}

func Test_Time_Marshal_Yaml(t *testing.T) {
	input := testTime()
	expected := fmt.Sprintf("\"%v\"\n", TEST_TIME)
	bs, err := yaml.Marshal(input)
	actual := string(bs)

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}
