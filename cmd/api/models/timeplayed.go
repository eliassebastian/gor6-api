package models

import "time"

type TimePlayed struct {
	Value        string    `json:"value"`
	LastModified time.Time `json:"lastModified"`
}
type TimeStats struct {
	TimePlayed TimePlayed `json:"PPvPTimePlayed"`
}
type TimePlayedProfiles struct {
	ProfileID string    `json:"profileId"`
	Stats     TimeStats `json:"stats"`
}

type TimePlayedModel struct {
	Profiles []TimePlayedProfiles `json:"profiles"`
}
