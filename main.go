package rfc3339date

import (
	"strings"

	"github.com/boundedinfinity/rfc3339date/internal"
)

func TranslateFormat(s string) (string, error) {
	o := s

	for a, b := range internal.FORMAT_MAP {
		strings.ReplaceAll(o, a, b)
	}

	return o, nil
}
