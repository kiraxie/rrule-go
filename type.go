package rrule

import (
	"fmt"
	"strconv"
)

// Frequency denotes the period on which the rule is evaluated.
type Frequency int

// Constants
const (
	Yearly Frequency = iota
	Monthly
	Weekly
	Daily
	Hourly
	Minutely
	Secondly
)

func (t *Frequency) Parse(src any) error {
	switch v := src.(type) {
	case string:
		switch v {
		case "YEARLY":
			*t = Yearly
		case "MONTHLY":
			*t = Monthly
		case "WEEKLY":
			*t = Weekly
		case "DAILY":
			*t = Daily
		case "HOURLY":
			*t = Hourly
		case "MINUTELY":
			*t = Minutely
		case "SECONDLY":
			*t = Secondly
		default:
			return fmt.Errorf("%w: %s", ErrInvalidFreq, v)
		}
	case int:
		*t = Frequency(v)
	default:
		return fmt.Errorf("%w: %s", ErrInvalidFreq, src)
	}

	return nil
}

func (t Frequency) String() string {
	switch t {
	case Secondly:
		return "SECONDLY"
	case Minutely:
		return "MINUTELY"
	case Hourly:
		return "HOURLY"
	case Daily:
		return "DAILY"
	case Weekly:
		return "WEEKLY"
	case Monthly:
		return "MONTHLY"
	case Yearly:
		return "YEARLY"
	default:
		return "UNKNOWN"
	}
}

// Weekday specifying the nth weekday.
// Field N could be positive or negative (like MO(+2) or MO(-3).
// Not specifying N (0) is the same as specifying +1.
type Weekday struct {
	weekday int
	n       int
}

// Weekdays
var (
	Monday    = Weekday{weekday: 0}
	Tuesday   = Weekday{weekday: 1}
	Wednesday = Weekday{weekday: 2}
	Thursday  = Weekday{weekday: 3}
	Friday    = Weekday{weekday: 4}
	Saturday  = Weekday{weekday: 5}
	Sunday    = Weekday{weekday: 6}
)

// Nth return the nth weekday
// __call__ - Cannot call the object directly,
// do it through e.g. TH.nth(-1) instead,
func (t *Weekday) Nth(n int) Weekday {
	return Weekday{t.weekday, n}
}

// N returns index of the week, e.g. for 3MO, N() will return 3
func (t *Weekday) N() int {
	return t.n
}

// Day returns index of the day in a week (0 for MO, 6 for SU)
func (t *Weekday) Day() int {
	return t.weekday
}

func (t Weekday) String() string {
	s := ""
	switch t.weekday {
	case 0:
		s = "MO"
	case 1:
		s = "TU"
	case 2:
		s = "WE"
	case 3:
		s = "TH"
	case 4:
		s = "FR"
	case 5:
		s = "SA"
	case 6:
		s = "SU"
	default:
		return "UNKNOWN"
	}
	if t.n == 0 {
		return s
	}
	return fmt.Sprintf("%+d%s", t.n, s)
}

func (t *Weekday) Parse(s string) error {
	if len(s) < 2 {
		return ErrInvalidWeekday
	}
	switch s[len(s)-2:] {
	case "MO":
		*t = Monday
	case "TU":
		*t = Tuesday
	case "WE":
		*t = Wednesday
	case "TH":
		*t = Thursday
	case "FR":
		*t = Friday
	case "SA":
		*t = Saturday
	case "SU":
		*t = Sunday
	default:
		return fmt.Errorf("%w: %s", ErrInvalidWeekday, s)
	}
	if len(s) > 2 {
		n, err := strconv.Atoi(s[:len(s)-2])
		if err != nil {
			return err
		}
		t.n = n
	}

	return nil
}
