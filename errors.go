package rrule

import "errors"

var (
	ErrInvalidFreq        = errors.New("invalid frequency")
	ErrIndexOutOfBounds   = errors.New("index out of bounds")
	ErrInvalidWeekday     = errors.New("invalid weekday")
	ErrInvalidRRuleFormat = errors.New("invalid rrule format")
	ErrInvalidateBound    = errors.New("invalid bound")
	ErrBadFormat          = errors.New("bad format")
)
