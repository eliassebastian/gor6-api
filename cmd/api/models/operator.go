package models

type OperatorModel struct {
	ProfileID string            `json:"profileId"`
	Platforms OperatorPlatforms `json:"platforms"`
}

/*
type KillDeathRatio struct {
	Value float64 `json:"value"`
	P     int     `json:"p"`
}

type HeadshotAccuracy struct {
	Value float64 `json:"value"`
	P     int     `json:"p"`
}

type KillsPerRound struct {
	Value int `json:"value"`
	P     int `json:"p"`
}

type RoundsWithAKill struct {
	Value float64 `json:"value"`
	P     int     `json:"p"`
}

type RoundsWithMultiKill struct {
	Value float64 `json:"value"`
	P     int     `json:"p"`
}

type RoundsWithOpeningKill struct {
	Value float64 `json:"value"`
	P     int     `json:"p"`
}

type RoundsWithOpeningDeath struct {
	Value float64 `json:"value"`
	P     int     `json:"p"`
}

type RoundsWithKOST struct {
	Value float64 `json:"value"`
	P     int     `json:"p"`
}

type RoundsSurvived struct {
	Value float64 `json:"value"`
	P     int     `json:"p"`
}

type RoundsWithAnAce struct {
	Value int `json:"value"`
	P     int `json:"p"`
}

type RoundsWithClutch struct {
	Value int `json:"value"`
	P     int `json:"p"`
}
*/
type OperatorSide struct {
	//Type      string `json:"type"`
	//StatsType string `json:"statsType"`
	StatsDetail   string `json:"statsDetail"`
	MatchesPlayed int    `json:"matchesPlayed"`
	RoundsPlayed  int    `json:"roundsPlayed"`
	MinutesPlayed int    `json:"minutesPlayed"`
	MatchesWon    int    `json:"matchesWon"`
	MatchesLost   int    `json:"matchesLost"`
	RoundsWon     int    `json:"roundsWon"`
	RoundsLost    int    `json:"roundsLost"`
	Kills         int    `json:"kills"`
	Assists       int    `json:"assists"`
	Death         int    `json:"death"`
	Headshots     int    `json:"headshots"`
	MeleeKills    int    `json:"meleeKills"`
	//TeamKills              int                    `json:"teamKills"`
	//OpeningKills           int                    `json:"openingKills"`
	//OpeningDeaths          int                    `json:"openingDeaths"`
	//Trades                 int                    `json:"trades"`
	//OpeningKillTrades      int                    `json:"openingKillTrades"`
	//OpeningDeathTrades     int                    `json:"openingDeathTrades"`
	//Revives                int                    `json:"revives"`
	//DistanceTravelled      int                    `json:"distanceTravelled"`
	//WinLossRatio int `json:"winLossRatio"`
	//KillDeathRatio         KillDeathRatio         `json:"killDeathRatio"`
	//HeadshotAccuracy       HeadshotAccuracy       `json:"headshotAccuracy"`
	//KillsPerRound          KillsPerRound          `json:"killsPerRound"`
	//RoundsWithAKill        RoundsWithAKill        `json:"roundsWithAKill"`
	//RoundsWithMultiKill    RoundsWithMultiKill    `json:"roundsWithMultiKill"`
	//RoundsWithOpeningKill  RoundsWithOpeningKill  `json:"roundsWithOpeningKill"`
	//RoundsWithOpeningDeath RoundsWithOpeningDeath `json:"roundsWithOpeningDeath"`
	//RoundsWithKOST         RoundsWithKOST         `json:"roundsWithKOST"`
	//RoundsSurvived         RoundsSurvived         `json:"roundsSurvived"`
	//RoundsWithAnAce        RoundsWithAnAce        `json:"roundsWithAnAce"`
	//RoundsWithClutch       RoundsWithClutch       `json:"roundsWithClutch"`
	TimeAlivePerMatch float32 `json:"timeAlivePerMatch" bson:"timeAlivePerMatch_F32"`
	TimeDeadPerMatch  float32 `json:"timeDeadPerMatch" bson:"timeDeadPerMatch_F32"`
	DistancePerRound  float32 `json:"distancePerRound" bson:"distancePerRound_F32"`
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
