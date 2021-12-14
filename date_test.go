package rfc3339date_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/boundedinfinity/rfc3339date"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
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
})
