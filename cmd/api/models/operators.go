package model

type OperatorModel struct {
	ProfileID string            `json:"profileId"`
	Platforms OperatorPlatforms `json:"platforms"`
}

type OperatorSide struct {
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

type OperatorTeamRoles struct {
	Attacker []OperatorSide `json:"attacker"`
	Defender []OperatorSide `json:"defender"`
}

type OperatorGameMode struct {
	//Type      string    `json:"type"`
	TeamRoles OperatorTeamRoles `json:"teamRoles"`
}

type OperatorGameModes struct {
	All      OperatorGameMode `json:"all"`
	Casual   OperatorGameMode `json:"casual"`
	Ranked   OperatorGameMode `json:"ranked"`
	Unranked OperatorGameMode `json:"unranked"`
}

type OperatorPlatform struct {
	GameModes OperatorGameModes `json:"gameModes"`
}

type OperatorPlatforms struct {
	Pc   OperatorPlatform `json:"PC,omitempty"`
	Xbox OperatorPlatform `json:"XONE,omitempty"`
	Ps4  OperatorPlatform `json:"PS4,omitempty"`
}
