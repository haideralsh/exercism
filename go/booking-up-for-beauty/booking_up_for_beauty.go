package booking

import (
    "fmt"
    "log"
    "time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
    t, err := time.Parse("1/2/2006 15:04:05", date)

    if err != nil {
        log.Fatal(err)
    }

    return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
    t, err := time.Parse("January 2, 2006 15:04:05", date)

    if err != nil {
        log.Fatal(err)
    }

    return t.Unix() < time.Now().Unix()

}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
    t, err := time.Parse("Monday, January 2, 2006 15:04:05", date)

    if err != nil {
        log.Fatal(err)
    }

    return t.Hour() >= 12 && t.Hour() < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
    // => "You have an appointment on Thursday, July 25, 2019, at 13:45."

    return fmt.Sprintf("You have an appointment on %s.", Schedule(date).Format("Monday, January 2, 2006, at 15:04"))
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
    currentYear := time.Now().Year()

    t, err := time.Parse("2006-01-02", fmt.Sprintf("%d-09-15", currentYear))

    if err != nil {
        log.Fatal("invalid date")
    }

    return t
}
