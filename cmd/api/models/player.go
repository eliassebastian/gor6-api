package models

type PlayerIDModel struct {
	Profiles []struct {
		ProfileID      string `json:"profileId"`
		UserID         string `json:"userId"`
		PlatformType   string `json:"platformType"`
		IDOnPlatform   string `json:"idOnPlatform"`
		NameOnPlatform string `json:"nameOnPlatform"`
	} `json:"profiles"`
}
