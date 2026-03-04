package http

import "time"

type Record struct {
	Username  string    `json:"username"`
	Time      float32   `json:"time"`
	CreatedAt time.Time `json:"createdAt"`
}
