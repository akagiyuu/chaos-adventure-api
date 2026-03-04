package domain

import "time"

type Record struct {
	Username  string
	Time      float32
	CreatedAt time.Time
}
