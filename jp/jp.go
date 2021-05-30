package jp

import (
	"time"

	"github.com/kenkyu392/go-holidays"
)

// JST is Japan Standard Time location
var JST = time.FixedZone("Asia/Tokyo", 9*60*60)

// Add executes Holidays.Add.
func Add(h ...*holidays.Holiday) holidays.Holidays {
	return Holidays.Add(h...)
}

// Remove executes Holidays.Remove.
func Remove(t ...time.Time) holidays.Holidays {
	return Holidays.Remove(t...)
}

// Clone executes Holidays.Clone.
func Clone() holidays.Holidays {
	return Holidays.Clone()
}

// Range executes Holidays.Range.
func Range(f func(h *holidays.Holiday) bool) {
	Holidays.Range(f)
}

// IsHoliday executes Holidays.IsHoliday.
func IsHoliday(t time.Time) *holidays.Holiday {
	return Holidays.IsHoliday(t)
}

// NextHoliday executes Holidays.NextHoliday.
func NextHoliday(t time.Time) *holidays.Holiday {
	return Holidays.NextHoliday(t)
}

// PrevHoliday executes Holidays.PrevHoliday.
func PrevHoliday(t time.Time) *holidays.Holiday {
	return Holidays.PrevHoliday(t)
}

// Between executes Holidays.Between.
func Between(t1, t2 time.Time) holidays.Holidays {
	return Holidays.Between(t1, t2)
}
