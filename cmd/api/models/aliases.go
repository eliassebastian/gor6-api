package models

import "time"

type Aliases struct {
	Username string    `json:"username"`
	Date     time.Time `json:"date"`
}
