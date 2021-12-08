package models

import "time"

type Session struct {
	SessionExpiry time.Time `json:"expiration"`
	SessionKey    string    `json:"sessionKey"`
	SpaceID       string    `json:"spaceId"`
	SessionTicket string    `json:"ticket"`
}
