package model

type RankedSeason struct {
	Season             int     `json:"season"`
	MaxMmr             float32 `json:"max_mmr"`
	SkillMean          float32 `json:"skill_mean"`
	Deaths             int     `json:"deaths"`
	Rank               int     `json:"rank"`
	MaxRank            int     `json:"max_rank"`
	SkillStdev         float32 `json:"skill_stdev"`
	Kills              int     `json:"kills"`
	LastMatchMmrChange float32 `json:"last_match_mmr_change"`
	Abandons           int     `json:"abandons"`
	Mmr                float32 `json:"mmr"`
	LastMatchResult    int     `json:"last_match_result"`
	Wins               int     `json:"wins"`
	Losses             int     `json:"losses"`
}

type RankedSeasons struct {
	Seasons []RankedSeason `json:"players_skill_records"`
}

type RegionsPlayerSkillRecords struct {
	BoardsPlayerSkillRecords []RankedSeasons `json:"boards_player_skill_records"`
}

type SeasonsPlayerSkillRecords struct {
	SeasonID                  int                         `json:"season_id"`
	RegionsPlayerSkillRecords []RegionsPlayerSkillRecords `json:"regions_player_skill_records"`
}

type RankedModel struct {
	SeasonsPlayerSkillRecords []SeasonsPlayerSkillRecords `json:"seasons_player_skill_records"`
}
