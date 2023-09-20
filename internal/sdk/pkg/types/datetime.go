// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package types

import "time"

// MustTimeFromString returns a time.Time from a string formatted as "2006-01-02T15:04:05Z07:00" or panics.
// Avoid using this function in production code.
func MustTimeFromString(str string) time.Time {
	t, err := time.Parse(time.RFC3339Nano, str)
	if err != nil {
		panic(err)
	}

	return t
}

// MustNewTimeFromString returns an instance of time.Time from a string formatted as "2006-01-02T15:04:05Z07:00" or panics.
// Avoid using this function in production code.
func MustNewTimeFromString(str string) *time.Time {
	t := MustTimeFromString(str)
	return &t
}
