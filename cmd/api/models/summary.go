package model

type SummaryModel struct {
	ProfileID string           `json:"profileId"`
	Platforms SummaryPlatforms `json:"platforms"`
}

type SummarySeason struct {
	SeasonYear        string  `json:"seasonYear"`
	SeasonNumber      string  `json:"seasonNumber"`
	MatchesPlayed     int     `json:"matchesPlayed"`
	RoundsPlayed      int     `json:"roundsPlayed"`
	MinutesPlayed     int     `json:"minutesPlayed"`
	MatchesWon        int     `json:"matchesWon"`
	MatchesLost       int     `json:"matchesLost"`
	RoundsWon         int     `json:"roundsWon"`
	RoundsLost        int     `json:"roundsLost"`
	Kills             int     `json:"kills"`
	Assists           int     `json:"assists"`
	Death             int     `json:"death"`
	Headshots         int     `json:"headshots"`
	MeleeKills        int     `json:"meleeKills"`
	OpeningKills      int     `json:"openingKills"`
	OpeningDeaths     int     `json:"openingDeaths"`
	Trades            int     `json:"trades"`
	TimeAlivePerMatch float32 `json:"timeAlivePerMatch" bson:"timeAlivePerMatch_F32"`
	TimeDeadPerMatch  float32 `json:"timeDeadPerMatch" bson:"timeDeadPerMatch_F32"`
	DistancePerRound  float32 `json:"distancePerRound" bson:"distancePerRound_F32"`
}

type SummaryTeamRoles struct {
	All []SummarySeason `json:"all"`
}

type SummaryGameMode struct {
	TeamRoles SummaryTeamRoles `json:"teamRoles"`
}

type SummaryGameModes struct {
	Casual   SummaryGameMode `json:"casual"`
	Ranked   SummaryGameMode `json:"ranked"`
	Unranked SummaryGameMode `json:"unranked"`
	All      SummaryGameMode `json:"all"`
}

type SummaryPlatform struct {
	SGameModes SummaryGameModes `json:"gameModes"`
}

type SummaryPlatforms struct {
	Pc   SummaryPlatform `json:"PC,omitempty"`
	Xbox SummaryPlatform `json:"XONE,omitempty"`
	Ps4  SummaryPlatform `json:"PS4,omitempty"`
}
