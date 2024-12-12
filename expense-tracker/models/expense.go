package models

import (
	"time"
)

type Expense struct {
	ID          int64
	Date        time.Time
	Description string
	Amount      int64
	Category    string
}

type Budget struct {
	Month  string
	Amount int64
}
