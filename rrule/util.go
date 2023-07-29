package rrule

import (
	"fmt"
	"strconv"
	"strings"
)

func repeat(value, count int) []int {
	result := []int{}
	for i := 0; i < count; i++ {
		result = append(result, value)
	}

	return result
}

func concat(slices ...[]int) []int {
	result := []int{}
	for _, item := range slices {
		result = append(result, item...)
	}

	return result
}

func rang(start, end int) []int {
	result := []int{}
	for i := start; i < end; i++ {
		result = append(result, i)
	}

	return result
}

func appendOption(options []string, key string, value []int) []string {
	if len(value) == 0 {
		return options
	}

	return append(options, fmt.Sprintf("%s=%s", key, strings.Join(toStringSlice(value), ",")))
}

func toStringSlice(v []int) []string {
	slice := make([]string, 0, len(v))
	for _, v := range v {
		slice = append(slice, strconv.Itoa(v))
	}

	return slice
}

func checkBounds(param string, value int, bounds []int, plusMinus bool) error {
	if !(value >= bounds[0] && value <= bounds[1]) && (!plusMinus || !(value <= -bounds[0] && value >= -bounds[1])) {
		plusMinusBounds := ""
		if plusMinus {
			plusMinusBounds = fmt.Sprintf(" or %d and %d", -bounds[0], -bounds[1])
		}

		return fmt.Errorf("%w: %s must be between %d and %d%s", ErrInvalidBound, param, bounds[0], bounds[1], plusMinusBounds)
	}

	return nil
}
