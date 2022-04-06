package model

import "time"

type Alias struct {
	Name string    `json:"name"`
	Date time.Time `json:"date"`
}
