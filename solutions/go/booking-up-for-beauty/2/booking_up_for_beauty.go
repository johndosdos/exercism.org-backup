package booking

import (
    "time"
)

const (
	layout1 = "1/2/2006 15:04:05"
    layout2 = "January 2, 2006 15:04:05"
    layout3 = "Monday, January 2, 2006 15:04:05"
)

func parse(layout, s string) time.Time {
    t, _ := time.ParseInLocation(layout, s, time.UTC) 
    return t
}

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	return parse(layout1, date)
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	parsed := parse(layout2, date)
	return parsed.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
    parsed := parse(layout3, date)
    hh := parsed.Hour()
	return hh >= 12 && hh < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
    parsed := parse(layout1, date)
	return "You have an appointment on " + parsed.Format("Monday, January 2, 2006, at 15:04.")
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), time.September, 15, 0, 0, 0, 0, time.UTC)
}
