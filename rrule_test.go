// 2017-2022, Teambition. All rights reserved.

package rrule

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var testDtstart = time.Date(1997, 9, 2, 9, 0, 0, 0, time.UTC)

func timesEqual(value, want []time.Time) bool {
	if len(value) != len(want) {
		return false
	}
	for index := range value {
		if value[index] != want[index] {
			return false
		}
	}
	return true
}

func TestNoDtstart(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq: Monthly,
	})
	assert.NoError(t, err)
	assert.WithinDuration(t, time.Now(), r.dtstart, time.Second)
}

func TestBadBySetPos(t *testing.T) {
	t.Parallel()
	_, err := NewRRule(ROption{
		Freq: Monthly, Count: 1, Bysetpos: []int{0}, Dtstart: testDtstart,
	})
	assert.Error(t, err)
}

func TestBadBySetPosMany(t *testing.T) {
	t.Parallel()
	_, err := NewRRule(ROption{
		Freq: Monthly, Count: 1, Bysetpos: []int{-1, 0, 1}, Dtstart: testDtstart,
	})
	assert.Error(t, err)
}

func TestByNegativeMonthDay(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq: Monthly, Count: 3, Bymonthday: []int{-1}, Dtstart: testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1997, 9, 30, 9, 0, 0, 0, time.UTC),
		time.Date(1997, 10, 31, 9, 0, 0, 0, time.UTC),
		time.Date(1997, 11, 30, 9, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.True(t, timesEqual(value, want))
}

func TestMonthlyMaxYear(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq: Monthly, Interval: 15, Dtstart: testDtstart,
	})
	assert.NoError(t, err)
	value := r.All()[1]
	want := time.Date(1998, 12, 2, 9, 0, 0, 0, time.UTC)
	assert.Equal(t, want, value)
}

func TestWeeklyMaxYear(t *testing.T) {
	t.Parallel()
	// Purposefully doesn't match anything for code coverage.
	r, err := NewRRule(ROption{
		Freq: Weekly, Bymonthday: []int{31}, Byyearday: []int{1}, Dtstart: testDtstart,
	})
	assert.NoError(t, err)
	value := r.All()
	want := []time.Time{}
	assert.Equal(t, want, value)
}

func TestInvalidRRules(t *testing.T) {
	t.Parallel()
	tests := []struct {
		desc  string
		rrule ROption
	}{
		{
			desc:  "Bysecond under",
			rrule: ROption{Freq: Yearly, Bysecond: []int{-1}},
		},
		{
			desc:  "Bysecond over",
			rrule: ROption{Freq: Yearly, Bysecond: []int{60}},
		},
		{
			desc:  "Byminute under",
			rrule: ROption{Freq: Yearly, Byminute: []int{-1}},
		},
		{
			desc:  "Byminute over",
			rrule: ROption{Freq: Yearly, Byminute: []int{60}},
		},
		{
			desc:  "Byhour under",
			rrule: ROption{Freq: Yearly, Byhour: []int{-1}},
		},
		{
			desc:  "Byhour over",
			rrule: ROption{Freq: Yearly, Byhour: []int{24}},
		},
		{
			desc:  "Bymonthday under",
			rrule: ROption{Freq: Yearly, Bymonthday: []int{0}},
		},
		{
			desc:  "Bymonthday over",
			rrule: ROption{Freq: Yearly, Bymonthday: []int{32}},
		},
		{
			desc:  "Bymonthday under negative",
			rrule: ROption{Freq: Yearly, Bymonthday: []int{-32}},
		},
		{
			desc:  "Byyearday under",
			rrule: ROption{Freq: Yearly, Byyearday: []int{0}},
		},
		{
			desc:  "Byyearday over",
			rrule: ROption{Freq: Yearly, Byyearday: []int{367}},
		},
		{
			desc:  "Byyearday under negative",
			rrule: ROption{Freq: Yearly, Byyearday: []int{-367}},
		},
		{
			desc:  "Byweekno under",
			rrule: ROption{Freq: Yearly, Byweekno: []int{0}},
		},
		{
			desc:  "Byweekno over",
			rrule: ROption{Freq: Yearly, Byweekno: []int{54}},
		},
		{
			desc:  "Byweekno under negative",
			rrule: ROption{Freq: Yearly, Byweekno: []int{-54}},
		},
		{
			desc:  "Bymonth under",
			rrule: ROption{Freq: Yearly, Bymonth: []int{0}},
		},
		{
			desc:  "Bymonth over",
			rrule: ROption{Freq: Yearly, Bymonth: []int{13}},
		},
		{
			desc:  "Bysetpos under",
			rrule: ROption{Freq: Yearly, Bysetpos: []int{0}},
		},
		{
			desc:  "Bysetpos over",
			rrule: ROption{Freq: Yearly, Bysetpos: []int{367}},
		},
		{
			desc:  "Bysetpos under negative",
			rrule: ROption{Freq: Yearly, Bysetpos: []int{-367}},
		},
		{
			desc:  "Byday under",
			rrule: ROption{Freq: Yearly, Byweekday: []Weekday{{1, -54}}},
		},
		{
			desc:  "Byday over",
			rrule: ROption{Freq: Yearly, Byweekday: []Weekday{{1, 54}}},
		},
		{
			desc:  "Interval under",
			rrule: ROption{Freq: Daily, Interval: -1},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			t.Parallel()
			_, err := NewRRule(tc.rrule)
			assert.ErrorIs(t, err, ErrInvalidateBound)
		})
	}
}

func TestHourlyInvalidAndRepeatedBysetpos(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq: Hourly, Bysetpos: []int{1, -1, 2},
		Dtstart: testDtstart,
		Until:   time.Date(1997, 9, 2, 11, 0, 0, 0, time.UTC),
	})
	assert.NoError(t, err)
	value := r.All()
	want := []time.Time{
		testDtstart,
		time.Date(1997, 9, 2, 10, 0, 0, 0, time.UTC),
		time.Date(1997, 9, 2, 11, 0, 0, 0, time.UTC),
	}
	assert.Equal(t, want, value)
}

func TestNoAfter(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:    Daily,
		Count:   5,
		Dtstart: testDtstart,
	})
	assert.NoError(t, err)
	want := time.Time{}
	value := r.After(time.Date(1997, 9, 6, 9, 0, 0, 0, time.UTC), false)
	assert.Equal(t, want, value)
}

// Test cases from Python Dateutil

