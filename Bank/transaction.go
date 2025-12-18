package Bank

import "time"

type Transaction struct {
	Type      string
	Amount    float64
	Balance   float64
	Timestamp time.Time
}
