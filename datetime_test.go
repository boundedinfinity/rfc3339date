package rfc3339date_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/boundedinfinity/rfc3339date"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestDateTime(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "RFC3339 DateTime Suite")
}

var _ = Describe("RFC3339 DateTime", func() {
	Context("parse", func() {
		It("should parse date", func() {
			input := TEST_DATETIME
			expected := testDateTime()
			actual, err := rfc3339date.ParseDateTime(input)

			Expect(err).To(BeNil())
			compareDateTime(actual, expected)
		})
	})

	Context("string", func() {
		It("should create string", func() {
			actual := testDateTime().String()
			expected := TEST_DATETIME

			Expect(actual).To(Equal(expected))
		})
	})

	Context("json", func() {
		It("should unmarshal date", func() {
			input := fmt.Sprintf(`"%v"`, TEST_DATETIME)
			expected := testDateTime()
			var actual rfc3339date.Rfc3339DateTime
			err := json.Unmarshal([]byte(input), &actual)

			Expect(err).To(BeNil())
			compareDateTime(actual, expected)
		})

		It("should marshal date", func() {
			input := testDateTime()
			expected := fmt.Sprintf(`"%v"`, TEST_DATETIME)
			bs, err := json.Marshal(input)
			actual := string(bs)

			Expect(err).To(BeNil())
			Expect(actual).To(Equal(expected))
		})
	})
})
