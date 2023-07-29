package rrule

import (
	"fmt"
	"strings"
	"time"
)

// Note that the RFC allow ByDay carry preceded by a integer, e.g. 1MO, 2MO, -1MO, -2MO.
// However, the purpose of this implementation is a library for https://github.com/volkanunsal/postgres-rrule.
// So, regarding the extension implementation, we don't support this feature here.
type RRule struct {
	Frequency  Frequency      `json:"freq"`
	Interval   int            `json:"interval"`
	Count      int            `json:"count,omitempty"`
	Until      time.Time      `json:"until,omitempty"`
	BySecond   []int          `json:"bysecond,omitempty"`
	ByMinute   []int          `json:"byminute,omitempty"`
	ByHour     []int          `json:"byhour,omitempty"`
	ByDay      []time.Weekday `json:"byday,omitempty"`
	ByMonthDay []int          `json:"bymonthday,omitempty"`
	ByYearDay  []int          `json:"byyearday,omitempty"`
	ByWeekNo   []int          `json:"byweekno,omitempty"`
	ByMonth    []int          `json:"bymonth,omitempty"`
	BySetpos   []int          `json:"bysetpos,omitempty"`
	WeekStart  time.Weekday   `json:"wkst,omitempty"` // default: Monday
}

func New(freq Frequency, interval int, opts ...Option) (*RRule, error) {
	// default values
	if freq < Yearly || freq > Daily {
		freq = Yearly
	}
	// Regarding the RFC 5545 3.3.10, this field could be omitted.
	// However, in postgres-rrule implementation, the value CANNOT less than 1.
	if interval < 1 {
		interval = 1
	}
	t := &RRule{
		Frequency: freq,
		Interval:  interval,
	}
	// apply options
	for _, opt := range opts {
		opt(t)
	}

	if err := t.Validate(); err != nil {
		return nil, err
	}

	return t, nil
}

func (t *RRule) Validate() error {
	if t.Frequency < Yearly || t.Frequency > Daily {
		return ErrInvalidFrequency
	}
	if t.Interval < 0 {
		return ErrIntervalLessZero
	}

	// Define by RFC 5545 3.3.10
	if t.Count != 0 && !t.Until.IsZero() {
		return fmt.Errorf("%w: COUNT and UNTIL cannot be used together", ErrRuleConflict)
	}

	for _, b := range []struct {
		field     []int
		param     string
		bound     []int
		plusMinus bool // If the bound also applies for -x to -y.
	}{
		{t.BySecond, "bysecond", []int{0, 59}, false},
		{t.ByMinute, "byminute", []int{0, 59}, false},
		{t.ByHour, "byhour", []int{0, 23}, false},
		{t.ByMonthDay, "bymonthday", []int{1, 31}, true},
		{t.ByYearDay, "byyearday", []int{1, 366}, true},
		{t.ByWeekNo, "byweekno", []int{1, 53}, true},
		{t.ByMonth, "bymonth", []int{1, 12}, false},
		{t.BySetpos, "bysetpos", []int{1, 366}, true},
	} {
		for _, value := range b.field {
			if err := checkBounds(b.param, value, b.bound, b.plusMinus); err != nil {
				return err
			}
		}
	}

	// CONSTRAINT freq_yearly_if_byweekno CHECK("freq" = 'YEARLY' OR "byweekno" IS NULL)
	if t.Frequency != Yearly && len(t.ByWeekNo) > 0 {
		return ErrInvalidFrequency
	}

	return nil
}

func (t *RRule) String() string {
	result := []string{fmt.Sprintf("FREQ=%s", t.Frequency)}
	if t.Interval != 0 {
		result = append(result, fmt.Sprintf("INTERVAL=%d", t.Interval))
	}
	if t.WeekStart != time.Monday {
		result = append(result, fmt.Sprintf("WKST=%s", toWeekdayString(t.WeekStart)))
	}
	if t.Count != 0 {
		result = append(result, fmt.Sprintf("COUNT=%d", t.Count))
	}
	if !t.Until.IsZero() {
		result = append(result, fmt.Sprintf("UNTIL=%s", toRFC5545(t.Until)))
	}

	result = appendOption(result, "BYSECOND", t.BySecond)
	result = appendOption(result, "BYMINUTE", t.ByMinute)
	result = appendOption(result, "BYHOUR", t.ByHour)
	if len(t.ByDay) != 0 {
		slice := make([]string, len(t.ByDay))
		for i, wday := range t.ByDay {
			slice[i] = wday.String()
		}
		result = append(result, fmt.Sprintf("BYDAY=%s", strings.Join(slice, ",")))
	}
	result = appendOption(result, "BYMONTHDAY", t.ByMonthDay)
	result = appendOption(result, "BYYEARDAY", t.ByYearDay)
	result = appendOption(result, "BYWEEKNO", t.ByWeekNo)
	result = appendOption(result, "BYMONTH", t.ByMonth)
	result = appendOption(result, "BYSETPOS", t.BySetpos)

	return strings.Join(result, ";")
}

// The interface of sql.scanner.
func (t *RRule) Scan(src interface{}) error {
	return nil
}

// The interface of sql.valuer.
func (t *RRule) Value() (interface{}, error) {
	s := []string{}
	s = append(s, t.Frequency.String())

}
