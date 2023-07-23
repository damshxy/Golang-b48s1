package utilities

import (
	"strconv"
	"time"
)

func CountDuration(d1 time.Time, d2 time.Time) string {

	diff := d2.Sub(d1)
	days := int(diff.Hours() / 24)
	weeks := days / 7
	months := days / 30

	if months > 12 {
		return strconv.Itoa(months/12) + " Year"
	}
	if months > 0 {
		return strconv.Itoa(months) + " Month"
	}
	if weeks > 0 {
		return strconv.Itoa(weeks) + " Week"
	}
	return strconv.Itoa(days) + " Day"
}