package rrule

import (
	"time"
)

const (
	// DateTimeFormat is date-time format used in iCalendar (RFC 5545)
	DateTimeFormat = "20060102T150405Z"
	// LocalDateTimeFormat is a date-time format without Z prefix
	LocalDateTimeFormat = "20060102T150405"
	// DateFormat is date format used in iCalendar (RFC 5545)
	DateFormat = "20060102"
)

func toRFC5545(t time.Time) string {
	return t.Format(DateTimeFormat)
}

func toWeekdayString(v time.Weekday) string {
	switch v {
	case time.Monday:
		return "MO"
	case time.Tuesday:
		return "TU"
	case time.Wednesday:
		return "WE"
	case time.Thursday:
		return "TH"
	case time.Friday:
		return "FR"
	case time.Saturday:
		return "SA"
	case time.Sunday:
		return "SU"
	default:
		return "UNKNOWN"
	}
}
