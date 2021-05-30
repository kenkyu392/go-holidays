package holidays

//go:generate go run cmd/gen_jp_code/main.go
//go:generate goimports -w .

import (
	"time"
)

// Holiday ...
type Holiday struct {
	Name string    `json:"name"`
	Time time.Time `json:"time"`
}

// Holidays ...
type Holidays []*Holiday
