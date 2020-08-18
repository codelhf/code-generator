package util

import (
	"time"
)

func DateToStr(datetime time.Time) string {
	return DateToStrFormat(datetime, "2006-01-02 15:04:05")
}

func DateToStrFormat(datetime time.Time, layout string) string {
	if IsBlank(layout) {
		layout = "2006-01-02 15:04:05"
	}
	return datetime.Format(layout)
}
