package rrule

import "time"

type Option func(*RRule)

func Count(v int) func(r *RRule) {
	return func(r *RRule) {
		r.Count = v
	}
}

func Until(v time.Time) func(r *RRule) {
	return func(r *RRule) {
		r.Until = v
	}
}

func BySecond(v ...int) func(r *RRule) {
	return func(r *RRule) {
		r.BySecond = v
	}
}

func ByMinute(v ...int) func(r *RRule) {
	return func(r *RRule) {
		r.ByMinute = v
	}
}

func ByHour(v ...int) func(r *RRule) {
	return func(r *RRule) {
		r.ByHour = v
	}
}

func ByDay(v []time.Weekday) func(r *RRule) {
	return func(r *RRule) {
		r.ByDay = v
	}
}

func ByMonthDay(v ...int) func(r *RRule) {
	return func(r *RRule) {
		r.ByMonthDay = v
	}
}

func ByYearDay(v ...int) func(r *RRule) {
	return func(r *RRule) {
		r.ByYearDay = v
	}
}

func ByWeekNo(v ...int) func(r *RRule) {
	return func(r *RRule) {
		r.ByWeekNo = v
	}
}

func ByMonth(v ...int) func(r *RRule) {
	return func(r *RRule) {
		r.ByMonth = v
	}
}

func BySetPos(v ...int) func(r *RRule) {
	return func(r *RRule) {
		r.BySetpos = v
	}
}

func WeekStart(v time.Weekday) func(r *RRule) {
	return func(r *RRule) {
		r.WeekStart = v
	}
}
