package internal

import (
	"fmt"
	"time"
)

const (
	FORMAT_DATETIME = time.RFC3339
	FORMAT_TIME     = "15:04:05Z07:00"
	FORMAT_DATE     = "2006-01-02"
)

var (
	FAKE_TIME  = fmt.Sprintf("T%v", zero().Format(FORMAT_TIME))
	FAKE_DATE1 = fmt.Sprintf("%vT", zero().Format(FORMAT_DATE))
	FAKE_DATE2 = "Z"

	FORMAT_MAP = map[string]string{
		"YYYY": "2006",
		"YY":   "06",
		"MM":   "01",
		"DD":   "02",
		"hh":   "15",
		"mm":   "04",
		"ss":   "05",
		"zz":   "Z07:00",
	}
)

func zero() time.Time {
	var zero time.Time
	return zero
}
