package holidays

//go:generate go run cmd/gen_jp_code/main.go
//go:generate goimports -w .

import (
	"sort"
	"time"
)

// Holiday ...
type Holiday struct {
	Time time.Time         `json:"time"`
	I18n map[string]string `json:"i18n"`
	Lang string            `json:"lang"`
}

// String ...
func (h *Holiday) String() string {
	return h.I18n[h.Lang]
}

// Clone returns a new Holiday with the same value.
func (h *Holiday) Clone() *Holiday {
	i18n := make(map[string]string)
	for k, v := range h.I18n {
		i18n[k] = v
	}
	return &Holiday{Time: h.Time, I18n: i18n, Lang: h.Lang}
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
	for i := 0; i < len(hs); i++ {
		if !f(hs[i].Clone()) {
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
	for i := 0; i < len(hs); i++ {
		if hs[i].Time.After(t) {
			return hs[i].Clone()
		}
	}
	return nil
}

// PrevHoliday returns the previous holiday before t.
func (hs Holidays) PrevHoliday(t time.Time) *Holiday {
	for i := len(hs) - 1; i >= 0; i-- {
		if hs[i].Time.Before(t) {
			return hs[i].Clone()
		}
	}
	return nil
}

// Between returns the holidays that exist between t1 and t2.
func (hs Holidays) Between(t1, t2 time.Time) Holidays {
	_hs := make(Holidays, 0)
	for next := hs.NextHoliday(t1); next != nil && next.Time.Before(t2); {
		_hs = append(_hs, next.Clone())
		next = hs.NextHoliday(next.Time)
	}
	_hs.sort()
	return _hs
}

func (hs Holidays) sort() {
	sort.Slice(hs, func(i, j int) bool {
		return hs[i].Time.Before(hs[j].Time)
	})
}
