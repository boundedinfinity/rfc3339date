package rfc3339date

import (
	"fmt"
	"time"
)

const (
	_FORMAT_DATETIME = time.RFC3339
	_FORMAT_TIME     = "15:04:05Z07:00"
	_FORMAT_DATE     = "2006-01-02"
)

var (
	_FAKE_TIME  = fmt.Sprintf("T%v", ZeroDateTime().Format(_FORMAT_TIME))
	_FAKE_DATE1 = fmt.Sprintf("%vT", ZeroDateTime().Format(_FORMAT_DATE))
	_FAKE_DATE2 = "Z"
)
