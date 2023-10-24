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

func Test_DateTime_Marshal_Json(t *testing.T) {
	testCases := []struct {
		name     string
		input    rfc3339date.Rfc3339DateTime
		expected string
		err      error
	}{
		{
			name:     "non-zero date",
			input:    testDateTime(),
			expected: fmt.Sprintf(`"%v"`, TEST_DATETIME),
		},
		{
			name:     "zero date",
			input:    rfc3339date.ZeroDateTime,
			expected: "null",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			bs, err := json.Marshal(tc.input)
			actual := string(bs)

			assert.Nil(t, err, tc.name)
			assert.Equal(t, tc.expected, actual, tc.name)
		})
	}
}

func Test_DateTime_Unmarshal_Json(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected rfc3339date.Rfc3339DateTime
		err      error
	}{
		{
			name:     "non-zero date",
			input:    fmt.Sprintf(`"%v"`, TEST_DATETIME),
			expected: testDateTime(),
		},
		{
			name:     "zero date",
			input:    "null",
			expected: rfc3339date.ZeroDateTime,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			var actual rfc3339date.Rfc3339DateTime
			err := json.Unmarshal([]byte(tc.input), &actual)

			assert.Nil(t, err)
			compareDateTime(t, actual, tc.expected)
		})
	}
}

func Test_DateTime_Marshal_Xml(t *testing.T) {
	testCases := []struct {
		name     string
		input    rfc3339date.Rfc3339DateTime
		expected string
		err      error
	}{
		{
			name:     "non-zero date",
			input:    testDateTime(),
			expected: fmt.Sprintf(`<Rfc3339DateTime>%v</Rfc3339DateTime>`, TEST_DATETIME),
		},
		{
			name:     "zero date",
			input:    rfc3339date.ZeroDateTime,
			expected: "<Rfc3339DateTime></Rfc3339DateTime>",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			bs, err := xml.Marshal(tc.input)
			actual := string(bs)

			assert.Nil(t, err)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func Test_DateTime_Unmarshal_Xml(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected rfc3339date.Rfc3339DateTime
		err      error
	}{
		{
			name:     "non-zero date",
			input:    fmt.Sprintf(`<Rfc3339DateTime>%v</Rfc3339DateTime>`, TEST_DATETIME),
			expected: testDateTime(),
		},
		{
			name:     "zero date",
			input:    `<Rfc3339DateTime></Rfc3339DateTime>`,
			expected: rfc3339date.ZeroDateTime,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			var actual rfc3339date.Rfc3339DateTime
			err := xml.Unmarshal([]byte(tc.input), &actual)

			assert.Nil(t, err)
			compareDateTime(t, tc.expected, actual)
		})
	}
}

func Test_DateTime_Marshal_Yaml(t *testing.T) {
	testCases := []struct {
		name     string
		input    rfc3339date.Rfc3339DateTime
		expected string
		err      error
	}{
		{
			name:     "non-zero date",
			input:    testDateTime(),
			expected: fmt.Sprintf("\"%v\"\n", TEST_DATETIME),
		},
		{
			name:     "zero date",
			input:    rfc3339date.ZeroDateTime,
			expected: "null\n",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			bs, err := yaml.Marshal(tc.input)
			actual := string(bs)

			assert.Nil(t, err)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func Test_DateTime_Unmarshal_Yaml(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected rfc3339date.Rfc3339DateTime
		err      error
	}{
		{
			name:     "non-zero date",
			input:    fmt.Sprintf(`"%v"`, TEST_DATETIME),
			expected: testDateTime(),
		},
		{
			name:     "zero date",
			input:    "null",
			expected: rfc3339date.ZeroDateTime,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {

			var actual rfc3339date.Rfc3339DateTime
			err := yaml.Unmarshal([]byte(tc.input), &actual)

			assert.Nil(t, err)
			compareDateTime(t, tc.expected, actual)
		})
	}
}
