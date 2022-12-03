package internal

import (
	"fmt"
	"time"
)

const (
	FORMAT_DATETIME = time.RFC3339
	FORMAT_TIME     = "15:04:05Z07:00"
	FORMAT_DATE     = "2006-01-02"

	SECONDS_PER_MINUTE = 60
	SECONDS_PER_HOUR   = 60 * SECONDS_PER_MINUTE
	SECONDS_PER_DAY    = 24 * SECONDS_PER_HOUR
	SECONDS_PER_WEEK   = 7 * SECONDS_PER_DAY
	SECONDS_PER_MONTH  = 4 * SECONDS_PER_WEEK
	SECONDS_PER_YEAR   = 12 * SECONDS_PER_MONTH
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

	DURATION_MAP = map[string]int64{
		"s": 1,
		"m": SECONDS_PER_MINUTE,
		"h": SECONDS_PER_MINUTE,
		"D": SECONDS_PER_DAY,
		"W": SECONDS_PER_WEEK,
		"M": SECONDS_PER_MONTH,
		"Y": SECONDS_PER_YEAR,
	}
)

func zero() time.Time {
	var zero time.Time
	return zero
}
