package model

import "time"

type Level struct {
	Value        *string    `json:"value"`
	Lastmodified *time.Time `json:"lastModified"`
}

type TimePlayed struct {
	Value        *string    `json:"value"`
	Lastmodified *time.Time `json:"lastModified"`
}

type Stats struct {
	TimePlayed TimePlayed `json:"PPvPTimePlayed"`
	Level      Level      `json:"PClearanceLevel"`
}

type TimeAndLevelProfiles struct {
	ProfileID string `json:"profileId"`
	Stats     Stats  `json:"stats"`
}

type TimeAndLevelModel struct {
	Profiles []TimeAndLevelProfiles `json:"profiles"`
}
