# go-holidays

go-holidays is a Go library for handling holidays and business days.


## Installation

```
go get -u github.com/kenkyu392/go-holidays
```

## Usage

### Japan

Currently, it supports Japanese holidays from 1955 to 2022.

```go
package main

import (
	"fmt"
	"time"

	"github.com/kenkyu392/go-holidays/jp"
)

func main() {
	healthAndSportsDay := time.Date(2021, 7, 23, 0, 0, 0, 0, jp.JST)
	newYearsDay := time.Date(2022, 1, 1, 0, 0, 0, 0, jp.JST)

	// Change the display language to English.
	// jp.Holidays.SetTag(language.English)

	h := jp.IsHoliday(healthAndSportsDay)
	fmt.Printf("%s: %s\n", h.Time.Format("2006/01/02"), h.String())

	hs := jp.Between(healthAndSportsDay, newYearsDay)
	for _, h := range hs {
		fmt.Printf("%s: %s\n", h.Time.Format("2006/01/02"), h.String())
	}
}
```

<details>
<summary><b><i>Output :</i></b></summary>

```
2021/07/23: スポーツの日
2021/08/08: 山の日
2021/08/09: 振替休日（山の日）
2021/09/20: 敬老の日
2021/09/23: 秋分の日
2021/11/03: 文化の日
2021/11/23: 勤労感謝の日
```

</details>


## License

[MIT](LICENSE)
