package model

type TrendsRole struct {
	//Type                   string    `json:"type"`
	//StatsDetail            string    `json:"statsDetail"`
	StatsPeriod   []string `json:"statsPeriod"`
	MatchesPlayed []int    `json:"matchesPlayed"`
	RoundsPlayed  []int    `json:"roundsPlayed"`
	//MinutesPlayed          []int     `json:"minutesPlayed"`
	MatchesWon  []int `json:"matchesWon"`
	MatchesLost []int `json:"matchesLost"`
	RoundsWon   []int `json:"roundsWon"`
	RoundsLost  []int `json:"roundsLost"`
	Kills       []int `json:"kills"`
	//Assists                []int     `json:"assists"`
	Death []int `json:"death"`
	//Headshots              []int     `json:"headshots"`
	//MeleeKills             []int     `json:"meleeKills"`
	TeamKills              []int     `json:"teamKills"`
	OpeningKills           []int     `json:"openingKills"`
	OpeningDeaths          []int     `json:"openingDeaths"`
	Trades                 []int     `json:"trades"`
	OpeningKillTrades      []int     `json:"openingKillTrades"`
	OpeningDeathTrades     []int     `json:"openingDeathTrades"`
	Revives                []int     `json:"revives"`
	DistanceTravelled      []int     `json:"distanceTravelled"`
	WinLossRatio           []float32 `json:"winLossRatio"`
	KillDeathRatio         []float32 `json:"killDeathRatio"`
	HeadshotAccuracy       []float32 `json:"headshotAccuracy"`
	KillsPerRound          []float32 `json:"killsPerRound"`
	RoundsWithAKill        []float32 `json:"roundsWithAKill"`
	RoundsWithMultiKill    []float32 `json:"roundsWithMultiKill"`
	RoundsWithOpeningKill  []float32 `json:"roundsWithOpeningKill"`
	RoundsWithOpeningDeath []float32 `json:"roundsWithOpeningDeath"`
	RoundsWithKOST         []float32 `json:"roundsWithKOST"`
	RoundsSurvived         []float32 `json:"roundsSurvived"`
	RoundsWithAnAce        []float32 `json:"roundsWithAnAce"`
	RoundsWithClutch       []float32 `json:"roundsWithClutch"`
	//TimeAlivePerMatch      []float32 `json:"timeAlivePerMatch"`
	//TimeDeadPerMatch       []float32 `json:"timeDeadPerMatch"`
	//DistancePerRound       []float32 `json:"distancePerRound"`
}

type TrendsRoles struct {
	All      []TrendsRole `json:"all"`
	Attacker []TrendsRole `json:"attacker"`
	Defender []TrendsRole `json:"defender"`
}

type TrendsGameMode struct {
	//Type      string      `json:"type"`
	TeamRoles TrendsRoles `json:"teamRoles"`
}

type TrendGameModes struct {
	All      TrendsGameMode `json:"all"`
	Casual   TrendsGameMode `json:"casual"`
	Ranked   TrendsGameMode `json:"ranked"`
	Unranked TrendsGameMode `json:"unranked"`
}

type TrendsPlatform struct {
	GameModes TrendGameModes `json:"gameModes"`
}

type TrendsPlatforms struct {
	Pc   TrendsPlatform `json:"PC,omitempty"`
	Xbox TrendsPlatform `json:"XONE,omitempty"`
	Ps4  TrendsPlatform `json:"PS4,omitempty"`
}

type TrendsModel struct {
	//ProfileID string          `json:"profileId"`
	//StartDate int             `json:"startDate"`
	//EndDate   int             `json:"endDate"`
	//Region    string          `json:"region"`
	//StatType  string          `json:"statType"`
	Platforms TrendsPlatforms `json:"platforms"`
}
