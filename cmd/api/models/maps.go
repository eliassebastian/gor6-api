package models

type MapsModel struct {
	ProfileID string `json:"profileId"`
	//StartDate int       `json:"startDate"`
	//EndDate   int       `json:"endDate"`
	//Region    string    `json:"region"`
	//StatType  string    `json:"statType"`
	Platforms MapsPlatforms `json:"platforms"`
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
	Value float64 `json:"value"`
	P     int     `json:"p"`
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
	Value int `json:"value"`
	P     int `json:"p"`
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
type Map struct {
	//Type          string `json:"type"`
	//StatsType     string `json:"statsType"`
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
	//WinLossRatio           int                    `json:"winLossRatio"`
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
type MapsTeamRoles struct {
	All      []Map `json:"all"`
	Attacker []Map `json:"attacker"`
	Defender []Map `json:"defender"`
}
type MapsGameMode struct {
	Type      string        `json:"type"`
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
