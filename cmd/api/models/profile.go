package models

type PlayerProfile struct {
	//ProfileID string `json:"profileId"`
	//UserID         string `json:"userId"`
	PlatformType   string `json:"platformType"`
	IDOnPlatform   string `json:"idOnPlatform"`
	NameOnPlatform string `json:"nameOnPlatform"`
}

type PlayerIDModel struct {
	Profiles []PlayerProfile `json:"profiles"`
}
