package models

import "time"

type Level struct {
	Value        string    `json:"value"`
	LastModified time.Time `json:"lastModified"`
}

type TimePlayed struct {
	Value        string    `json:"value"`
	LastModified time.Time `json:"lastModified"`
}

type Stats struct {
	TimePlayedO TimePlayed `json:"PPvPTimePlayed"`
	LevelO      Level      `json:"PClearanceLevel"`
}

type TimeAndLevelProfiles struct {
	ProfileID string `json:"profileId"`
	StatsO    Stats  `json:"stats"`
}

type TimeAndLevelModel struct {
	Profiles []TimeAndLevelProfiles `json:"profiles"`
}