func TestYearly(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:    Yearly,
		Count:   3,
		Dtstart: testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		testDtstart,
		time.Date(1998, 9, 2, 9, 0, 0, 0, time.UTC),
		time.Date(1999, 9, 2, 9, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestYearlyInterval(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:     Yearly,
		Count:    3,
		Interval: 2,
		Dtstart:  testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		testDtstart,
		time.Date(1999, 9, 2, 9, 0, 0, 0, time.UTC),
		time.Date(2001, 9, 2, 9, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestYearlyIntervalLarge(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:     Yearly,
		Count:    3,
		Interval: 100,
		Dtstart:  testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		testDtstart,
		time.Date(2097, 9, 2, 9, 0, 0, 0, time.UTC),
		time.Date(2197, 9, 2, 9, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestYearlyByMonth(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:    Yearly,
		Count:   3,
		Bymonth: []int{1, 3},
		Dtstart: testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1998, 1, 2, 9, 0, 0, 0, time.UTC),
		time.Date(1998, 3, 2, 9, 0, 0, 0, time.UTC),
		time.Date(1999, 1, 2, 9, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestYearlyByMonthDay(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:       Yearly,
		Count:      3,
		Bymonthday: []int{1, 3},
		Dtstart:    testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1997, 9, 3, 9, 0, 0, 0, time.UTC),
		time.Date(1997, 10, 1, 9, 0, 0, 0, time.UTC),
		time.Date(1997, 10, 3, 9, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestYearlyByMonthAndMonthDay(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:       Yearly,
		Count:      3,
		Bymonth:    []int{1, 3},
		Bymonthday: []int{5, 7},
		Dtstart:    testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1998, 1, 5, 9, 0, 0, 0, time.UTC),
		time.Date(1998, 1, 7, 9, 0, 0, 0, time.UTC),
		time.Date(1998, 3, 5, 9, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestYearlyByWeekDay(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:      Yearly,
		Count:     3,
		Byweekday: []Weekday{Tuesday, Thursday},
		Dtstart:   testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		testDtstart,
		time.Date(1997, 9, 4, 9, 0, 0, 0, time.UTC),
		time.Date(1997, 9, 9, 9, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestYearlyByNWeekDay(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:      Yearly,
		Count:     3,
		Byweekday: []Weekday{Tuesday.Nth(1), Thursday.Nth(-1)},
		Dtstart:   testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1997, 12, 25, 9, 0, 0, 0, time.UTC),
		time.Date(1998, 1, 6, 9, 0, 0, 0, time.UTC),
		time.Date(1998, 12, 31, 9, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestYearlyByNWeekDayLarge(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:      Yearly,
		Count:     3,
		Byweekday: []Weekday{Tuesday.Nth(3), Thursday.Nth(-3)},
		Dtstart:   testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1997, 12, 11, 9, 0, 0, 0, time.UTC),
		time.Date(1998, 1, 20, 9, 0, 0, 0, time.UTC),
		time.Date(1998, 12, 17, 9, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestYearlyByMonthAndWeekDay(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:      Yearly,
		Count:     3,
		Bymonth:   []int{1, 3},
		Byweekday: []Weekday{Tuesday, Thursday},
		Dtstart:   testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1998, 1, 1, 9, 0, 0, 0, time.UTC),
		time.Date(1998, 1, 6, 9, 0, 0, 0, time.UTC),
		time.Date(1998, 1, 8, 9, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestYearlyByMonthAndNWeekDay(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:      Yearly,
		Count:     3,
		Bymonth:   []int{1, 3},
		Byweekday: []Weekday{Tuesday.Nth(1), Thursday.Nth(-1)},
		Dtstart:   testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1998, 1, 6, 9, 0, 0, 0, time.UTC),
		time.Date(1998, 1, 29, 9, 0, 0, 0, time.UTC),
		time.Date(1998, 3, 3, 9, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestYearlyByMonthAndNWeekDayLarge(t *testing.T) {
	t.Parallel()
	// This is interesting because the TH.Nth(-3) ends up before
	// the TU.Nth(3).
	r, err := NewRRule(ROption{
		Freq:      Yearly,
		Count:     3,
		Bymonth:   []int{1, 3},
		Byweekday: []Weekday{Tuesday.Nth(3), Thursday.Nth(-3)},
		Dtstart:   testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1998, 1, 15, 9, 0, 0, 0, time.UTC),
		time.Date(1998, 1, 20, 9, 0, 0, 0, time.UTC),
		time.Date(1998, 3, 12, 9, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestYearlyByMonthDayAndWeekDay(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:       Yearly,
		Count:      3,
		Bymonthday: []int{1, 3},
		Byweekday:  []Weekday{Tuesday, Thursday},
		Dtstart:    testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1998, 1, 1, 9, 0, 0, 0, time.UTC),
		time.Date(1998, 2, 3, 9, 0, 0, 0, time.UTC),
		time.Date(1998, 3, 3, 9, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestYearlyByMonthAndMonthDayAndWeekDay(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:       Yearly,
		Count:      3,
		Bymonth:    []int{1, 3},
		Bymonthday: []int{1, 3},
		Byweekday:  []Weekday{Tuesday, Thursday},
		Dtstart:    testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1998, 1, 1, 9, 0, 0, 0, time.UTC),
		time.Date(1998, 3, 3, 9, 0, 0, 0, time.UTC),
		time.Date(2001, 3, 1, 9, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestYearlyByYearDay(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:      Yearly,
		Count:     4,
		Byyearday: []int{1, 100, 200, 365},
		Dtstart:   testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1997, 12, 31, 9, 0, 0, 0, time.UTC),
		time.Date(1998, 1, 1, 9, 0, 0, 0, time.UTC),
		time.Date(1998, 4, 10, 9, 0, 0, 0, time.UTC),
		time.Date(1998, 7, 19, 9, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestYearlyByYearDayNeg(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:      Yearly,
		Count:     4,
		Byyearday: []int{-365, -266, -166, -1},
		Dtstart:   testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1997, 12, 31, 9, 0, 0, 0, time.UTC),
		time.Date(1998, 1, 1, 9, 0, 0, 0, time.UTC),
		time.Date(1998, 4, 10, 9, 0, 0, 0, time.UTC),
		time.Date(1998, 7, 19, 9, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestYearlyByMonthAndYearDay(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:      Yearly,
		Count:     4,
		Bymonth:   []int{4, 7},
		Byyearday: []int{1, 100, 200, 365},
		Dtstart:   testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1998, 4, 10, 9, 0, 0, 0, time.UTC),
		time.Date(1998, 7, 19, 9, 0, 0, 0, time.UTC),
		time.Date(1999, 4, 10, 9, 0, 0, 0, time.UTC),
		time.Date(1999, 7, 19, 9, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestYearlyByMonthAndYearDayNeg(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:      Yearly,
		Count:     4,
		Bymonth:   []int{4, 7},
		Byyearday: []int{-365, -266, -166, -1},
		Dtstart:   testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1998, 4, 10, 9, 0, 0, 0, time.UTC),
		time.Date(1998, 7, 19, 9, 0, 0, 0, time.UTC),
		time.Date(1999, 4, 10, 9, 0, 0, 0, time.UTC),
		time.Date(1999, 7, 19, 9, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestYearlyByWeekNo(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:     Yearly,
		Count:    3,
		Byweekno: []int{20},
		Dtstart:  testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1998, 5, 11, 9, 0, 0, 0, time.UTC),
		time.Date(1998, 5, 12, 9, 0, 0, 0, time.UTC),
		time.Date(1998, 5, 13, 9, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestYearlyByWeekNoAndWeekDay(t *testing.T) {
	t.Parallel()
	// That's a nice one. The first days of week number one
	// may be in the last year.
	r, err := NewRRule(ROption{
		Freq:      Yearly,
		Count:     3,
		Byweekno:  []int{1},
		Byweekday: []Weekday{Monday},
		Dtstart:   testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1997, 12, 29, 9, 0, 0, 0, time.UTC),
		time.Date(1999, 1, 4, 9, 0, 0, 0, time.UTC),
		time.Date(2000, 1, 3, 9, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestYearlyByWeekNoAndWeekDayLarge(t *testing.T) {
	t.Parallel()
	// Another nice test. The last days of week number 52/53
	// may be in the next year.
	r, err := NewRRule(ROption{
		Freq:      Yearly,
		Count:     3,
		Byweekno:  []int{52},
		Byweekday: []Weekday{Sunday},
		Dtstart:   testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1997, 12, 28, 9, 0, 0, 0, time.UTC),
		time.Date(1998, 12, 27, 9, 0, 0, 0, time.UTC),
		time.Date(2000, 1, 2, 9, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestYearlyByWeekNoAndWeekDayLast(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:      Yearly,
		Count:     3,
		Byweekno:  []int{-1},
		Byweekday: []Weekday{Sunday},
		Dtstart:   testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1997, 12, 28, 9, 0, 0, 0, time.UTC),
		time.Date(1999, 1, 3, 9, 0, 0, 0, time.UTC),
		time.Date(2000, 1, 2, 9, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestYearlyByEaster(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:     Yearly,
		Count:    3,
		Byeaster: []int{0},
		Dtstart:  testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1998, 4, 12, 9, 0, 0, 0, time.UTC),
		time.Date(1999, 4, 4, 9, 0, 0, 0, time.UTC),
		time.Date(2000, 4, 23, 9, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestYearlyByEasterPos(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:     Yearly,
		Count:    3,
		Byeaster: []int{1},
		Dtstart:  testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1998, 4, 13, 9, 0, 0, 0, time.UTC),
		time.Date(1999, 4, 5, 9, 0, 0, 0, time.UTC),
		time.Date(2000, 4, 24, 9, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestYearlyByEasterNeg(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:     Yearly,
		Count:    3,
		Byeaster: []int{-1},
		Dtstart:  testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1998, 4, 11, 9, 0, 0, 0, time.UTC),
		time.Date(1999, 4, 3, 9, 0, 0, 0, time.UTC),
		time.Date(2000, 4, 22, 9, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestYearlyByWeekNoAndWeekDay53(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:      Yearly,
		Count:     3,
		Byweekno:  []int{53},
		Byweekday: []Weekday{Monday},
		Dtstart:   testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1998, 12, 28, 9, 0, 0, 0, time.UTC),
		time.Date(2004, 12, 27, 9, 0, 0, 0, time.UTC),
		time.Date(2009, 12, 28, 9, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestYearlyByHour(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:    Yearly,
		Count:   3,
		Byhour:  []int{6, 18},
		Dtstart: testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1997, 9, 2, 18, 0, 0, 0, time.UTC),
		time.Date(1998, 9, 2, 6, 0, 0, 0, time.UTC),
		time.Date(1998, 9, 2, 18, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestYearlyByMinute(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:     Yearly,
		Count:    3,
		Byminute: []int{6, 18},
		Dtstart:  testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1997, 9, 2, 9, 6, 0, 0, time.UTC),
		time.Date(1997, 9, 2, 9, 18, 0, 0, time.UTC),
		time.Date(1998, 9, 2, 9, 6, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestYearlyBySecond(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:     Yearly,
		Count:    3,
		Bysecond: []int{6, 18},
		Dtstart:  testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1997, 9, 2, 9, 0, 6, 0, time.UTC),
		time.Date(1997, 9, 2, 9, 0, 18, 0, time.UTC),
		time.Date(1998, 9, 2, 9, 0, 6, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestYearlyByHourAndMinute(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:     Yearly,
		Count:    3,
		Byhour:   []int{6, 18},
		Byminute: []int{6, 18},
		Dtstart:  testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1997, 9, 2, 18, 6, 0, 0, time.UTC),
		time.Date(1997, 9, 2, 18, 18, 0, 0, time.UTC),
		time.Date(1998, 9, 2, 6, 6, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestYearlyByHourAndSecond(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:     Yearly,
		Count:    3,
		Byhour:   []int{6, 18},
		Bysecond: []int{6, 18},
		Dtstart:  testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1997, 9, 2, 18, 0, 6, 0, time.UTC),
		time.Date(1997, 9, 2, 18, 0, 18, 0, time.UTC),
		time.Date(1998, 9, 2, 6, 0, 6, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestYearlyByMinuteAndSecond(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:     Yearly,
		Count:    3,
		Byminute: []int{6, 18},
		Bysecond: []int{6, 18},
		Dtstart:  testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1997, 9, 2, 9, 6, 6, 0, time.UTC),
		time.Date(1997, 9, 2, 9, 6, 18, 0, time.UTC),
		time.Date(1997, 9, 2, 9, 18, 6, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestYearlyByHourAndMinuteAndSecond(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:     Yearly,
		Count:    3,
		Byhour:   []int{6, 18},
		Byminute: []int{6, 18},
		Bysecond: []int{6, 18},
		Dtstart:  testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1997, 9, 2, 18, 6, 6, 0, time.UTC),
		time.Date(1997, 9, 2, 18, 6, 18, 0, time.UTC),
		time.Date(1997, 9, 2, 18, 18, 6, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestYearlyBySetPos(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:       Yearly,
		Count:      3,
		Bymonthday: []int{15},
		Byhour:     []int{6, 18},
		Bysetpos:   []int{3, -3},
		Dtstart:    testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1997, 11, 15, 18, 0, 0, 0, time.UTC),
		time.Date(1998, 2, 15, 6, 0, 0, 0, time.UTC),
		time.Date(1998, 11, 15, 18, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestMonthly(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:    Monthly,
		Count:   3,
		Dtstart: testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		testDtstart,
		time.Date(1997, 10, 2, 9, 0, 0, 0, time.UTC),
		time.Date(1997, 11, 2, 9, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestMonthlyInterval(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:     Monthly,
		Count:    3,
		Interval: 2,
		Dtstart:  testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		testDtstart,
		time.Date(1997, 11, 2, 9, 0, 0, 0, time.UTC),
		time.Date(1998, 1, 2, 9, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestMonthlyIntervalLarge(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:     Monthly,
		Count:    3,
		Interval: 18,
		Dtstart:  testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		testDtstart,
		time.Date(1999, 3, 2, 9, 0, 0, 0, time.UTC),
		time.Date(2000, 9, 2, 9, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestMonthlyByMonth(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:    Monthly,
		Count:   3,
		Bymonth: []int{1, 3},
		Dtstart: testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1998, 1, 2, 9, 0, 0, 0, time.UTC),
		time.Date(1998, 3, 2, 9, 0, 0, 0, time.UTC),
		time.Date(1999, 1, 2, 9, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestMonthlyByMonthDay(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:       Monthly,
		Count:      3,
		Bymonthday: []int{1, 3},
		Dtstart:    testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1997, 9, 3, 9, 0, 0, 0, time.UTC),
		time.Date(1997, 10, 1, 9, 0, 0, 0, time.UTC),
		time.Date(1997, 10, 3, 9, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestMonthlyByMonthAndMonthDay(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:       Monthly,
		Count:      3,
		Bymonth:    []int{1, 3},
		Bymonthday: []int{5, 7},
		Dtstart:    testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1998, 1, 5, 9, 0, 0, 0, time.UTC),
		time.Date(1998, 1, 7, 9, 0, 0, 0, time.UTC),
		time.Date(1998, 3, 5, 9, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestMonthlyByWeekDay(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:      Monthly,
		Count:     3,
		Byweekday: []Weekday{Tuesday, Thursday},
		Dtstart:   testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		testDtstart,
		time.Date(1997, 9, 4, 9, 0, 0, 0, time.UTC),
		time.Date(1997, 9, 9, 9, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestMonthlyByNWeekDay(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:      Monthly,
		Count:     3,
		Byweekday: []Weekday{Tuesday.Nth(1), Thursday.Nth(-1)},
		Dtstart:   testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		testDtstart,
		time.Date(1997, 9, 25, 9, 0, 0, 0, time.UTC),
		time.Date(1997, 10, 7, 9, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestMonthlyByNWeekDayLarge(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:      Monthly,
		Count:     3,
		Byweekday: []Weekday{Tuesday.Nth(3), Thursday.Nth(-3)},
		Dtstart:   testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1997, 9, 11, 9, 0, 0, 0, time.UTC),
		time.Date(1997, 9, 16, 9, 0, 0, 0, time.UTC),
		time.Date(1997, 10, 16, 9, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestMonthlyByMonthAndWeekDay(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:      Monthly,
		Count:     3,
		Bymonth:   []int{1, 3},
		Byweekday: []Weekday{Tuesday, Thursday},
		Dtstart:   testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1998, 1, 1, 9, 0, 0, 0, time.UTC),
		time.Date(1998, 1, 6, 9, 0, 0, 0, time.UTC),
		time.Date(1998, 1, 8, 9, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestMonthlyByMonthAndNWeekDay(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:      Monthly,
		Count:     3,
		Bymonth:   []int{1, 3},
		Byweekday: []Weekday{Tuesday.Nth(1), Thursday.Nth(-1)},
		Dtstart:   testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1998, 1, 6, 9, 0, 0, 0, time.UTC),
		time.Date(1998, 1, 29, 9, 0, 0, 0, time.UTC),
		time.Date(1998, 3, 3, 9, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestMonthlyByMonthAndNWeekDayLarge(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:      Monthly,
		Count:     3,
		Bymonth:   []int{1, 3},
		Byweekday: []Weekday{Tuesday.Nth(3), Thursday.Nth(-3)},
		Dtstart:   testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1998, 1, 15, 9, 0, 0, 0, time.UTC),
		time.Date(1998, 1, 20, 9, 0, 0, 0, time.UTC),
		time.Date(1998, 3, 12, 9, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestMonthlyByMonthDayAndWeekDay(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:       Monthly,
		Count:      3,
		Bymonthday: []int{1, 3},
		Byweekday:  []Weekday{Tuesday, Thursday},
		Dtstart:    testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1998, 1, 1, 9, 0, 0, 0, time.UTC),
		time.Date(1998, 2, 3, 9, 0, 0, 0, time.UTC),
		time.Date(1998, 3, 3, 9, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestMonthlyByMonthAndMonthDayAndWeekDay(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:       Monthly,
		Count:      3,
		Bymonth:    []int{1, 3},
		Bymonthday: []int{1, 3},
		Byweekday:  []Weekday{Tuesday, Thursday},
		Dtstart:    testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1998, 1, 1, 9, 0, 0, 0, time.UTC),
		time.Date(1998, 3, 3, 9, 0, 0, 0, time.UTC),
		time.Date(2001, 3, 1, 9, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestMonthlyByYearDay(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:      Monthly,
		Count:     4,
		Byyearday: []int{1, 100, 200, 365},
		Dtstart:   testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1997, 12, 31, 9, 0, 0, 0, time.UTC),
		time.Date(1998, 1, 1, 9, 0, 0, 0, time.UTC),
		time.Date(1998, 4, 10, 9, 0, 0, 0, time.UTC),
		time.Date(1998, 7, 19, 9, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestMonthlyByYearDayNeg(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:      Monthly,
		Count:     4,
		Byyearday: []int{-365, -266, -166, -1},
		Dtstart:   testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1997, 12, 31, 9, 0, 0, 0, time.UTC),
		time.Date(1998, 1, 1, 9, 0, 0, 0, time.UTC),
		time.Date(1998, 4, 10, 9, 0, 0, 0, time.UTC),
		time.Date(1998, 7, 19, 9, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestMonthlyByMonthAndYearDay(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:      Monthly,
		Count:     4,
		Bymonth:   []int{4, 7},
		Byyearday: []int{1, 100, 200, 365},
		Dtstart:   testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1998, 4, 10, 9, 0, 0, 0, time.UTC),
		time.Date(1998, 7, 19, 9, 0, 0, 0, time.UTC),
		time.Date(1999, 4, 10, 9, 0, 0, 0, time.UTC),
		time.Date(1999, 7, 19, 9, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestMonthlyByMonthAndYearDayNeg(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:      Monthly,
		Count:     4,
		Bymonth:   []int{4, 7},
		Byyearday: []int{-365, -266, -166, -1},
		Dtstart:   testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1998, 4, 10, 9, 0, 0, 0, time.UTC),
		time.Date(1998, 7, 19, 9, 0, 0, 0, time.UTC),
		time.Date(1999, 4, 10, 9, 0, 0, 0, time.UTC),
		time.Date(1999, 7, 19, 9, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestMonthlyByWeekNo(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:     Monthly,
		Count:    3,
		Byweekno: []int{20},
		Dtstart:  testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1998, 5, 11, 9, 0, 0, 0, time.UTC),
		time.Date(1998, 5, 12, 9, 0, 0, 0, time.UTC),
		time.Date(1998, 5, 13, 9, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestMonthlyByWeekNoAndWeekDay(t *testing.T) {
	t.Parallel()
	// That's a nice one. The first days of week number one
	// may be in the last year.
	r, err := NewRRule(ROption{
		Freq:      Monthly,
		Count:     3,
		Byweekno:  []int{1},
		Byweekday: []Weekday{Monday},
		Dtstart:   testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1997, 12, 29, 9, 0, 0, 0, time.UTC),
		time.Date(1999, 1, 4, 9, 0, 0, 0, time.UTC),
		time.Date(2000, 1, 3, 9, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestMonthlyByWeekNoAndWeekDayLarge(t *testing.T) {
	t.Parallel()
	// Another nice test. The last days of week number 52/53
	// may be in the next year.
	r, err := NewRRule(ROption{
		Freq:      Monthly,
		Count:     3,
		Byweekno:  []int{52},
		Byweekday: []Weekday{Sunday},
		Dtstart:   testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1997, 12, 28, 9, 0, 0, 0, time.UTC),
		time.Date(1998, 12, 27, 9, 0, 0, 0, time.UTC),
		time.Date(2000, 1, 2, 9, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestMonthlyByWeekNoAndWeekDayLast(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:      Monthly,
		Count:     3,
		Byweekno:  []int{-1},
		Byweekday: []Weekday{Sunday},
		Dtstart:   testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1997, 12, 28, 9, 0, 0, 0, time.UTC),
		time.Date(1999, 1, 3, 9, 0, 0, 0, time.UTC),
		time.Date(2000, 1, 2, 9, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestMonthlyByWeekNoAndWeekDay53(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:      Monthly,
		Count:     3,
		Byweekno:  []int{53},
		Byweekday: []Weekday{Monday},
		Dtstart:   testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1998, 12, 28, 9, 0, 0, 0, time.UTC),
		time.Date(2004, 12, 27, 9, 0, 0, 0, time.UTC),
		time.Date(2009, 12, 28, 9, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestMonthlyByEaster(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:     Monthly,
		Count:    3,
		Byeaster: []int{0},
		Dtstart:  testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1998, 4, 12, 9, 0, 0, 0, time.UTC),
		time.Date(1999, 4, 4, 9, 0, 0, 0, time.UTC),
		time.Date(2000, 4, 23, 9, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestMonthlyByEasterPos(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:     Monthly,
		Count:    3,
		Byeaster: []int{1},
		Dtstart:  testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1998, 4, 13, 9, 0, 0, 0, time.UTC),
		time.Date(1999, 4, 5, 9, 0, 0, 0, time.UTC),
		time.Date(2000, 4, 24, 9, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestMonthlyByEasterNeg(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:     Monthly,
		Count:    3,
		Byeaster: []int{-1},
		Dtstart:  testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1998, 4, 11, 9, 0, 0, 0, time.UTC),
		time.Date(1999, 4, 3, 9, 0, 0, 0, time.UTC),
		time.Date(2000, 4, 22, 9, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestMonthlyByHour(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:    Monthly,
		Count:   3,
		Byhour:  []int{6, 18},
		Dtstart: testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1997, 9, 2, 18, 0, 0, 0, time.UTC),
		time.Date(1997, 10, 2, 6, 0, 0, 0, time.UTC),
		time.Date(1997, 10, 2, 18, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestMonthlyByMinute(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:     Monthly,
		Count:    3,
		Byminute: []int{6, 18},
		Dtstart:  testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1997, 9, 2, 9, 6, 0, 0, time.UTC),
		time.Date(1997, 9, 2, 9, 18, 0, 0, time.UTC),
		time.Date(1997, 10, 2, 9, 6, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestMonthlyBySecond(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:     Monthly,
		Count:    3,
		Bysecond: []int{6, 18},
		Dtstart:  testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1997, 9, 2, 9, 0, 6, 0, time.UTC),
		time.Date(1997, 9, 2, 9, 0, 18, 0, time.UTC),
		time.Date(1997, 10, 2, 9, 0, 6, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestMonthlyByHourAndMinute(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:     Monthly,
		Count:    3,
		Byhour:   []int{6, 18},
		Byminute: []int{6, 18},
		Dtstart:  testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1997, 9, 2, 18, 6, 0, 0, time.UTC),
		time.Date(1997, 9, 2, 18, 18, 0, 0, time.UTC),
		time.Date(1997, 10, 2, 6, 6, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestMonthlyByHourAndSecond(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:     Monthly,
		Count:    3,
		Byhour:   []int{6, 18},
		Bysecond: []int{6, 18},
		Dtstart:  testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1997, 9, 2, 18, 0, 6, 0, time.UTC),
		time.Date(1997, 9, 2, 18, 0, 18, 0, time.UTC),
		time.Date(1997, 10, 2, 6, 0, 6, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestMonthlyByMinuteAndSecond(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:     Monthly,
		Count:    3,
		Byminute: []int{6, 18},
		Bysecond: []int{6, 18},
		Dtstart:  testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1997, 9, 2, 9, 6, 6, 0, time.UTC),
		time.Date(1997, 9, 2, 9, 6, 18, 0, time.UTC),
		time.Date(1997, 9, 2, 9, 18, 6, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestMonthlyByHourAndMinuteAndSecond(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:     Monthly,
		Count:    3,
		Byhour:   []int{6, 18},
		Byminute: []int{6, 18},
		Bysecond: []int{6, 18},
		Dtstart:  testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1997, 9, 2, 18, 6, 6, 0, time.UTC),
		time.Date(1997, 9, 2, 18, 6, 18, 0, time.UTC),
		time.Date(1997, 9, 2, 18, 18, 6, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestMonthlyBySetPos(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:       Monthly,
		Count:      3,
		Bymonthday: []int{13, 17},
		Byhour:     []int{6, 18},
		Bysetpos:   []int{3, -3},
		Dtstart:    testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1997, 9, 13, 18, 0, 0, 0, time.UTC),
		time.Date(1997, 9, 17, 6, 0, 0, 0, time.UTC),
		time.Date(1997, 10, 13, 18, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestWeekly(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:    Weekly,
		Count:   3,
		Dtstart: testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		testDtstart,
		time.Date(1997, 9, 9, 9, 0, 0, 0, time.UTC),
		time.Date(1997, 9, 16, 9, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestWeeklyInterval(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:     Weekly,
		Count:    3,
		Interval: 2,
		Dtstart:  testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		testDtstart,
		time.Date(1997, 9, 16, 9, 0, 0, 0, time.UTC),
		time.Date(1997, 9, 30, 9, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestWeeklyIntervalLarge(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:     Weekly,
		Count:    3,
		Interval: 20,
		Dtstart:  testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		testDtstart,
		time.Date(1998, 1, 20, 9, 0, 0, 0, time.UTC),
		time.Date(1998, 6, 9, 9, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestWeeklyByMonth(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:    Weekly,
		Count:   3,
		Bymonth: []int{1, 3},
		Dtstart: testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1998, 1, 6, 9, 0, 0, 0, time.UTC),
		time.Date(1998, 1, 13, 9, 0, 0, 0, time.UTC),
		time.Date(1998, 1, 20, 9, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestWeeklyByMonthDay(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:       Weekly,
		Count:      3,
		Bymonthday: []int{1, 3},
		Dtstart:    testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1997, 9, 3, 9, 0, 0, 0, time.UTC),
		time.Date(1997, 10, 1, 9, 0, 0, 0, time.UTC),
		time.Date(1997, 10, 3, 9, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestWeeklyByMonthAndMonthDay(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:       Weekly,
		Count:      3,
		Bymonth:    []int{1, 3},
		Bymonthday: []int{5, 7},
		Dtstart:    testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1998, 1, 5, 9, 0, 0, 0, time.UTC),
		time.Date(1998, 1, 7, 9, 0, 0, 0, time.UTC),
		time.Date(1998, 3, 5, 9, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestWeeklyByWeekDay(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:      Weekly,
		Count:     3,
		Byweekday: []Weekday{Tuesday, Thursday},
		Dtstart:   testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		testDtstart,
		time.Date(1997, 9, 4, 9, 0, 0, 0, time.UTC),
		time.Date(1997, 9, 9, 9, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestWeeklyByNWeekDay(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:      Weekly,
		Count:     3,
		Byweekday: []Weekday{Tuesday.Nth(1), Thursday.Nth(-1)},
		Dtstart:   testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		testDtstart,
		time.Date(1997, 9, 4, 9, 0, 0, 0, time.UTC),
		time.Date(1997, 9, 9, 9, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestWeeklyByMonthAndWeekDay(t *testing.T) {
	t.Parallel()
	// This test is interesting, because it crosses the year
	// boundary in a weekly period to find day '1' as a
	// valid recurrence.
	r, err := NewRRule(ROption{
		Freq:      Weekly,
		Count:     3,
		Bymonth:   []int{1, 3},
		Byweekday: []Weekday{Tuesday, Thursday},
		Dtstart:   testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1998, 1, 1, 9, 0, 0, 0, time.UTC),
		time.Date(1998, 1, 6, 9, 0, 0, 0, time.UTC),
		time.Date(1998, 1, 8, 9, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestWeeklyByMonthAndNWeekDay(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:      Weekly,
		Count:     3,
		Bymonth:   []int{1, 3},
		Byweekday: []Weekday{Tuesday.Nth(1), Thursday.Nth(-1)},
		Dtstart:   testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1998, 1, 1, 9, 0, 0, 0, time.UTC),
		time.Date(1998, 1, 6, 9, 0, 0, 0, time.UTC),
		time.Date(1998, 1, 8, 9, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestWeeklyByMonthDayAndWeekDay(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:       Weekly,
		Count:      3,
		Bymonthday: []int{1, 3},
		Byweekday:  []Weekday{Tuesday, Thursday},
		Dtstart:    testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1998, 1, 1, 9, 0, 0, 0, time.UTC),
		time.Date(1998, 2, 3, 9, 0, 0, 0, time.UTC),
		time.Date(1998, 3, 3, 9, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestWeeklyByMonthAndMonthDayAndWeekDay(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:       Weekly,
		Count:      3,
		Bymonth:    []int{1, 3},
		Bymonthday: []int{1, 3},
		Byweekday:  []Weekday{Tuesday, Thursday},
		Dtstart:    testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1998, 1, 1, 9, 0, 0, 0, time.UTC),
		time.Date(1998, 3, 3, 9, 0, 0, 0, time.UTC),
		time.Date(2001, 3, 1, 9, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestWeeklyByYearDay(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:      Weekly,
		Count:     4,
		Byyearday: []int{1, 100, 200, 365},
		Dtstart:   testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1997, 12, 31, 9, 0, 0, 0, time.UTC),
		time.Date(1998, 1, 1, 9, 0, 0, 0, time.UTC),
		time.Date(1998, 4, 10, 9, 0, 0, 0, time.UTC),
		time.Date(1998, 7, 19, 9, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestWeeklyByYearDayNeg(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:      Weekly,
		Count:     4,
		Byyearday: []int{-365, -266, -166, -1},
		Dtstart:   testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1997, 12, 31, 9, 0, 0, 0, time.UTC),
		time.Date(1998, 1, 1, 9, 0, 0, 0, time.UTC),
		time.Date(1998, 4, 10, 9, 0, 0, 0, time.UTC),
		time.Date(1998, 7, 19, 9, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestWeeklyByMonthAndYearDay(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:      Weekly,
		Count:     4,
		Bymonth:   []int{1, 7},
		Byyearday: []int{1, 100, 200, 365},
		Dtstart:   testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1998, 1, 1, 9, 0, 0, 0, time.UTC),
		time.Date(1998, 7, 19, 9, 0, 0, 0, time.UTC),
		time.Date(1999, 1, 1, 9, 0, 0, 0, time.UTC),
		time.Date(1999, 7, 19, 9, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestWeeklyByMonthAndYearDayNeg(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:      Weekly,
		Count:     4,
		Bymonth:   []int{1, 7},
		Byyearday: []int{-365, -266, -166, -1},
		Dtstart:   testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1998, 1, 1, 9, 0, 0, 0, time.UTC),
		time.Date(1998, 7, 19, 9, 0, 0, 0, time.UTC),
		time.Date(1999, 1, 1, 9, 0, 0, 0, time.UTC),
		time.Date(1999, 7, 19, 9, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestWeeklyByWeekNo(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:     Weekly,
		Count:    3,
		Byweekno: []int{20},
		Dtstart:  testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1998, 5, 11, 9, 0, 0, 0, time.UTC),
		time.Date(1998, 5, 12, 9, 0, 0, 0, time.UTC),
		time.Date(1998, 5, 13, 9, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestWeeklyByWeekNoAndWeekDay(t *testing.T) {
	t.Parallel()
	// That's a nice one. The first days of week number one
	// may be in the last year.
	r, err := NewRRule(ROption{
		Freq:      Weekly,
		Count:     3,
		Byweekno:  []int{1},
		Byweekday: []Weekday{Monday},
		Dtstart:   testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1997, 12, 29, 9, 0, 0, 0, time.UTC),
		time.Date(1999, 1, 4, 9, 0, 0, 0, time.UTC),
		time.Date(2000, 1, 3, 9, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestWeeklyByWeekNoAndWeekDayLarge(t *testing.T) {
	t.Parallel()
	// Another nice test. The last days of week number 52/53
	// may be in the next year.
	r, err := NewRRule(ROption{
		Freq:      Weekly,
		Count:     3,
		Byweekno:  []int{52},
		Byweekday: []Weekday{Sunday},
		Dtstart:   testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1997, 12, 28, 9, 0, 0, 0, time.UTC),
		time.Date(1998, 12, 27, 9, 0, 0, 0, time.UTC),
		time.Date(2000, 1, 2, 9, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestWeeklyByWeekNoAndWeekDayLast(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:      Weekly,
		Count:     3,
		Byweekno:  []int{-1},
		Byweekday: []Weekday{Sunday},
		Dtstart:   testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1997, 12, 28, 9, 0, 0, 0, time.UTC),
		time.Date(1999, 1, 3, 9, 0, 0, 0, time.UTC),
		time.Date(2000, 1, 2, 9, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestWeeklyByWeekNoAndWeekDay53(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:      Weekly,
		Count:     3,
		Byweekno:  []int{53},
		Byweekday: []Weekday{Monday},
		Dtstart:   testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1998, 12, 28, 9, 0, 0, 0, time.UTC),
		time.Date(2004, 12, 27, 9, 0, 0, 0, time.UTC),
		time.Date(2009, 12, 28, 9, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestWeeklyByEaster(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:     Weekly,
		Count:    3,
		Byeaster: []int{0},
		Dtstart:  testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1998, 4, 12, 9, 0, 0, 0, time.UTC),
		time.Date(1999, 4, 4, 9, 0, 0, 0, time.UTC),
		time.Date(2000, 4, 23, 9, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestWeeklyByEasterPos(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:     Weekly,
		Count:    3,
		Byeaster: []int{1},
		Dtstart:  testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1998, 4, 13, 9, 0, 0, 0, time.UTC),
		time.Date(1999, 4, 5, 9, 0, 0, 0, time.UTC),
		time.Date(2000, 4, 24, 9, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestWeeklyByEasterNeg(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:     Weekly,
		Count:    3,
		Byeaster: []int{-1},
		Dtstart:  testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1998, 4, 11, 9, 0, 0, 0, time.UTC),
		time.Date(1999, 4, 3, 9, 0, 0, 0, time.UTC),
		time.Date(2000, 4, 22, 9, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestWeeklyByHour(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:    Weekly,
		Count:   3,
		Byhour:  []int{6, 18},
		Dtstart: testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1997, 9, 2, 18, 0, 0, 0, time.UTC),
		time.Date(1997, 9, 9, 6, 0, 0, 0, time.UTC),
		time.Date(1997, 9, 9, 18, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestWeeklyByMinute(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:     Weekly,
		Count:    3,
		Byminute: []int{6, 18},
		Dtstart:  testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1997, 9, 2, 9, 6, 0, 0, time.UTC),
		time.Date(1997, 9, 2, 9, 18, 0, 0, time.UTC),
		time.Date(1997, 9, 9, 9, 6, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestWeeklyBySecond(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:     Weekly,
		Count:    3,
		Bysecond: []int{6, 18},
		Dtstart:  testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1997, 9, 2, 9, 0, 6, 0, time.UTC),
		time.Date(1997, 9, 2, 9, 0, 18, 0, time.UTC),
		time.Date(1997, 9, 9, 9, 0, 6, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestWeeklyByHourAndMinute(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:     Weekly,
		Count:    3,
		Byhour:   []int{6, 18},
		Byminute: []int{6, 18},
		Dtstart:  testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1997, 9, 2, 18, 6, 0, 0, time.UTC),
		time.Date(1997, 9, 2, 18, 18, 0, 0, time.UTC),
		time.Date(1997, 9, 9, 6, 6, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestWeeklyByHourAndSecond(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:     Weekly,
		Count:    3,
		Byhour:   []int{6, 18},
		Bysecond: []int{6, 18},
		Dtstart:  testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1997, 9, 2, 18, 0, 6, 0, time.UTC),
		time.Date(1997, 9, 2, 18, 0, 18, 0, time.UTC),
		time.Date(1997, 9, 9, 6, 0, 6, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestWeeklyByMinuteAndSecond(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:     Weekly,
		Count:    3,
		Byminute: []int{6, 18},
		Bysecond: []int{6, 18},
		Dtstart:  testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1997, 9, 2, 9, 6, 6, 0, time.UTC),
		time.Date(1997, 9, 2, 9, 6, 18, 0, time.UTC),
		time.Date(1997, 9, 2, 9, 18, 6, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestWeeklyByHourAndMinuteAndSecond(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:     Weekly,
		Count:    3,
		Byhour:   []int{6, 18},
		Byminute: []int{6, 18},
		Bysecond: []int{6, 18},
		Dtstart:  testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1997, 9, 2, 18, 6, 6, 0, time.UTC),
		time.Date(1997, 9, 2, 18, 6, 18, 0, time.UTC),
		time.Date(1997, 9, 2, 18, 18, 6, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestWeeklyBySetPos(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:      Weekly,
		Count:     3,
		Byweekday: []Weekday{Tuesday, Thursday},
		Byhour:    []int{6, 18},
		Bysetpos:  []int{3, -3},
		Dtstart:   testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1997, 9, 2, 18, 0, 0, 0, time.UTC),
		time.Date(1997, 9, 4, 6, 0, 0, 0, time.UTC),
		time.Date(1997, 9, 9, 18, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestDaily(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:    Daily,
		Count:   3,
		Dtstart: testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		testDtstart,
		time.Date(1997, 9, 3, 9, 0, 0, 0, time.UTC),
		time.Date(1997, 9, 4, 9, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestDailyInterval(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:     Daily,
		Count:    3,
		Interval: 2,
		Dtstart:  testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		testDtstart,
		time.Date(1997, 9, 4, 9, 0, 0, 0, time.UTC),
		time.Date(1997, 9, 6, 9, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestDailyIntervalLarge(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:     Daily,
		Count:    3,
		Interval: 92,
		Dtstart:  testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		testDtstart,
		time.Date(1997, 12, 3, 9, 0, 0, 0, time.UTC),
		time.Date(1998, 3, 5, 9, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestDailyByMonth(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:    Daily,
		Count:   3,
		Bymonth: []int{1, 3},
		Dtstart: testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1998, 1, 1, 9, 0, 0, 0, time.UTC),
		time.Date(1998, 1, 2, 9, 0, 0, 0, time.UTC),
		time.Date(1998, 1, 3, 9, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestDailyByMonthDay(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:       Daily,
		Count:      3,
		Bymonthday: []int{1, 3},
		Dtstart:    testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1997, 9, 3, 9, 0, 0, 0, time.UTC),
		time.Date(1997, 10, 1, 9, 0, 0, 0, time.UTC),
		time.Date(1997, 10, 3, 9, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestDailyByMonthAndMonthDay(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:       Daily,
		Count:      3,
		Bymonth:    []int{1, 3},
		Bymonthday: []int{5, 7},
		Dtstart:    testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1998, 1, 5, 9, 0, 0, 0, time.UTC),
		time.Date(1998, 1, 7, 9, 0, 0, 0, time.UTC),
		time.Date(1998, 3, 5, 9, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestDailyByWeekDay(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:      Daily,
		Count:     3,
		Byweekday: []Weekday{Tuesday, Thursday},
		Dtstart:   testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		testDtstart,
		time.Date(1997, 9, 4, 9, 0, 0, 0, time.UTC),
		time.Date(1997, 9, 9, 9, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestDailyByNWeekDay(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:      Daily,
		Count:     3,
		Byweekday: []Weekday{Tuesday.Nth(1), Thursday.Nth(-1)},
		Dtstart:   testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		testDtstart,
		time.Date(1997, 9, 4, 9, 0, 0, 0, time.UTC),
		time.Date(1997, 9, 9, 9, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestDailyByMonthAndWeekDay(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:      Daily,
		Count:     3,
		Bymonth:   []int{1, 3},
		Byweekday: []Weekday{Tuesday, Thursday},
		Dtstart:   testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1998, 1, 1, 9, 0, 0, 0, time.UTC),
		time.Date(1998, 1, 6, 9, 0, 0, 0, time.UTC),
		time.Date(1998, 1, 8, 9, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestDailyByMonthAndNWeekDay(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:      Daily,
		Count:     3,
		Bymonth:   []int{1, 3},
		Byweekday: []Weekday{Tuesday.Nth(1), Thursday.Nth(-1)},
		Dtstart:   testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1998, 1, 1, 9, 0, 0, 0, time.UTC),
		time.Date(1998, 1, 6, 9, 0, 0, 0, time.UTC),
		time.Date(1998, 1, 8, 9, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestDailyByMonthDayAndWeekDay(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:       Daily,
		Count:      3,
		Bymonthday: []int{1, 3},
		Byweekday:  []Weekday{Tuesday, Thursday},
		Dtstart:    testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1998, 1, 1, 9, 0, 0, 0, time.UTC),
		time.Date(1998, 2, 3, 9, 0, 0, 0, time.UTC),
		time.Date(1998, 3, 3, 9, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestDailyByMonthAndMonthDayAndWeekDay(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:       Daily,
		Count:      3,
		Bymonth:    []int{1, 3},
		Bymonthday: []int{1, 3},
		Byweekday:  []Weekday{Tuesday, Thursday},
		Dtstart:    testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1998, 1, 1, 9, 0, 0, 0, time.UTC),
		time.Date(1998, 3, 3, 9, 0, 0, 0, time.UTC),
		time.Date(2001, 3, 1, 9, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestDailyByYearDay(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:      Daily,
		Count:     4,
		Byyearday: []int{1, 100, 200, 365},
		Dtstart:   testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1997, 12, 31, 9, 0, 0, 0, time.UTC),
		time.Date(1998, 1, 1, 9, 0, 0, 0, time.UTC),
		time.Date(1998, 4, 10, 9, 0, 0, 0, time.UTC),
		time.Date(1998, 7, 19, 9, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestDailyByYearDayNeg(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:      Daily,
		Count:     4,
		Byyearday: []int{-365, -266, -166, -1},
		Dtstart:   testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1997, 12, 31, 9, 0, 0, 0, time.UTC),
		time.Date(1998, 1, 1, 9, 0, 0, 0, time.UTC),
		time.Date(1998, 4, 10, 9, 0, 0, 0, time.UTC),
		time.Date(1998, 7, 19, 9, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestDailyByMonthAndYearDay(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:      Daily,
		Count:     4,
		Bymonth:   []int{1, 7},
		Byyearday: []int{1, 100, 200, 365},
		Dtstart:   testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1998, 1, 1, 9, 0, 0, 0, time.UTC),
		time.Date(1998, 7, 19, 9, 0, 0, 0, time.UTC),
		time.Date(1999, 1, 1, 9, 0, 0, 0, time.UTC),
		time.Date(1999, 7, 19, 9, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestDailyByMonthAndYearDayNeg(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:      Daily,
		Count:     4,
		Bymonth:   []int{1, 7},
		Byyearday: []int{-365, -266, -166, -1},
		Dtstart:   testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1998, 1, 1, 9, 0, 0, 0, time.UTC),
		time.Date(1998, 7, 19, 9, 0, 0, 0, time.UTC),
		time.Date(1999, 1, 1, 9, 0, 0, 0, time.UTC),
		time.Date(1999, 7, 19, 9, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestDailyByWeekNo(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:     Daily,
		Count:    3,
		Byweekno: []int{20},
		Dtstart:  testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1998, 5, 11, 9, 0, 0, 0, time.UTC),
		time.Date(1998, 5, 12, 9, 0, 0, 0, time.UTC),
		time.Date(1998, 5, 13, 9, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestDailyByWeekNoAndWeekDay(t *testing.T) {
	t.Parallel()
	// That's a nice one. The first days of week number one
	// may be in the last year.
	r, err := NewRRule(ROption{
		Freq:      Daily,
		Count:     3,
		Byweekno:  []int{1},
		Byweekday: []Weekday{Monday},
		Dtstart:   testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1997, 12, 29, 9, 0, 0, 0, time.UTC),
		time.Date(1999, 1, 4, 9, 0, 0, 0, time.UTC),
		time.Date(2000, 1, 3, 9, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestDailyByWeekNoAndWeekDayLarge(t *testing.T) {
	t.Parallel()
	// Another nice test. The last days of week number 52/53
	// may be in the next year.
	r, err := NewRRule(ROption{
		Freq:      Daily,
		Count:     3,
		Byweekno:  []int{52},
		Byweekday: []Weekday{Sunday},
		Dtstart:   testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1997, 12, 28, 9, 0, 0, 0, time.UTC),
		time.Date(1998, 12, 27, 9, 0, 0, 0, time.UTC),
		time.Date(2000, 1, 2, 9, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestDailyByWeekNoAndWeekDayLast(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:      Daily,
		Count:     3,
		Byweekno:  []int{-1},
		Byweekday: []Weekday{Sunday},
		Dtstart:   testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1997, 12, 28, 9, 0, 0, 0, time.UTC),
		time.Date(1999, 1, 3, 9, 0, 0, 0, time.UTC),
		time.Date(2000, 1, 2, 9, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestDailyByWeekNoAndWeekDay53(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:      Daily,
		Count:     3,
		Byweekno:  []int{53},
		Byweekday: []Weekday{Monday},
		Dtstart:   testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1998, 12, 28, 9, 0, 0, 0, time.UTC),
		time.Date(2004, 12, 27, 9, 0, 0, 0, time.UTC),
		time.Date(2009, 12, 28, 9, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestDailyByEaster(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:     Daily,
		Count:    3,
		Byeaster: []int{0},
		Dtstart:  testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1998, 4, 12, 9, 0, 0, 0, time.UTC),
		time.Date(1999, 4, 4, 9, 0, 0, 0, time.UTC),
		time.Date(2000, 4, 23, 9, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestDailyByEasterPos(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:     Daily,
		Count:    3,
		Byeaster: []int{1},
		Dtstart:  testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1998, 4, 13, 9, 0, 0, 0, time.UTC),
		time.Date(1999, 4, 5, 9, 0, 0, 0, time.UTC),
		time.Date(2000, 4, 24, 9, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestDailyByEasterNeg(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:     Daily,
		Count:    3,
		Byeaster: []int{-1},
		Dtstart:  testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1998, 4, 11, 9, 0, 0, 0, time.UTC),
		time.Date(1999, 4, 3, 9, 0, 0, 0, time.UTC),
		time.Date(2000, 4, 22, 9, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestDailyByHour(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:    Daily,
		Count:   3,
		Byhour:  []int{6, 18},
		Dtstart: testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1997, 9, 2, 18, 0, 0, 0, time.UTC),
		time.Date(1997, 9, 3, 6, 0, 0, 0, time.UTC),
		time.Date(1997, 9, 3, 18, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestDailyByMinute(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:     Daily,
		Count:    3,
		Byminute: []int{6, 18},
		Dtstart:  testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1997, 9, 2, 9, 6, 0, 0, time.UTC),
		time.Date(1997, 9, 2, 9, 18, 0, 0, time.UTC),
		time.Date(1997, 9, 3, 9, 6, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestDailyBySecond(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:     Daily,
		Count:    3,
		Bysecond: []int{6, 18},
		Dtstart:  testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1997, 9, 2, 9, 0, 6, 0, time.UTC),
		time.Date(1997, 9, 2, 9, 0, 18, 0, time.UTC),
		time.Date(1997, 9, 3, 9, 0, 6, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestDailyByHourAndMinute(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:     Daily,
		Count:    3,
		Byhour:   []int{6, 18},
		Byminute: []int{6, 18},
		Dtstart:  testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1997, 9, 2, 18, 6, 0, 0, time.UTC),
		time.Date(1997, 9, 2, 18, 18, 0, 0, time.UTC),
		time.Date(1997, 9, 3, 6, 6, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestDailyByHourAndSecond(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:     Daily,
		Count:    3,
		Byhour:   []int{6, 18},
		Bysecond: []int{6, 18},
		Dtstart:  testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1997, 9, 2, 18, 0, 6, 0, time.UTC),
		time.Date(1997, 9, 2, 18, 0, 18, 0, time.UTC),
		time.Date(1997, 9, 3, 6, 0, 6, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestDailyByMinuteAndSecond(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:     Daily,
		Count:    3,
		Byminute: []int{6, 18},
		Bysecond: []int{6, 18},
		Dtstart:  testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1997, 9, 2, 9, 6, 6, 0, time.UTC),
		time.Date(1997, 9, 2, 9, 6, 18, 0, time.UTC),
		time.Date(1997, 9, 2, 9, 18, 6, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestDailyByHourAndMinuteAndSecond(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:     Daily,
		Count:    3,
		Byhour:   []int{6, 18},
		Byminute: []int{6, 18},
		Bysecond: []int{6, 18},
		Dtstart:  testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1997, 9, 2, 18, 6, 6, 0, time.UTC),
		time.Date(1997, 9, 2, 18, 6, 18, 0, time.UTC),
		time.Date(1997, 9, 2, 18, 18, 6, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestDailyBySetPos(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:     Daily,
		Count:    3,
		Byhour:   []int{6, 18},
		Byminute: []int{15, 45},
		Bysetpos: []int{3, -3},
		Dtstart:  testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1997, 9, 2, 18, 15, 0, 0, time.UTC),
		time.Date(1997, 9, 3, 6, 45, 0, 0, time.UTC),
		time.Date(1997, 9, 3, 18, 15, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestHourly(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:    Hourly,
		Count:   3,
		Dtstart: testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		testDtstart,
		time.Date(1997, 9, 2, 10, 0, 0, 0, time.UTC),
		time.Date(1997, 9, 2, 11, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestHourlyInterval(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:     Hourly,
		Count:    3,
		Interval: 2,
		Dtstart:  testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		testDtstart,
		time.Date(1997, 9, 2, 11, 0, 0, 0, time.UTC),
		time.Date(1997, 9, 2, 13, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestHourlyIntervalLarge(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:     Hourly,
		Count:    3,
		Interval: 769,
		Dtstart:  testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		testDtstart,
		time.Date(1997, 10, 4, 10, 0, 0, 0, time.UTC),
		time.Date(1997, 11, 5, 11, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestHourlyByMonth(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:    Hourly,
		Count:   3,
		Bymonth: []int{1, 3},
		Dtstart: testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1998, 1, 1, 0, 0, 0, 0, time.UTC),
		time.Date(1998, 1, 1, 1, 0, 0, 0, time.UTC),
		time.Date(1998, 1, 1, 2, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestHourlyByMonthDay(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:       Hourly,
		Count:      3,
		Bymonthday: []int{1, 3},
		Dtstart:    testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1997, 9, 3, 0, 0, 0, 0, time.UTC),
		time.Date(1997, 9, 3, 1, 0, 0, 0, time.UTC),
		time.Date(1997, 9, 3, 2, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestHourlyByMonthAndMonthDay(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:       Hourly,
		Count:      3,
		Bymonth:    []int{1, 3},
		Bymonthday: []int{5, 7},
		Dtstart:    testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1998, 1, 5, 0, 0, 0, 0, time.UTC),
		time.Date(1998, 1, 5, 1, 0, 0, 0, time.UTC),
		time.Date(1998, 1, 5, 2, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestHourlyByWeekDay(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:      Hourly,
		Count:     3,
		Byweekday: []Weekday{Tuesday, Thursday},
		Dtstart:   testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		testDtstart,
		time.Date(1997, 9, 2, 10, 0, 0, 0, time.UTC),
		time.Date(1997, 9, 2, 11, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestHourlyByNWeekDay(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:      Hourly,
		Count:     3,
		Byweekday: []Weekday{Tuesday.Nth(1), Thursday.Nth(-1)},
		Dtstart:   testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		testDtstart,
		time.Date(1997, 9, 2, 10, 0, 0, 0, time.UTC),
		time.Date(1997, 9, 2, 11, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestHourlyByMonthAndWeekDay(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:      Hourly,
		Count:     3,
		Bymonth:   []int{1, 3},
		Byweekday: []Weekday{Tuesday, Thursday},
		Dtstart:   testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1998, 1, 1, 0, 0, 0, 0, time.UTC),
		time.Date(1998, 1, 1, 1, 0, 0, 0, time.UTC),
		time.Date(1998, 1, 1, 2, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestHourlyByMonthAndNWeekDay(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:      Hourly,
		Count:     3,
		Bymonth:   []int{1, 3},
		Byweekday: []Weekday{Tuesday.Nth(1), Thursday.Nth(-1)},
		Dtstart:   testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1998, 1, 1, 0, 0, 0, 0, time.UTC),
		time.Date(1998, 1, 1, 1, 0, 0, 0, time.UTC),
		time.Date(1998, 1, 1, 2, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestHourlyByMonthDayAndWeekDay(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:       Hourly,
		Count:      3,
		Bymonthday: []int{1, 3},
		Byweekday:  []Weekday{Tuesday, Thursday},
		Dtstart:    testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1998, 1, 1, 0, 0, 0, 0, time.UTC),
		time.Date(1998, 1, 1, 1, 0, 0, 0, time.UTC),
		time.Date(1998, 1, 1, 2, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestHourlyByMonthAndMonthDayAndWeekDay(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:       Hourly,
		Count:      3,
		Bymonth:    []int{1, 3},
		Bymonthday: []int{1, 3},
		Byweekday:  []Weekday{Tuesday, Thursday},
		Dtstart:    testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1998, 1, 1, 0, 0, 0, 0, time.UTC),
		time.Date(1998, 1, 1, 1, 0, 0, 0, time.UTC),
		time.Date(1998, 1, 1, 2, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestHourlyByYearDay(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:      Hourly,
		Count:     4,
		Byyearday: []int{1, 100, 200, 365},
		Dtstart:   testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1997, 12, 31, 0, 0, 0, 0, time.UTC),
		time.Date(1997, 12, 31, 1, 0, 0, 0, time.UTC),
		time.Date(1997, 12, 31, 2, 0, 0, 0, time.UTC),
		time.Date(1997, 12, 31, 3, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestHourlyByYearDayNeg(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:      Hourly,
		Count:     4,
		Byyearday: []int{-365, -266, -166, -1},
		Dtstart:   testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1997, 12, 31, 0, 0, 0, 0, time.UTC),
		time.Date(1997, 12, 31, 1, 0, 0, 0, time.UTC),
		time.Date(1997, 12, 31, 2, 0, 0, 0, time.UTC),
		time.Date(1997, 12, 31, 3, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestHourlyByMonthAndYearDay(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:      Hourly,
		Count:     4,
		Bymonth:   []int{4, 7},
		Byyearday: []int{1, 100, 200, 365},
		Dtstart:   testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1998, 4, 10, 0, 0, 0, 0, time.UTC),
		time.Date(1998, 4, 10, 1, 0, 0, 0, time.UTC),
		time.Date(1998, 4, 10, 2, 0, 0, 0, time.UTC),
		time.Date(1998, 4, 10, 3, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestHourlyByMonthAndYearDayNeg(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:      Hourly,
		Count:     4,
		Bymonth:   []int{4, 7},
		Byyearday: []int{-365, -266, -166, -1},
		Dtstart:   testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1998, 4, 10, 0, 0, 0, 0, time.UTC),
		time.Date(1998, 4, 10, 1, 0, 0, 0, time.UTC),
		time.Date(1998, 4, 10, 2, 0, 0, 0, time.UTC),
		time.Date(1998, 4, 10, 3, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestHourlyByWeekNo(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:     Hourly,
		Count:    3,
		Byweekno: []int{20},
		Dtstart:  testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1998, 5, 11, 0, 0, 0, 0, time.UTC),
		time.Date(1998, 5, 11, 1, 0, 0, 0, time.UTC),
		time.Date(1998, 5, 11, 2, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestHourlyByWeekNoAndWeekDay(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:      Hourly,
		Count:     3,
		Byweekno:  []int{1},
		Byweekday: []Weekday{Monday},
		Dtstart:   testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1997, 12, 29, 0, 0, 0, 0, time.UTC),
		time.Date(1997, 12, 29, 1, 0, 0, 0, time.UTC),
		time.Date(1997, 12, 29, 2, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestHourlyByWeekNoAndWeekDayLarge(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:      Hourly,
		Count:     3,
		Byweekno:  []int{52},
		Byweekday: []Weekday{Sunday},
		Dtstart:   testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1997, 12, 28, 0, 0, 0, 0, time.UTC),
		time.Date(1997, 12, 28, 1, 0, 0, 0, time.UTC),
		time.Date(1997, 12, 28, 2, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestHourlyByWeekNoAndWeekDayLast(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:      Hourly,
		Count:     3,
		Byweekno:  []int{-1},
		Byweekday: []Weekday{Sunday},
		Dtstart:   testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1997, 12, 28, 0, 0, 0, 0, time.UTC),
		time.Date(1997, 12, 28, 1, 0, 0, 0, time.UTC),
		time.Date(1997, 12, 28, 2, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestHourlyByWeekNoAndWeekDay53(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:      Hourly,
		Count:     3,
		Byweekno:  []int{53},
		Byweekday: []Weekday{Monday},
		Dtstart:   testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1998, 12, 28, 0, 0, 0, 0, time.UTC),
		time.Date(1998, 12, 28, 1, 0, 0, 0, time.UTC),
		time.Date(1998, 12, 28, 2, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestHourlyByEaster(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:     Hourly,
		Count:    3,
		Byeaster: []int{0},
		Dtstart:  testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1998, 4, 12, 0, 0, 0, 0, time.UTC),
		time.Date(1998, 4, 12, 1, 0, 0, 0, time.UTC),
		time.Date(1998, 4, 12, 2, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestHourlyByEasterPos(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:     Hourly,
		Count:    3,
		Byeaster: []int{1},
		Dtstart:  testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1998, 4, 13, 0, 0, 0, 0, time.UTC),
		time.Date(1998, 4, 13, 1, 0, 0, 0, time.UTC),
		time.Date(1998, 4, 13, 2, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestHourlyByEasterNeg(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:     Hourly,
		Count:    3,
		Byeaster: []int{-1},
		Dtstart:  testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1998, 4, 11, 0, 0, 0, 0, time.UTC),
		time.Date(1998, 4, 11, 1, 0, 0, 0, time.UTC),
		time.Date(1998, 4, 11, 2, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestHourlyByHour(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:    Hourly,
		Count:   3,
		Byhour:  []int{6, 18},
		Dtstart: testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1997, 9, 2, 18, 0, 0, 0, time.UTC),
		time.Date(1997, 9, 3, 6, 0, 0, 0, time.UTC),
		time.Date(1997, 9, 3, 18, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestHourlyByMinute(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:     Hourly,
		Count:    3,
		Byminute: []int{6, 18},
		Dtstart:  testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1997, 9, 2, 9, 6, 0, 0, time.UTC),
		time.Date(1997, 9, 2, 9, 18, 0, 0, time.UTC),
		time.Date(1997, 9, 2, 10, 6, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestHourlyBySecond(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:     Hourly,
		Count:    3,
		Bysecond: []int{6, 18},
		Dtstart:  testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1997, 9, 2, 9, 0, 6, 0, time.UTC),
		time.Date(1997, 9, 2, 9, 0, 18, 0, time.UTC),
		time.Date(1997, 9, 2, 10, 0, 6, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestHourlyByHourAndMinute(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:     Hourly,
		Count:    3,
		Byhour:   []int{6, 18},
		Byminute: []int{6, 18},
		Dtstart:  testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1997, 9, 2, 18, 6, 0, 0, time.UTC),
		time.Date(1997, 9, 2, 18, 18, 0, 0, time.UTC),
		time.Date(1997, 9, 3, 6, 6, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestHourlyByHourAndSecond(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:     Hourly,
		Count:    3,
		Byhour:   []int{6, 18},
		Bysecond: []int{6, 18},
		Dtstart:  testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1997, 9, 2, 18, 0, 6, 0, time.UTC),
		time.Date(1997, 9, 2, 18, 0, 18, 0, time.UTC),
		time.Date(1997, 9, 3, 6, 0, 6, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestHourlyByMinuteAndSecond(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:     Hourly,
		Count:    3,
		Byminute: []int{6, 18},
		Bysecond: []int{6, 18},
		Dtstart:  testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1997, 9, 2, 9, 6, 6, 0, time.UTC),
		time.Date(1997, 9, 2, 9, 6, 18, 0, time.UTC),
		time.Date(1997, 9, 2, 9, 18, 6, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestHourlyByHourAndMinuteAndSecond(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:     Hourly,
		Count:    3,
		Byhour:   []int{6, 18},
		Byminute: []int{6, 18},
		Bysecond: []int{6, 18},
		Dtstart:  testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1997, 9, 2, 18, 6, 6, 0, time.UTC),
		time.Date(1997, 9, 2, 18, 6, 18, 0, time.UTC),
		time.Date(1997, 9, 2, 18, 18, 6, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestHourlyBySetPos(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:     Hourly,
		Count:    3,
		Byminute: []int{15, 45},
		Bysecond: []int{15, 45},
		Bysetpos: []int{3, -3},
		Dtstart:  testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1997, 9, 2, 9, 15, 45, 0, time.UTC),
		time.Date(1997, 9, 2, 9, 45, 15, 0, time.UTC),
		time.Date(1997, 9, 2, 10, 15, 45, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestMinutely(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:    Minutely,
		Count:   3,
		Dtstart: testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		testDtstart,
		time.Date(1997, 9, 2, 9, 1, 0, 0, time.UTC),
		time.Date(1997, 9, 2, 9, 2, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestMinutelyInterval(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:     Minutely,
		Count:    3,
		Interval: 2,
		Dtstart:  testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		testDtstart,
		time.Date(1997, 9, 2, 9, 2, 0, 0, time.UTC),
		time.Date(1997, 9, 2, 9, 4, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestMinutelyIntervalLarge(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:     Minutely,
		Count:    3,
		Interval: 1501,
		Dtstart:  testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		testDtstart,
		time.Date(1997, 9, 3, 10, 1, 0, 0, time.UTC),
		time.Date(1997, 9, 4, 11, 2, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestMinutelyByMonth(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:    Minutely,
		Count:   3,
		Bymonth: []int{1, 3},
		Dtstart: testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1998, 1, 1, 0, 0, 0, 0, time.UTC),
		time.Date(1998, 1, 1, 0, 1, 0, 0, time.UTC),
		time.Date(1998, 1, 1, 0, 2, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestMinutelyByMonthDay(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:       Minutely,
		Count:      3,
		Bymonthday: []int{1, 3},
		Dtstart:    testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1997, 9, 3, 0, 0, 0, 0, time.UTC),
		time.Date(1997, 9, 3, 0, 1, 0, 0, time.UTC),
		time.Date(1997, 9, 3, 0, 2, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestMinutelyByMonthAndMonthDay(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:       Minutely,
		Count:      3,
		Bymonth:    []int{1, 3},
		Bymonthday: []int{5, 7},
		Dtstart:    testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1998, 1, 5, 0, 0, 0, 0, time.UTC),
		time.Date(1998, 1, 5, 0, 1, 0, 0, time.UTC),
		time.Date(1998, 1, 5, 0, 2, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestMinutelyByWeekDay(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:      Minutely,
		Count:     3,
		Byweekday: []Weekday{Tuesday, Thursday},
		Dtstart:   testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		testDtstart,
		time.Date(1997, 9, 2, 9, 1, 0, 0, time.UTC),
		time.Date(1997, 9, 2, 9, 2, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestMinutelyByNWeekDay(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:      Minutely,
		Count:     3,
		Byweekday: []Weekday{Tuesday.Nth(1), Thursday.Nth(-1)},
		Dtstart:   testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		testDtstart,
		time.Date(1997, 9, 2, 9, 1, 0, 0, time.UTC),
		time.Date(1997, 9, 2, 9, 2, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestMinutelyByMonthAndWeekDay(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:      Minutely,
		Count:     3,
		Bymonth:   []int{1, 3},
		Byweekday: []Weekday{Tuesday, Thursday},
		Dtstart:   testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1998, 1, 1, 0, 0, 0, 0, time.UTC),
		time.Date(1998, 1, 1, 0, 1, 0, 0, time.UTC),
		time.Date(1998, 1, 1, 0, 2, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestMinutelyByMonthAndNWeekDay(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:      Minutely,
		Count:     3,
		Bymonth:   []int{1, 3},
		Byweekday: []Weekday{Tuesday.Nth(1), Thursday.Nth(-1)},
		Dtstart:   testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1998, 1, 1, 0, 0, 0, 0, time.UTC),
		time.Date(1998, 1, 1, 0, 1, 0, 0, time.UTC),
		time.Date(1998, 1, 1, 0, 2, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestMinutelyByMonthDayAndWeekDay(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:       Minutely,
		Count:      3,
		Bymonthday: []int{1, 3},
		Byweekday:  []Weekday{Tuesday, Thursday},
		Dtstart:    testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1998, 1, 1, 0, 0, 0, 0, time.UTC),
		time.Date(1998, 1, 1, 0, 1, 0, 0, time.UTC),
		time.Date(1998, 1, 1, 0, 2, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestMinutelyByMonthAndMonthDayAndWeekDay(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:       Minutely,
		Count:      3,
		Bymonth:    []int{1, 3},
		Bymonthday: []int{1, 3},
		Byweekday:  []Weekday{Tuesday, Thursday},
		Dtstart:    testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1998, 1, 1, 0, 0, 0, 0, time.UTC),
		time.Date(1998, 1, 1, 0, 1, 0, 0, time.UTC),
		time.Date(1998, 1, 1, 0, 2, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestMinutelyByYearDay(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:      Minutely,
		Count:     4,
		Byyearday: []int{1, 100, 200, 365},
		Dtstart:   testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1997, 12, 31, 0, 0, 0, 0, time.UTC),
		time.Date(1997, 12, 31, 0, 1, 0, 0, time.UTC),
		time.Date(1997, 12, 31, 0, 2, 0, 0, time.UTC),
		time.Date(1997, 12, 31, 0, 3, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestMinutelyByYearDayNeg(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:      Minutely,
		Count:     4,
		Byyearday: []int{-365, -266, -166, -1},
		Dtstart:   testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1997, 12, 31, 0, 0, 0, 0, time.UTC),
		time.Date(1997, 12, 31, 0, 1, 0, 0, time.UTC),
		time.Date(1997, 12, 31, 0, 2, 0, 0, time.UTC),
		time.Date(1997, 12, 31, 0, 3, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestMinutelyByMonthAndYearDay(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:      Minutely,
		Count:     4,
		Bymonth:   []int{4, 7},
		Byyearday: []int{1, 100, 200, 365},
		Dtstart:   testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1998, 4, 10, 0, 0, 0, 0, time.UTC),
		time.Date(1998, 4, 10, 0, 1, 0, 0, time.UTC),
		time.Date(1998, 4, 10, 0, 2, 0, 0, time.UTC),
		time.Date(1998, 4, 10, 0, 3, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestMinutelyByMonthAndYearDayNeg(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:      Minutely,
		Count:     4,
		Bymonth:   []int{4, 7},
		Byyearday: []int{-365, -266, -166, -1},
		Dtstart:   testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1998, 4, 10, 0, 0, 0, 0, time.UTC),
		time.Date(1998, 4, 10, 0, 1, 0, 0, time.UTC),
		time.Date(1998, 4, 10, 0, 2, 0, 0, time.UTC),
		time.Date(1998, 4, 10, 0, 3, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestMinutelyByWeekNo(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:     Minutely,
		Count:    3,
		Byweekno: []int{20},
		Dtstart:  testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1998, 5, 11, 0, 0, 0, 0, time.UTC),
		time.Date(1998, 5, 11, 0, 1, 0, 0, time.UTC),
		time.Date(1998, 5, 11, 0, 2, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestMinutelyByWeekNoAndWeekDay(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:      Minutely,
		Count:     3,
		Byweekno:  []int{1},
		Byweekday: []Weekday{Monday},
		Dtstart:   testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1997, 12, 29, 0, 0, 0, 0, time.UTC),
		time.Date(1997, 12, 29, 0, 1, 0, 0, time.UTC),
		time.Date(1997, 12, 29, 0, 2, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestMinutelyByWeekNoAndWeekDayLarge(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:      Minutely,
		Count:     3,
		Byweekno:  []int{52},
		Byweekday: []Weekday{Sunday},
		Dtstart:   testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1997, 12, 28, 0, 0, 0, 0, time.UTC),
		time.Date(1997, 12, 28, 0, 1, 0, 0, time.UTC),
		time.Date(1997, 12, 28, 0, 2, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestMinutelyByWeekNoAndWeekDayLast(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:      Minutely,
		Count:     3,
		Byweekno:  []int{-1},
		Byweekday: []Weekday{Sunday},
		Dtstart:   testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1997, 12, 28, 0, 0, 0, 0, time.UTC),
		time.Date(1997, 12, 28, 0, 1, 0, 0, time.UTC),
		time.Date(1997, 12, 28, 0, 2, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestMinutelyByWeekNoAndWeekDay53(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:      Minutely,
		Count:     3,
		Byweekno:  []int{53},
		Byweekday: []Weekday{Monday},
		Dtstart:   testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1998, 12, 28, 0, 0, 0, 0, time.UTC),
		time.Date(1998, 12, 28, 0, 1, 0, 0, time.UTC),
		time.Date(1998, 12, 28, 0, 2, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestMinutelyByEaster(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:     Minutely,
		Count:    3,
		Byeaster: []int{0},
		Dtstart:  testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1998, 4, 12, 0, 0, 0, 0, time.UTC),
		time.Date(1998, 4, 12, 0, 1, 0, 0, time.UTC),
		time.Date(1998, 4, 12, 0, 2, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestMinutelyByEasterPos(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:     Minutely,
		Count:    3,
		Byeaster: []int{1},
		Dtstart:  testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1998, 4, 13, 0, 0, 0, 0, time.UTC),
		time.Date(1998, 4, 13, 0, 1, 0, 0, time.UTC),
		time.Date(1998, 4, 13, 0, 2, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestMinutelyByEasterNeg(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:     Minutely,
		Count:    3,
		Byeaster: []int{-1},
		Dtstart:  testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1998, 4, 11, 0, 0, 0, 0, time.UTC),
		time.Date(1998, 4, 11, 0, 1, 0, 0, time.UTC),
		time.Date(1998, 4, 11, 0, 2, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestMinutelyByHour(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:    Minutely,
		Count:   3,
		Byhour:  []int{6, 18},
		Dtstart: testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1997, 9, 2, 18, 0, 0, 0, time.UTC),
		time.Date(1997, 9, 2, 18, 1, 0, 0, time.UTC),
		time.Date(1997, 9, 2, 18, 2, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestMinutelyByMinute(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:     Minutely,
		Count:    3,
		Byminute: []int{6, 18},
		Dtstart:  testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1997, 9, 2, 9, 6, 0, 0, time.UTC),
		time.Date(1997, 9, 2, 9, 18, 0, 0, time.UTC),
		time.Date(1997, 9, 2, 10, 6, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestMinutelyBySecond(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:     Minutely,
		Count:    3,
		Bysecond: []int{6, 18},
		Dtstart:  testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1997, 9, 2, 9, 0, 6, 0, time.UTC),
		time.Date(1997, 9, 2, 9, 0, 18, 0, time.UTC),
		time.Date(1997, 9, 2, 9, 1, 6, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestMinutelyByHourAndMinute(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:     Minutely,
		Count:    3,
		Byhour:   []int{6, 18},
		Byminute: []int{6, 18},
		Dtstart:  testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1997, 9, 2, 18, 6, 0, 0, time.UTC),
		time.Date(1997, 9, 2, 18, 18, 0, 0, time.UTC),
		time.Date(1997, 9, 3, 6, 6, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestMinutelyByHourAndSecond(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:     Minutely,
		Count:    3,
		Byhour:   []int{6, 18},
		Bysecond: []int{6, 18},
		Dtstart:  testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1997, 9, 2, 18, 0, 6, 0, time.UTC),
		time.Date(1997, 9, 2, 18, 0, 18, 0, time.UTC),
		time.Date(1997, 9, 2, 18, 1, 6, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestMinutelyByMinuteAndSecond(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:     Minutely,
		Count:    3,
		Byminute: []int{6, 18},
		Bysecond: []int{6, 18},
		Dtstart:  testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1997, 9, 2, 9, 6, 6, 0, time.UTC),
		time.Date(1997, 9, 2, 9, 6, 18, 0, time.UTC),
		time.Date(1997, 9, 2, 9, 18, 6, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestMinutelyByHourAndMinuteAndSecond(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:     Minutely,
		Count:    3,
		Byhour:   []int{6, 18},
		Byminute: []int{6, 18},
		Bysecond: []int{6, 18},
		Dtstart:  testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1997, 9, 2, 18, 6, 6, 0, time.UTC),
		time.Date(1997, 9, 2, 18, 6, 18, 0, time.UTC),
		time.Date(1997, 9, 2, 18, 18, 6, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestMinutelyBySetPos(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:     Minutely,
		Count:    3,
		Bysecond: []int{15, 30, 45},
		Bysetpos: []int{3, -3},
		Dtstart:  testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1997, 9, 2, 9, 0, 15, 0, time.UTC),
		time.Date(1997, 9, 2, 9, 0, 45, 0, time.UTC),
		time.Date(1997, 9, 2, 9, 1, 15, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestSecondly(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:    Secondly,
		Count:   3,
		Dtstart: testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		testDtstart,
		time.Date(1997, 9, 2, 9, 0, 1, 0, time.UTC),
		time.Date(1997, 9, 2, 9, 0, 2, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestSecondlyInterval(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:     Secondly,
		Count:    3,
		Interval: 2,
		Dtstart:  testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		testDtstart,
		time.Date(1997, 9, 2, 9, 0, 2, 0, time.UTC),
		time.Date(1997, 9, 2, 9, 0, 4, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestSecondlyIntervalLarge(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:     Secondly,
		Count:    3,
		Interval: 90061,
		Dtstart:  testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		testDtstart,
		time.Date(1997, 9, 3, 10, 1, 1, 0, time.UTC),
		time.Date(1997, 9, 4, 11, 2, 2, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestSecondlyByMonth(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:    Secondly,
		Count:   3,
		Bymonth: []int{1, 3},
		Dtstart: testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1998, 1, 1, 0, 0, 0, 0, time.UTC),
		time.Date(1998, 1, 1, 0, 0, 1, 0, time.UTC),
		time.Date(1998, 1, 1, 0, 0, 2, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestSecondlyByMonthDay(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:       Secondly,
		Count:      3,
		Bymonthday: []int{1, 3},
		Dtstart:    testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1997, 9, 3, 0, 0, 0, 0, time.UTC),
		time.Date(1997, 9, 3, 0, 0, 1, 0, time.UTC),
		time.Date(1997, 9, 3, 0, 0, 2, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestSecondlyByMonthAndMonthDay(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:       Secondly,
		Count:      3,
		Bymonth:    []int{1, 3},
		Bymonthday: []int{5, 7},
		Dtstart:    testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1998, 1, 5, 0, 0, 0, 0, time.UTC),
		time.Date(1998, 1, 5, 0, 0, 1, 0, time.UTC),
		time.Date(1998, 1, 5, 0, 0, 2, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestSecondlyByWeekDay(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:      Secondly,
		Count:     3,
		Byweekday: []Weekday{Tuesday, Thursday},
		Dtstart:   testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		testDtstart,
		time.Date(1997, 9, 2, 9, 0, 1, 0, time.UTC),
		time.Date(1997, 9, 2, 9, 0, 2, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestSecondlyByNWeekDay(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:      Secondly,
		Count:     3,
		Byweekday: []Weekday{Tuesday.Nth(1), Thursday.Nth(-1)},
		Dtstart:   testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		testDtstart,
		time.Date(1997, 9, 2, 9, 0, 1, 0, time.UTC),
		time.Date(1997, 9, 2, 9, 0, 2, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestSecondlyByMonthAndWeekDay(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:      Secondly,
		Count:     3,
		Bymonth:   []int{1, 3},
		Byweekday: []Weekday{Tuesday, Thursday},
		Dtstart:   testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1998, 1, 1, 0, 0, 0, 0, time.UTC),
		time.Date(1998, 1, 1, 0, 0, 1, 0, time.UTC),
		time.Date(1998, 1, 1, 0, 0, 2, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestSecondlyByMonthAndNWeekDay(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:      Secondly,
		Count:     3,
		Bymonth:   []int{1, 3},
		Byweekday: []Weekday{Tuesday.Nth(1), Thursday.Nth(-1)},
		Dtstart:   testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1998, 1, 1, 0, 0, 0, 0, time.UTC),
		time.Date(1998, 1, 1, 0, 0, 1, 0, time.UTC),
		time.Date(1998, 1, 1, 0, 0, 2, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestSecondlyByMonthDayAndWeekDay(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:       Secondly,
		Count:      3,
		Bymonthday: []int{1, 3},
		Byweekday:  []Weekday{Tuesday, Thursday},
		Dtstart:    testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1998, 1, 1, 0, 0, 0, 0, time.UTC),
		time.Date(1998, 1, 1, 0, 0, 1, 0, time.UTC),
		time.Date(1998, 1, 1, 0, 0, 2, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestSecondlyByMonthAndMonthDayAndWeekDay(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:       Secondly,
		Count:      3,
		Bymonth:    []int{1, 3},
		Bymonthday: []int{1, 3},
		Byweekday:  []Weekday{Tuesday, Thursday},
		Dtstart:    testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1998, 1, 1, 0, 0, 0, 0, time.UTC),
		time.Date(1998, 1, 1, 0, 0, 1, 0, time.UTC),
		time.Date(1998, 1, 1, 0, 0, 2, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestSecondlyByYearDay(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:      Secondly,
		Count:     4,
		Byyearday: []int{1, 100, 200, 365},
		Dtstart:   testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1997, 12, 31, 0, 0, 0, 0, time.UTC),
		time.Date(1997, 12, 31, 0, 0, 1, 0, time.UTC),
		time.Date(1997, 12, 31, 0, 0, 2, 0, time.UTC),
		time.Date(1997, 12, 31, 0, 0, 3, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestSecondlyByYearDayNeg(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:      Secondly,
		Count:     4,
		Byyearday: []int{-365, -266, -166, -1},
		Dtstart:   testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1997, 12, 31, 0, 0, 0, 0, time.UTC),
		time.Date(1997, 12, 31, 0, 0, 1, 0, time.UTC),
		time.Date(1997, 12, 31, 0, 0, 2, 0, time.UTC),
		time.Date(1997, 12, 31, 0, 0, 3, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestSecondlyByMonthAndYearDay(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:      Secondly,
		Count:     4,
		Bymonth:   []int{4, 7},
		Byyearday: []int{1, 100, 200, 365},
		Dtstart:   testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1998, 4, 10, 0, 0, 0, 0, time.UTC),
		time.Date(1998, 4, 10, 0, 0, 1, 0, time.UTC),
		time.Date(1998, 4, 10, 0, 0, 2, 0, time.UTC),
		time.Date(1998, 4, 10, 0, 0, 3, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestSecondlyByMonthAndYearDayNeg(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:      Secondly,
		Count:     4,
		Bymonth:   []int{4, 7},
		Byyearday: []int{-365, -266, -166, -1},
		Dtstart:   testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1998, 4, 10, 0, 0, 0, 0, time.UTC),
		time.Date(1998, 4, 10, 0, 0, 1, 0, time.UTC),
		time.Date(1998, 4, 10, 0, 0, 2, 0, time.UTC),
		time.Date(1998, 4, 10, 0, 0, 3, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestSecondlyByWeekNo(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:     Secondly,
		Count:    3,
		Byweekno: []int{20},
		Dtstart:  testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1998, 5, 11, 0, 0, 0, 0, time.UTC),
		time.Date(1998, 5, 11, 0, 0, 1, 0, time.UTC),
		time.Date(1998, 5, 11, 0, 0, 2, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestSecondlyByWeekNoAndWeekDay(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:      Secondly,
		Count:     3,
		Byweekno:  []int{1},
		Byweekday: []Weekday{Monday},
		Dtstart:   testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1997, 12, 29, 0, 0, 0, 0, time.UTC),
		time.Date(1997, 12, 29, 0, 0, 1, 0, time.UTC),
		time.Date(1997, 12, 29, 0, 0, 2, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestSecondlyByWeekNoAndWeekDayLarge(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:      Secondly,
		Count:     3,
		Byweekno:  []int{52},
		Byweekday: []Weekday{Sunday},
		Dtstart:   testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1997, 12, 28, 0, 0, 0, 0, time.UTC),
		time.Date(1997, 12, 28, 0, 0, 1, 0, time.UTC),
		time.Date(1997, 12, 28, 0, 0, 2, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestSecondlyByWeekNoAndWeekDayLast(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:      Secondly,
		Count:     3,
		Byweekno:  []int{-1},
		Byweekday: []Weekday{Sunday},
		Dtstart:   testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1997, 12, 28, 0, 0, 0, 0, time.UTC),
		time.Date(1997, 12, 28, 0, 0, 1, 0, time.UTC),
		time.Date(1997, 12, 28, 0, 0, 2, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestSecondlyByWeekNoAndWeekDay53(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:      Secondly,
		Count:     3,
		Byweekno:  []int{53},
		Byweekday: []Weekday{Monday},
		Dtstart:   testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1998, 12, 28, 0, 0, 0, 0, time.UTC),
		time.Date(1998, 12, 28, 0, 0, 1, 0, time.UTC),
		time.Date(1998, 12, 28, 0, 0, 2, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestSecondlyByEaster(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:     Secondly,
		Count:    3,
		Byeaster: []int{0},
		Dtstart:  testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1998, 4, 12, 0, 0, 0, 0, time.UTC),
		time.Date(1998, 4, 12, 0, 0, 1, 0, time.UTC),
		time.Date(1998, 4, 12, 0, 0, 2, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestSecondlyByEasterPos(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:     Secondly,
		Count:    3,
		Byeaster: []int{1},
		Dtstart:  testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1998, 4, 13, 0, 0, 0, 0, time.UTC),
		time.Date(1998, 4, 13, 0, 0, 1, 0, time.UTC),
		time.Date(1998, 4, 13, 0, 0, 2, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestSecondlyByEasterNeg(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:     Secondly,
		Count:    3,
		Byeaster: []int{-1},
		Dtstart:  testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1998, 4, 11, 0, 0, 0, 0, time.UTC),
		time.Date(1998, 4, 11, 0, 0, 1, 0, time.UTC),
		time.Date(1998, 4, 11, 0, 0, 2, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestSecondlyByHour(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:    Secondly,
		Count:   3,
		Byhour:  []int{6, 18},
		Dtstart: testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1997, 9, 2, 18, 0, 0, 0, time.UTC),
		time.Date(1997, 9, 2, 18, 0, 1, 0, time.UTC),
		time.Date(1997, 9, 2, 18, 0, 2, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestSecondlyByMinute(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:     Secondly,
		Count:    3,
		Byminute: []int{6, 18},
		Dtstart:  testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1997, 9, 2, 9, 6, 0, 0, time.UTC),
		time.Date(1997, 9, 2, 9, 6, 1, 0, time.UTC),
		time.Date(1997, 9, 2, 9, 6, 2, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestSecondlyBySecond(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:     Secondly,
		Count:    3,
		Bysecond: []int{6, 18},
		Dtstart:  testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1997, 9, 2, 9, 0, 6, 0, time.UTC),
		time.Date(1997, 9, 2, 9, 0, 18, 0, time.UTC),
		time.Date(1997, 9, 2, 9, 1, 6, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestSecondlyByHourAndMinute(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:     Secondly,
		Count:    3,
		Byhour:   []int{6, 18},
		Byminute: []int{6, 18},
		Dtstart:  testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1997, 9, 2, 18, 6, 0, 0, time.UTC),
		time.Date(1997, 9, 2, 18, 6, 1, 0, time.UTC),
		time.Date(1997, 9, 2, 18, 6, 2, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestSecondlyByHourAndSecond(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:     Secondly,
		Count:    3,
		Byhour:   []int{6, 18},
		Bysecond: []int{6, 18},
		Dtstart:  testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1997, 9, 2, 18, 0, 6, 0, time.UTC),
		time.Date(1997, 9, 2, 18, 0, 18, 0, time.UTC),
		time.Date(1997, 9, 2, 18, 1, 6, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestSecondlyByMinuteAndSecond(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:     Secondly,
		Count:    3,
		Byminute: []int{6, 18},
		Bysecond: []int{6, 18},
		Dtstart:  testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1997, 9, 2, 9, 6, 6, 0, time.UTC),
		time.Date(1997, 9, 2, 9, 6, 18, 0, time.UTC),
		time.Date(1997, 9, 2, 9, 18, 6, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestSecondlyByHourAndMinuteAndSecond(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:     Secondly,
		Count:    3,
		Byhour:   []int{6, 18},
		Byminute: []int{6, 18},
		Bysecond: []int{6, 18},
		Dtstart:  testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1997, 9, 2, 18, 6, 6, 0, time.UTC),
		time.Date(1997, 9, 2, 18, 6, 18, 0, time.UTC),
		time.Date(1997, 9, 2, 18, 18, 6, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestSecondlyByHourAndMinuteAndSecondBug(t *testing.T) {
	t.Parallel()
	// This explores a bug found by Mathieu Bridon.
	r, err := NewRRule(ROption{
		Freq:     Secondly,
		Count:    3,
		Bysecond: []int{0},
		Byminute: []int{1},
		Dtstart:  time.Date(2010, 3, 22, 12, 1, 0, 0, time.UTC),
	},
	)
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(2010, 3, 22, 12, 1, 0, 0, time.UTC),
		time.Date(2010, 3, 22, 13, 1, 0, 0, time.UTC),
		time.Date(2010, 3, 22, 14, 1, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestUntilNotMatching(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:    Daily,
		Count:   3,
		Dtstart: testDtstart,
		Until:   time.Date(1997, 9, 5, 8, 0, 0, 0, time.UTC),
	},
	)
	assert.NoError(t, err)
	want := []time.Time{
		testDtstart,
		time.Date(1997, 9, 3, 9, 0, 0, 0, time.UTC),
		time.Date(1997, 9, 4, 9, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestUntilMatching(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:    Daily,
		Count:   3,
		Dtstart: testDtstart,
		Until:   time.Date(1997, 9, 4, 9, 0, 0, 0, time.UTC),
	},
	)
	assert.NoError(t, err)
	want := []time.Time{
		testDtstart,
		time.Date(1997, 9, 3, 9, 0, 0, 0, time.UTC),
		time.Date(1997, 9, 4, 9, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestUntilSingle(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:    Daily,
		Count:   3,
		Dtstart: testDtstart,
		Until:   testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{testDtstart}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestUntilEmpty(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:    Daily,
		Count:   3,
		Dtstart: testDtstart,
		Until:   time.Date(1997, 9, 1, 9, 0, 0, 0, time.UTC),
	},
	)
	assert.NoError(t, err)
	want := []time.Time{}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestUntilWithDate(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:    Daily,
		Count:   3,
		Dtstart: testDtstart,
		Until:   time.Date(1997, 9, 5, 0, 0, 0, 0, time.UTC),
	},
	)
	assert.NoError(t, err)
	want := []time.Time{
		testDtstart,
		time.Date(1997, 9, 3, 9, 0, 0, 0, time.UTC),
		time.Date(1997, 9, 4, 9, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestWkStIntervalMO(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:      Weekly,
		Count:     3,
		Interval:  2,
		Byweekday: []Weekday{Tuesday, Sunday},
		Wkst:      Monday,
		Dtstart:   testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		testDtstart,
		time.Date(1997, 9, 7, 9, 0, 0, 0, time.UTC),
		time.Date(1997, 9, 16, 9, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestWkStIntervalSU(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:      Weekly,
		Count:     3,
		Interval:  2,
		Byweekday: []Weekday{Tuesday, Sunday},
		Wkst:      Sunday,
		Dtstart:   testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		testDtstart,
		time.Date(1997, 9, 14, 9, 0, 0, 0, time.UTC),
		time.Date(1997, 9, 16, 9, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestDTStart(t *testing.T) {
	t.Parallel()
	dt := time.Now().UTC().Truncate(time.Second)
	r, err := NewRRule(ROption{Freq: Yearly, Count: 3})
	assert.NoError(t, err)
	want := []time.Time{dt, dt.AddDate(1, 0, 0), dt.AddDate(2, 0, 0)}
	value := r.All()
	assert.Equal(t, want, value)

	dt = dt.AddDate(0, 0, 3)
	r.DTStart(dt)
	want = []time.Time{dt, dt.AddDate(1, 0, 0), dt.AddDate(2, 0, 0)}
	value = r.All()
	assert.Equal(t, want, value)
}

func TestDTStartIsDate(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:    Daily,
		Count:   3,
		Dtstart: time.Date(1997, 9, 2, 0, 0, 0, 0, time.UTC),
	},
	)
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1997, 9, 2, 0, 0, 0, 0, time.UTC),
		time.Date(1997, 9, 3, 0, 0, 0, 0, time.UTC),
		time.Date(1997, 9, 4, 0, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestDTStartWithMicroseconds(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:    Daily,
		Count:   3,
		Dtstart: time.Date(1997, 9, 2, 9, 0, 0, 500000000, time.UTC),
	},
	)
	assert.NoError(t, err)
	want := []time.Time{
		testDtstart,
		time.Date(1997, 9, 3, 9, 0, 0, 0, time.UTC),
		time.Date(1997, 9, 4, 9, 0, 0, 0, time.UTC),
	}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestUntil(t *testing.T) {
	t.Parallel()
	r1, err := NewRRule(ROption{
		Freq:    Daily,
		Dtstart: time.Date(1997, 9, 2, 0, 0, 0, 0, time.UTC),
	},
	)
	assert.NoError(t, err)
	r1.Until(time.Date(1998, 9, 2, 0, 0, 0, 0, time.UTC))

	r2, err := NewRRule(ROption{
		Freq:    Daily,
		Dtstart: time.Date(1997, 9, 2, 0, 0, 0, 0, time.UTC),
		Until:   time.Date(1998, 9, 2, 0, 0, 0, 0, time.UTC),
	},
	)
	assert.NoError(t, err)

	v1 := r1.All()
	v2 := r2.All()
	if !timesEqual(v1, v2) {
		t.Errorf("get %v, want %v", v1, v2)
	}

	r3, _ := NewRRule(ROption{
		Freq: Monthly, Dtstart: time.Date(MAXYEAR-100, 1, 1, 0, 0, 0, 0, time.UTC),
	},
	)
	assert.NoError(t, err)
	r3.Until(time.Date(MAXYEAR+100, 1, 1, 0, 0, 0, 0, time.UTC))
	v3 := r3.All()
	if len(v3) != 101*12 {
		t.Errorf("get %v, want %v", len(v3), 101*12)
	}
}

func TestMaxYear(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:       Yearly,
		Count:      3,
		Bymonth:    []int{2},
		Bymonthday: []int{31},
		Dtstart:    testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{}
	value := r.All()
	assert.Equal(t, want, value)
}

func TestBefore(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq: Daily,
		// Count:5,
		Dtstart: testDtstart,
	})
	assert.NoError(t, err)
	want := time.Date(1997, 9, 4, 9, 0, 0, 0, time.UTC)
	value := r.Before(time.Date(1997, 9, 5, 9, 0, 0, 0, time.UTC), false)
	assert.Equal(t, want, value)
}

func TestBeforeInc(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq: Daily,
		// Count:5,
		Dtstart: testDtstart,
	})
	assert.NoError(t, err)
	want := time.Date(1997, 9, 5, 9, 0, 0, 0, time.UTC)
	value := r.Before(time.Date(1997, 9, 5, 9, 0, 0, 0, time.UTC), true)
	assert.Equal(t, want, value)
}

func TestAfter(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq: Daily,
		// Count:5,
		Dtstart: testDtstart,
	})
	assert.NoError(t, err)

	want := time.Date(1997, 9, 5, 9, 0, 0, 0, time.UTC)
	value := r.After(time.Date(1997, 9, 4, 9, 0, 0, 0, time.UTC), false)
	assert.Equal(t, want, value)
}

func TestAfterInc(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq: Daily,
		// Count:5,
		Dtstart: testDtstart,
	})
	assert.NoError(t, err)
	want := time.Date(1997, 9, 4, 9, 0, 0, 0, time.UTC)
	value := r.After(time.Date(1997, 9, 4, 9, 0, 0, 0, time.UTC), true)
	assert.Equal(t, want, value)
}

func TestBetween(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq: Daily,
		// Count:5,
		Dtstart: testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		time.Date(1997, 9, 3, 9, 0, 0, 0, time.UTC),
		time.Date(1997, 9, 4, 9, 0, 0, 0, time.UTC),
		time.Date(1997, 9, 5, 9, 0, 0, 0, time.UTC),
	}
	value := r.Between(testDtstart, time.Date(1997, 9, 6, 9, 0, 0, 0, time.UTC), false)
	assert.Equal(t, want, value)
}

func TestBetweenInc(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq: Daily,
		// Count:5,
		Dtstart: testDtstart,
	})
	assert.NoError(t, err)
	want := []time.Time{
		testDtstart,
		time.Date(1997, 9, 3, 9, 0, 0, 0, time.UTC),
		time.Date(1997, 9, 4, 9, 0, 0, 0, time.UTC),
		time.Date(1997, 9, 5, 9, 0, 0, 0, time.UTC),
		time.Date(1997, 9, 6, 9, 0, 0, 0, time.UTC),
	}
	value := r.Between(testDtstart, time.Date(1997, 9, 6, 9, 0, 0, 0, time.UTC), true)
	assert.Equal(t, want, value)
}

func TestAllWithDefaultUtil(t *testing.T) {
	t.Parallel()
	r, err := NewRRule(ROption{
		Freq:    Yearly,
		Dtstart: testDtstart,
	})
	assert.NoError(t, err)

	value := r.All()
	if len(value) > 300 || len(value) < 200 {
		t.Errorf("No default Util time")
	}

	r, _ = NewRRule(ROption{Freq: Yearly})
	if len(r.All()) != len(value) {
		t.Errorf("No default Util time")
	}
}

func TestWeekdayGetters(t *testing.T) {
	t.Parallel()
	wd := Weekday{n: 2, weekday: 0}
	if wd.N() != 2 {
		t.Errorf("Ord week getter is wrong")
	}
	if wd.Day() != 0 {
		t.Errorf("Day index getter is wrong")
	}
}

func TestRuleChangeDTStartTimezoneRespected(t *testing.T) {
	t.Parallel()
	/*
		https://golang.org/pkg/time/#LoadLocation

		"The time zone database needed by LoadLocation may not be present on all systems, especially non-Unix systems.
		LoadLocation looks in the directory or uncompressed zip file named by the ZONEINFO environment variable,
		if any, then looks in known installation locations on Unix systems, and finally looks in
		$GOROOT/lib/time/zoneinfo.zip."
	*/
	loc, err := time.LoadLocation("CET")
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	rule, err := NewRRule(
		ROption{
			Freq:    Daily,
			Count:   10,
			Wkst:    Monday,
			Dtstart: time.Date(2019, 3, 6, 1, 1, 1, 0, loc),
		},
	)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	rule.DTStart(time.Date(2019, 3, 6, 0, 0, 0, 0, time.UTC))

	events := rule.All()
	if len(events) != 10 {
		t.Fatal("expected", 10, "got", len(events))
	}

	for _, e := range events {
		if e.Location().String() != "UTC" {
			t.Fatal("expected", "UTC", "got", e.Location().String())
		}
		h, m, s := e.Clock()
		if (h + m + s) != 0 {
			t.Fatal("expected", "0", "got", h, m, s)
		}
	}
}

func BenchmarkIterator(b *testing.B) {
	type testCase struct {
		Name   string
		Option ROption
	}
	dtstart := time.Date(2000, 0o3, 22, 12, 0, 0, 0, time.UTC)
	for _, c := range []testCase{
		{
			Name: "simple secondly",
			Option: ROption{
				Dtstart: dtstart,
				Freq:    Secondly,
			},
		},
		{
			Name: "simple minutely",
			Option: ROption{
				Dtstart: dtstart,
				Freq:    Minutely,
			},
		},
		{
			Name: "simple hourly",
			Option: ROption{
				Dtstart: dtstart,
				Freq:    Hourly,
			},
		},
		{
			Name: "simple daily",
			Option: ROption{
				Dtstart: dtstart,
				Freq:    Daily,
			},
		},
		{
			Name: "simple weekly",
			Option: ROption{
				Dtstart: dtstart,
				Freq:    Weekly,
			},
		},
		{
			Name: "simple monthly",
			Option: ROption{
				Dtstart: dtstart,
				Freq:    Monthly,
			},
		},
		{
			Name: "simple yearly",
			Option: ROption{
				Dtstart: dtstart,
				Freq:    Yearly,
			},
		},
	} {
		c := c
		b.Run(c.Name, func(b *testing.B) {
			rrule, err := NewRRule(c.Option)
			if err != nil {
				b.Errorf("failed to init rrule: %s", err)
			}

			for i := 0; i < b.N; i++ {
				res := iterateNum(rrule.Iterator(), 200)
				if res.IsZero() {
					b.Error("expected not zero iterator result")
				}
			}
		})
	}
}

func iterateNum(iter Next, num int) (last time.Time) {
	for i := 0; i < num; i++ {
		var ok bool
		last, ok = iter()
		if !ok {
			return time.Time{}
		}
	}
	return last
}
