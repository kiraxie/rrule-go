package rrule

import "errors"

var (
	ErrInvalidValue     = errors.New("invalid value")
	ErrRuleConflict     = errors.New("rule conflict")
	ErrInvalidBound     = errors.New("invalid bound")
	ErrIntervalLessZero = errors.New("interval must be greater than 0")
	ErrInvalidFrequency = errors.New("invalid frequency")
)
