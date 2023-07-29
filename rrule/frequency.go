package rrule

import "fmt"

type Frequency int

const (
	Yearly Frequency = iota
	Monthly
	Weekly
	Daily
)

func (t Frequency) String() string {
	switch t {
	case Yearly:
		return "YEARLY"
	case Monthly:
		return "MONTHLY"
	case Weekly:
		return "WEEKLY"
	case Daily:
		return "DAILY"
	default:
		return "UNKNOWN"
	}
}

func NewFrequencyFromString(s string) (Frequency, error) {
	switch s {
	case "YEARLY":
		return Yearly, nil
	case "MONTHLY":
		return Monthly, nil
	case "WEEKLY":
		return Weekly, nil
	case "DAILY":
		return Daily, nil
	default:
		return -1, fmt.Errorf("%w: %s", ErrInvalidValue, s)
	}
}
