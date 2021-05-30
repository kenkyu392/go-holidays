package holidays

import (
	"testing"
	"time"
)

func TestHolidays(t *testing.T) {
	hs := Holidays{
		{Name: "元日", Time: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)},
		{Name: "成人の日", Time: time.Date(2021, 1, 11, 0, 0, 0, 0, time.UTC)},
		{Name: "建国記念の日", Time: time.Date(2021, 2, 11, 0, 0, 0, 0, time.UTC)},
		{Name: "天皇誕生日", Time: time.Date(2021, 2, 23, 0, 0, 0, 0, time.UTC)},
		{Name: "春分の日", Time: time.Date(2021, 3, 20, 0, 0, 0, 0, time.UTC)},
		{Name: "昭和の日", Time: time.Date(2021, 4, 29, 0, 0, 0, 0, time.UTC)},
		{Name: "憲法記念日", Time: time.Date(2021, 5, 3, 0, 0, 0, 0, time.UTC)},
		{Name: "みどりの日", Time: time.Date(2021, 5, 4, 0, 0, 0, 0, time.UTC)},
		{Name: "こどもの日", Time: time.Date(2021, 5, 5, 0, 0, 0, 0, time.UTC)},
		{Name: "海の日", Time: time.Date(2021, 7, 22, 0, 0, 0, 0, time.UTC)},
		{Name: "スポーツの日", Time: time.Date(2021, 7, 23, 0, 0, 0, 0, time.UTC)},
		{Name: "山の日", Time: time.Date(2021, 8, 8, 0, 0, 0, 0, time.UTC)},
		{Name: "振替休日（山の日）", Time: time.Date(2021, 8, 9, 0, 0, 0, 0, time.UTC)},
		{Name: "敬老の日", Time: time.Date(2021, 9, 20, 0, 0, 0, 0, time.UTC)},
		{Name: "秋分の日", Time: time.Date(2021, 9, 23, 0, 0, 0, 0, time.UTC)},
		{Name: "文化の日", Time: time.Date(2021, 11, 3, 0, 0, 0, 0, time.UTC)},
		{Name: "勤労感謝の日", Time: time.Date(2021, 11, 23, 0, 0, 0, 0, time.UTC)},
	}

	healthAndSportsDay := time.Date(2021, 7, 23, 1, 2, 3, 4, time.UTC)
	mountainDay := time.Date(2021, 8, 8, 0, 0, 0, 0, time.UTC)
	cultureDay := time.Date(2021, 11, 3, 0, 0, 0, 0, time.UTC)
	specialDay := time.Date(3000, 1, 1, 0, 0, 0, 0, time.UTC)

	t.Run("String", func(t *testing.T) {
		if want, got := "元日", hs[0].String(); want != got {
			t.Errorf(
				"No match\nwant: %v\ngot : %v",
				want, got,
			)
		}
	})

	t.Run("IsHoliday", func(t *testing.T) {
		if got := hs.IsHoliday(healthAndSportsDay); !got.Equal(healthAndSportsDay) {
			t.Errorf(
				"Must be health and sports day\nwant: %v\ngot : %v",
				healthAndSportsDay, got.Time,
			)
		}
	})

	t.Run("NextHoliday", func(t *testing.T) {
		if got := hs.NextHoliday(healthAndSportsDay); !got.Equal(mountainDay) {
			t.Errorf(
				"Must be mountain day\nwant: %v\ngot : %v",
				mountainDay, got.Time,
			)
		}
	})

	t.Run("PrevHoliday", func(t *testing.T) {
		if got := hs.PrevHoliday(mountainDay); !got.Equal(healthAndSportsDay) {
			t.Errorf(
				"Must be health and sports day\nwant: %v\ngot : %v",
				healthAndSportsDay, got.Time,
			)
		}

		if got := hs.PrevHoliday(hs[0].Time); got != nil {
			t.Errorf(
				"Must be nil\nwant: %v\ngot : %v",
				nil, got,
			)
		}
	})

	t.Run("Remove", func(t *testing.T) {
		_hs := hs.Remove(cultureDay)
		if got := _hs.IsHoliday(cultureDay); got != nil {
			t.Errorf(
				"Must not be a holiday\nwant: %v\ngot : %v",
				cultureDay, got.Time,
			)
		}
	})

	t.Run("Range", func(t *testing.T) {
		_hs := make(Holidays, 0)
		end := time.Date(2021, 4, 30, 0, 0, 0, 0, time.UTC)
		hs.Range(func(h *Holiday) bool {
			f := h.Time.Before(end)
			if f {
				_hs = append(_hs, h)
			}
			return f
		})
		if want, got := 6, len(_hs); want != got {
			t.Errorf(
				"No match\nwant: %v\ngot : %v",
				want, got,
			)
		}
	})

	t.Run("Add", func(t *testing.T) {
		_hs := hs.Add(&Holiday{Name: "特別な日", Time: specialDay})
		if got := _hs.IsHoliday(specialDay); !got.Equal(specialDay) {
			t.Errorf(
				"Must be special day\nwant: %v\ngot : %v",
				cultureDay, got.Time,
			)
		}
	})

	t.Run("Between", func(t *testing.T) {
		t1 := time.Date(2021, 7, 0, 0, 0, 0, 0, time.UTC)
		t2 := time.Date(2021, 12, 0, 0, 0, 0, 0, time.UTC)
		_hs := hs.Between(t1, t2)
		if want, got := 8, len(_hs); want != got {
			t.Errorf(
				"No match\nwant: %v\ngot : %v",
				want, got,
			)
		}
	})
}
