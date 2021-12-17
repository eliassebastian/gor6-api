package models

//ps4  https://r6s-stats.ubisoft.com/v1/current/summary/79239d25-3da1-401c-9231-ffb2e50f7c1d?gameMode=all,ranked,unranked,casual&platform=PS4&startDate=20210815&endDate=20211213
// https://r6s-stats.ubisoft.com/v1/seasonal/summary/ab1ff7ae-13e4-4a6a-9b03-317285f8057b?gameMode=all,ranked,casual,unranked&platform=PC
// https://r6s-stats.ubisoft.com/v1/seasonal/summary/ab1ff7ae-13e4-4a6a-9b03-317285f8057b?gameMode=all,ranked,casual,unranked&platform=PC
//GenomeId  fd4135bb-409a-4e90-8587-a945f92e6c6d

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
	TeamKills         int     `json:"teamKills"`
	OpeningKills      int     `json:"openingKills"`
	OpeningDeaths     int     `json:"openingDeaths"`
	Trades            int     `json:"trades"`
	TimeAlivePerMatch float64 `json:"timeAlivePerMatch"`
	TimeDeadPerMatch  float64 `json:"timeDeadPerMatch"`
	DistancePerRound  float64 `json:"distancePerRound"`
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
	Xbox SummaryPlatform `json:"XBOX,omitempty"`
	Ps4  SummaryPlatform `json:"PS4,omitempty"`
}
