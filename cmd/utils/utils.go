package utils

import(
    "time"
)

type Time struct {
    time.Time
}

func (t Time) IsLeapYear() bool {
    if t.Year() % 400 == 0 {
        return true
    }
    return t.Year() % 4 == 0 && t.Year() % 100 != 0
}

func GetCurrentTime() Time {
    return Time{Time: time.Now()}
}
