package holidays

//go:generate go run cmd/gen_jp_code/main.go
//go:generate goimports -w .

import (
	"sort"
	"time"
)

// Holiday ...
type Holiday struct {
	Name string    `json:"name"`
	Time time.Time `json:"time"`
}

// Clone returns a new Holiday with the same value.
func (h *Holiday) Clone() *Holiday {
	return &Holiday{Name: h.Name, Time: h.Time}
}

// Equal reports whether h and t represent the same date instant.
func (h *Holiday) Equal(t time.Time) bool {
	y1, m1, d1 := t.Date()
	y2, m2, d2 := h.Time.Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}

// Holidays ...
type Holidays []*Holiday

// Add h and return a new instance.
func (hs Holidays) Add(h ...*Holiday) Holidays {
	_hs := hs.Clone()
	_hs = append(_hs, h...)
	_hs.sort()
	return _hs
}

// Remove t and return a new instance.
func (hs Holidays) Remove(t ...time.Time) Holidays {
	_hs := make(Holidays, 0)
	for i := 0; i < len(hs); i++ {
		found := false
		for j := 0; j < len(t); j++ {
			if hs[i].Equal(t[j]) {
				found = true
				break
			}
		}
		if found {
			continue
		}
		_hs = append(_hs, hs[i].Clone())
	}
	_hs.sort()
	return _hs
}

// Clone returns a new Holidays with the same value.
func (hs Holidays) Clone() Holidays {
	_hs := make(Holidays, len(hs))
	for i := 0; i < len(hs); i++ {
		_hs[i] = hs[i].Clone()
	}
	_hs.sort()
	return _hs
}

// Range calls f sequentially for each h present in the hs.
// If f returns false, range stops the iteration.
func (hs Holidays) Range(f func(h *Holiday) bool) {
	for _, h := range hs {
		if !f(h.Clone()) {
			break
		}
	}
}

// IsHoliday returns Holiday if t is a holiday.
func (hs Holidays) IsHoliday(t time.Time) *Holiday {
	for _, h := range hs {
		if h.Equal(t) {
			return h.Clone()
		}
	}
	return nil
}

// NextHoliday returns the next holiday after t.
func (hs Holidays) NextHoliday(t time.Time) *Holiday {
	for _, h := range hs {
		if h.Time.After(t) {
			return h.Clone()
		}
	}
	return nil
}

func (hs Holidays) sort() {
	sort.Slice(hs, func(i, j int) bool {
		return hs[i].Time.Before(hs[j].Time)
	})
}
