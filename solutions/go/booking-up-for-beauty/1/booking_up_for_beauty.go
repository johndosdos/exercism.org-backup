package booking

import (
    "fmt"
    "time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	parsed, _ := time.Parse("1/02/2006 15:04:05", date)
    return parsed
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	layout := "January 2, 2006 15:04:05"
    parsed, _ := time.Parse(layout, date)
	return parsed.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
    layout := "Monday, January 2, 2006 15:04:05"
    parsed, _ := time.Parse(layout, date)
    hh, _, _ := parsed.Clock()
	return hh >= 12 && hh < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
    layout := "1/2/2006 15:04:05"
    parsed, _ := time.Parse(layout, date)
    yyyy, mm, dd := parsed.Date()
    h, m, _ := parsed.Clock()
    day := parsed.Weekday().String()
    return fmt.Sprintf("You have an appointment on %s, %s %d, %d, at %d:%d.", day, mm, dd, yyyy, h, m)
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
    layout := "2006-1-2 15:04:05 -0700"
    parsed, _ := time.Parse(layout, "2025-09-15 00:00:00 +0000")
    return parsed.UTC()
}
