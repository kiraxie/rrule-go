package rrule

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"
)

func (t RRule) Value() (driver.Value, error) {
	s := []string{}
	s = append(s, t.freq.String())
	if t.interval != 0 {
		s = append(s, fmt.Sprintf("%d", t.interval))
	} else {
		s = append(s, "")
	}
	if t.count != 0 {
		s = append(s, fmt.Sprintf("%d", t.count))
	} else {
		s = append(s, "")
	}
	if !t.until.IsZero() {
		s = append(s, fmt.Sprintf("\"%s\"", t.until.Format(time.DateTime)))
	} else {
		s = append(s, "")
	}
	if len(t.bysecond) != 0 {
		s = append(s, fmt.Sprintf("\"{%s}\"", strings.Join(intSliceToStringSlice(t.bysecond), ",")))
	} else {
		s = append(s, "")
	}
	if len(t.byminute) != 0 {
		s = append(s, fmt.Sprintf("\"{%s}\"", strings.Join(intSliceToStringSlice(t.byminute), ",")))
	} else {
		s = append(s, "")
	}
	if len(t.byhour) != 0 {
		s = append(s, fmt.Sprintf("\"{%s}\"", strings.Join(intSliceToStringSlice(t.byhour), ",")))
	} else {
		s = append(s, "")
	}
	if len(t.bynweekday) != 0 {
		s = append(s, fmt.Sprintf("\"{%s}\"", strings.Join(weekdaySliceToStringSlice(t.bynweekday), ",")))
	} else {
		s = append(s, "")
	}
	if len(t.bymonthday) != 0 {
		s = append(s, fmt.Sprintf("\"{%s}\"", strings.Join(intSliceToStringSlice(t.bymonthday), ",")))
	} else {
		s = append(s, "")
	}
	if len(t.byyearday) != 0 {
		s = append(s, fmt.Sprintf("\"{%s}\"", strings.Join(intSliceToStringSlice(t.byyearday), ",")))
	} else {
		s = append(s, "")
	}
	if len(t.byweekno) != 0 {
		s = append(s, fmt.Sprintf("\"{%s}\"", strings.Join(intSliceToStringSlice(t.byweekno), ",")))
	} else {
		s = append(s, "")
	}
	if len(t.bymonth) != 0 {
		s = append(s, fmt.Sprintf("\"{%s}\"", strings.Join(intSliceToStringSlice(t.bymonth), ",")))
	} else {
		s = append(s, "")
	}
	if len(t.bysetpos) != 0 {
		s = append(s, fmt.Sprintf("\"{%s}\"", strings.Join(intSliceToStringSlice(t.bysetpos), ",")))
	} else {
		s = append(s, "")
	}
	s = append(s, Weekday{weekday: t.wkst}.String())

	return fmt.Sprintf("(%s)", strings.Join(s, ",")), nil
}

func (t *RRule) Scan(value interface{}) (err error) {
	s := strings.Trim(value.(string), "()")
	values := strings.Split(s, ",")
	if len(values) != 14 {
		return ErrInvalidRRuleFormat
	}
	opt := ROption{}
	if err = opt.Freq.Parse(values[0]); err != nil {
		return
	}
	if opt.Interval, err = parseInt(values[1]); err != nil {
		return
	}
	if opt.Count, err = parseInt(values[2]); err != nil {
		return
	}
	if opt.Until, err = parseDate(values[3]); err != nil {
		return
	}
	if opt.Bysecond, err = parseIntSlice(values[4]); err != nil {
		return
	}
	if opt.Byminute, err = parseIntSlice(values[5]); err != nil {
		return
	}
	if opt.Byhour, err = parseIntSlice(values[6]); err != nil {
		return
	}
	if opt.Byweekday, err = parseWeekdaySlice(values[7]); err != nil {
		return
	}
	if opt.Bymonthday, err = parseIntSlice(values[8]); err != nil {
		return
	}
	if opt.Byyearday, err = parseIntSlice(values[9]); err != nil {
		return
	}
	if opt.Byweekno, err = parseIntSlice(values[10]); err != nil {
		return
	}
	if opt.Bymonth, err = parseIntSlice(values[11]); err != nil {
		return
	}
	if opt.Bysetpos, err = parseIntSlice(values[12]); err != nil {
		return
	}
	if err = opt.Wkst.Parse(values[13]); err != nil {
		return
	}

	v, err := NewRRule(opt)
	if err != nil {
		return
	}
	*t = *v

	return nil
}

func (t Set) Value() (driver.Value, error) {
	return "", nil
}

func splitRRuleSetValue(s string) (element []string, err error) {
	symbols := []string{}
	for prev, pos := 0, 0; pos < len(s); pos++ {
		switch s[pos] {
		case ',':
			if len(symbols) == 0 {
				element = append(element, s[prev:pos])
				prev = pos + 1
			}
		case '(':
			if len(element) == 0 { // the front of the string
				prev = pos + 1

				continue
			}
			symbols = append(symbols, "(")
			prev = pos
		case ')':
			switch {
			case len(element) == 5 && len(symbols) == 0: // the end of the string
				element = append(element, s[prev:pos])
			case len(symbols) == 0 || symbols[len(symbols)-1] != "(":
				return nil, fmt.Errorf("%w: symbol \")\" %s", ErrInvalidRRuleFormat, s)
			default:
				symbols = symbols[:len(symbols)-1]
				element = append(element, s[prev:pos+1])
				pos += 2
				prev = pos + 1
			}
		case '{':
			symbols = append(symbols, "{")
		case '}':
			if len(symbols) == 0 || symbols[len(symbols)-1] != "{" {
				return nil, fmt.Errorf("%w: symbol \"}\" %s", ErrInvalidRRuleFormat, s)
			}
			symbols = symbols[:len(symbols)-1]
			element = append(element, s[prev:pos+1])
			pos += 2
			prev = pos + 1
		default:
		}
	}
	if len(element) != 6 {
		return nil, fmt.Errorf("%w: %s(%d)", ErrInvalidRRuleFormat, s, len(element))
	}

	return element, nil
}

func (t *Set) Scan(value interface{}) (err error) {
	var element []string
	if element, err = splitRRuleSetValue(value.(string)); err != nil {
		return
	}
	if t.dtstart, err = parseDate(element[0]); err != nil {
		return
	}
	if element[2] != "" {
		t.rrule = &RRule{}
		if err = t.rrule.Scan(element[2]); err != nil {
			return
		}
	}
	if t.rdate, err = parseDateSlice(element[4]); err != nil {
		return
	}
	if t.exdate, err = parseDateSlice(element[5]); err != nil {
		return
	}

	return nil
}
