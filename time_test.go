package rfc3339date_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/boundedinfinity/rfc3339date"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestTime(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "RFC3339 Time Suite")
}

var _ = Describe("RFC3339 Time", func() {
	Context("parse", func() {
		It("should parse time", func() {
			input := TEST_TIME
			expected := testTime()
			actual, err := rfc3339date.ParseTime(input)

			Expect(err).To(BeNil())
			compareTime(actual, expected)
		})
	})

	Context("string", func() {
		It("should create string", func() {
			actual := testTime().String()
			expected := TEST_TIME

			Expect(actual).To(Equal(expected))
		})
	})

	Context("json", func() {
		It("should unmarshal time", func() {
			input := fmt.Sprintf(`"%v"`, TEST_TIME)
			expected := testTime()
			var actual rfc3339date.Rfc3339Time
			err := json.Unmarshal([]byte(input), &actual)

			Expect(err).To(BeNil())
			compareTime(actual, expected)
		})

		It("should marshal time", func() {
			input := testTime()
			expected := fmt.Sprintf(`"%v"`, TEST_TIME)
			bs, err := json.Marshal(input)
			actual := string(bs)

			Expect(err).To(BeNil())
			Expect(actual).To(Equal(expected))
		})
	})
})
