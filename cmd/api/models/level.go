package models

import "time"

type Level struct {
	Value        string    `json:"value"`
	LastModified time.Time `json:"lastModified"`
}

type LevelStats struct {
	LevelO Level `json:"PClearanceLevel"`
}

type LevelProfiles struct {
	ProfileID string     `json:"profileId"`
	Stats     LevelStats `json:"stats"`
}

type LevelModel struct {
	Profiles []LevelProfiles `json:"profiles"`
}
