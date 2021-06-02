package holidays

import (
	"testing"
	"time"

	"golang.org/x/text/language"
)

func TestHolidays(t *testing.T) {
	hs := Holidays{
		{
			Time: time.Date(2021, 7, 22, 0, 0, 0, 0, time.UTC),
			I18n: map[language.Tag]string{
				language.Japanese: "海の日",
				language.English:  "Marine Day",
			},
			Tag: language.Japanese,
		},
		{
			Time: time.Date(2021, 7, 23, 0, 0, 0, 0, time.UTC),
			I18n: map[language.Tag]string{
				language.Japanese: "スポーツの日",
				language.English:  "Health and Sports Day",
			},
			Tag: language.Japanese,
		},
		{
			Time: time.Date(2021, 8, 8, 0, 0, 0, 0, time.UTC),
			I18n: map[language.Tag]string{
				language.Japanese: "山の日",
				language.English:  "Mountain Day",
			},
			Tag: language.Japanese,
		},
		{
			Time: time.Date(2021, 8, 9, 0, 0, 0, 0, time.UTC),
			I18n: map[language.Tag]string{
				language.Japanese: "振替休日（山の日）",
				language.English:  "Substitute Holiday (Mountain Day)",
			},
			Tag: language.Japanese,
		},
		{
			Time: time.Date(2021, 9, 20, 0, 0, 0, 0, time.UTC),
			I18n: map[language.Tag]string{
				language.Japanese: "敬老の日",
				language.English:  "Respect for the Aged Day",
			},
			Tag: language.Japanese,
		},
		{
			Time: time.Date(2021, 9, 23, 0, 0, 0, 0, time.UTC),
			I18n: map[language.Tag]string{
				language.Japanese: "秋分の日",
				language.English:  "Autumnal Equinox Day",
			},
			Tag: language.Japanese,
		},
		{
			Time: time.Date(2021, 11, 3, 0, 0, 0, 0, time.UTC),
			I18n: map[language.Tag]string{
				language.Japanese: "文化の日",
				language.English:  "Culture Day",
			},
			Tag: language.Japanese,
		},
		{
			Time: time.Date(2021, 11, 23, 0, 0, 0, 0, time.UTC),
			I18n: map[language.Tag]string{
				language.Japanese: "勤労感謝の日",
				language.English:  "Labor Thanksgiving Day",
			},
			Tag: language.Japanese,
		},
	}

	healthAndSportsDay := time.Date(2021, 7, 23, 1, 2, 3, 4, time.UTC)
	mountainDay := time.Date(2021, 8, 8, 0, 0, 0, 0, time.UTC)
	cultureDay := time.Date(2021, 11, 3, 0, 0, 0, 0, time.UTC)
	specialDay := time.Date(3000, 1, 1, 0, 0, 0, 0, time.UTC)

	t.Run("String", func(t *testing.T) {
		if want, got := "海の日", hs[0].String(); want != got {
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
		end := time.Date(2021, 9, 30, 0, 0, 0, 0, time.UTC)
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
		_hs := hs.Add(&Holiday{Time: specialDay, Tag: language.Japanese, I18n: map[language.Tag]string{
			language.Japanese: "特別な日",
			language.English:  "Special Day",
		}})
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
