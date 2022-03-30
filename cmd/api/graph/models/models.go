package model

type All struct {
	Seasonyear        *string  `json:"seasonYear"`
	Seasonnumber      *string  `json:"seasonNumber"`
	Matchesplayed     *int     `json:"matchesPlayed"`
	Roundsplayed      *int     `json:"roundsPlayed"`
	Minutesplayed     *int     `json:"minutesPlayed"`
	Matcheswon        *int     `json:"matchesWon"`
	Matcheslost       *int     `json:"matchesLost"`
	Roundswon         *int     `json:"roundsWon"`
	Roundslost        *int     `json:"roundsLost"`
	Kills             *int     `json:"kills"`
	Assists           *int     `json:"assists"`
	Death             *int     `json:"death"`
	Headshots         *int     `json:"headshots"`
	Meleekills        *int     `json:"meleeKills"`
	Openingkills      *int     `json:"openingKills"`
	Openingdeaths     *int     `json:"openingDeaths"`
	Trades            *int     `json:"trades"`
	Timealivepermatch *float64 `json:"timeAlivePerMatch"`
	Timedeadpermatch  *float64 `json:"timeDeadPerMatch"`
	Distanceperround  *int     `json:"distancePerRound"`
}

type Attacker struct {
	Statsdetail       *string  `json:"statsdetail"`
	Matchesplayed     *int     `json:"matchesplayed"`
	Roundsplayed      *int     `json:"roundsplayed"`
	Minutesplayed     *int     `json:"minutesplayed"`
	Matcheswon        *int     `json:"matcheswon"`
	Matcheslost       *int     `json:"matcheslost"`
	Roundswon         *int     `json:"roundswon"`
	Roundslost        *int     `json:"roundslost"`
	Kills             *int     `json:"kills"`
	Assists           *int     `json:"assists"`
	Death             *int     `json:"death"`
	Headshots         *int     `json:"headshots"`
	Meleekills        *int     `json:"meleekills"`
	Timealivepermatch *int     `json:"timealivepermatch"`
	Timedeadpermatch  *float64 `json:"timedeadpermatch"`
	Distanceperround  *float64 `json:"distanceperround"`
}

type Casual struct {
	//Type      *string    `json:"type"`
	Teamroles *TeamRoles `json:"teamRoles"`
}

type Defender struct {
	Statsdetail       *string  `json:"statsdetail"`
	Matchesplayed     *int     `json:"matchesplayed"`
	Roundsplayed      *int     `json:"roundsplayed"`
	Minutesplayed     *int     `json:"minutesplayed"`
	Matcheswon        *int     `json:"matcheswon"`
	Matcheslost       *int     `json:"matcheslost"`
	Roundswon         *int     `json:"roundswon"`
	Roundslost        *int     `json:"roundslost"`
	Kills             *int     `json:"kills"`
	Assists           *int     `json:"assists"`
	Death             *int     `json:"death"`
	Headshots         *int     `json:"headshots"`
	Meleekills        *int     `json:"meleekills"`
	Timealivepermatch *int     `json:"timealivepermatch"`
	Timedeadpermatch  *int     `json:"timedeadpermatch"`
	Distanceperround  *float64 `json:"distanceperround"`
}

type Maps struct {
	Unranked *Unranked `json:"unranked"`
	Ranked   *Ranked   `json:"ranked"`
	Casual   *Casual   `json:"casual"`
	All      *All      `json:"all"`
}

type Operators struct {
	Unranked *Unranked `json:"unranked"`
	Ranked   *Ranked   `json:"ranked"`
	Casual   *Casual   `json:"casual"`
	All      *All      `json:"all"`
}

type Ranked struct {
	Season      *int     `json:"season"`
	MaxMmr      *int     `json:"max_mmr"`
	SkillMean   *float64 `json:"skill_mean"`
	Deaths      *int     `json:"deaths"`
	NextRankMmr *int     `json:"nextrankmmr"`
	Rank        *int     `json:"rank"`
	MaxRank     *int     `json:"max_rank"`
	SkillStDev  *float64 `json:"skill_stdev"`
	Kills       *int     `json:"kills"`
	//Lastmatchskillstdevchange *int     `json:"lastmatchskillstdevchange"`
	LastMatchMmrChange *int `json:"last_match_mmr_change"`
	Abandons           *int `json:"abandons"`
	//Lastmatchskillmeanchange  *int     `json:"lastmatchskillmeanchange"`
	Mmr             *int `json:"mmr"`
	LastMatchResult *int `json:"last_match_result"`
	Wins            *int `json:"wins"`
	Losses          *int `json:"losses"`
}

type Summary struct {
	All      *All      `json:"all"`
	Unranked *Unranked `json:"unranked"`
	Ranked   *Ranked   `json:"ranked"`
	Casual   *Casual   `json:"casual"`
}

type TeamRoles struct {
	Defender []*Defender `json:"defender"`
	Attacker []*Attacker `json:"attacker"`
	All      []*All      `json:"all"`
}

type Unranked struct {
	Type      *string    `json:"type"`
	Teamroles *TeamRoles `json:"teamRoles"`
}
