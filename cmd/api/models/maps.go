package model

type MapsModel struct {
	ProfileID string        `json:"profileId"`
	Platforms MapsPlatforms `json:"platforms"`
}
type Map struct {
	StatsDetail       string  `json:"statsDetail"`
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
	TimeAlivePerMatch float32 `json:"timeAlivePerMatch"`
	TimeDeadPerMatch  float32 `json:"timeDeadPerMatch"`
	DistancePerRound  float32 `json:"distancePerRound"`
}
type MapsTeamRoles struct {
	All      []Map `json:"all"`
	Attacker []Map `json:"attacker"`
	Defender []Map `json:"defender"`
}
type MapsGameMode struct {
	//Type      string        `json:"type"`
	TeamRoles MapsTeamRoles `json:"teamRoles"`
}
type MapsGameModes struct {
	All      MapsGameMode `json:"all"`
	Casual   MapsGameMode `json:"casual"`
	Ranked   MapsGameMode `json:"ranked"`
	Unranked MapsGameMode `json:"unranked"`
}
type MapsPlatform struct {
	GameModes MapsGameModes `json:"gameModes"`
}
type MapsPlatforms struct {
	Pc   MapsPlatform `json:"PC,omitempty"`
	Xbox MapsPlatform `json:"XONE,omitempty"`
	Ps4  MapsPlatform `json:"PS4,omitempty"`
}
