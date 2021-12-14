package rfc3339date_test

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"testing"

	"github.com/boundedinfinity/rfc3339date"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gopkg.in/yaml.v3"
)

func TestDate(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "RFC3339 Date Suite")
}

var _ = Describe("RFC3339 Date", func() {
	Context("parse", func() {
		It("should parse date", func() {
			input := TEST_DATE
			expected := testDate()
			actual, err := rfc3339date.ParseDate(input)

			Expect(err).To(BeNil())
			compareDate(actual, expected)
		})
	})

	Context("string", func() {
		It("should create string", func() {
			actual := testDate().String()
			expected := TEST_DATE

			Expect(actual).To(Equal(expected))
		})
	})

	Context("json", func() {
		It("should unmarshal date", func() {
			input := fmt.Sprintf(`"%v"`, TEST_DATE)
			expected := testDate()
			var actual rfc3339date.Rfc3339Date
			err := json.Unmarshal([]byte(input), &actual)

			Expect(err).To(BeNil())
			compareDate(actual, expected)
		})

		It("should marshal date", func() {
			input := testDate()
			expected := fmt.Sprintf(`"%v"`, TEST_DATE)
			bs, err := json.Marshal(input)
			actual := string(bs)

			Expect(err).To(BeNil())
			Expect(actual).To(Equal(expected))
		})
	})

	Context("xml", func() {
		It("should unmarshal date", func() {
			input := fmt.Sprintf(`<Rfc3339Date>%v</Rfc3339Date>`, TEST_DATE)
			expected := testDate()
			var actual rfc3339date.Rfc3339Date
			err := xml.Unmarshal([]byte(input), &actual)

			Expect(err).To(BeNil())
			compareDate(actual, expected)
		})

		It("should marshal date", func() {
			input := testDate()
			expected := fmt.Sprintf(`<Rfc3339Date>%v</Rfc3339Date>`, TEST_DATE)
			bs, err := xml.Marshal(input)
			actual := string(bs)

			Expect(err).To(BeNil())
			Expect(actual).To(Equal(expected))
		})
	})

	Context("yaml", func() {
		It("should unmarshal date", func() {
			input := fmt.Sprintf(`%v`, TEST_DATE)
			expected := testDate()
			var actual rfc3339date.Rfc3339Date
			err := yaml.Unmarshal([]byte(input), &actual)

			Expect(err).To(BeNil())
			compareDate(actual, expected)
		})

		It("should marshal date", func() {
			input := testDate()
			expected := fmt.Sprintf("\"%v\"\n", TEST_DATE)
			bs, err := yaml.Marshal(input)
			actual := string(bs)

			Expect(err).To(BeNil())
			Expect(actual).To(Equal(expected))
		})
	})
})
