package models

type RankedModel struct {
	SeasonsPlayerSkillRecords []SeasonsPlayerSkillRecords `json:"seasons_player_skill_records"`
}

type RankedSeason struct {
	Season    int     `json:"season"`
	MaxMmr    float32 `json:"max_mmr" bson:"maxMmr_F32"`
	SkillMean float32 `json:"skill_mean" bson:"skillMean_F32"`
	Deaths    int     `json:"deaths"`
	//ProfileID                 string    `json:"profile_id"`
	Rank    int `json:"rank"`
	MaxRank int `json:"max_rank"`
	//BoardID                   string    `json:"board_id"`
	SkillStdev float32 `json:"skill_stdev" bson:"skillStDev_F32"`
	Kills      int     `json:"kills"`
	//PastSeasonsWins           int       `json:"past_seasons_wins"`
	//UpdateTime                time.Time `json:"update_time"`
	LastMatchMmrChange float32 `json:"last_match_mmr_change" bson:"lastMatchMmrChange_F32"`
	Abandons           int     `json:"abandons"`
	//PastSeasonsLosses         int       `json:"past_seasons_losses"`
	//TopRankPosition          int `json:"top_rank_position"`
	Mmr             float32 `json:"mmr" bson:"mmr_F32"`
	LastMatchResult int     `json:"last_match_result"`
	Wins            int     `json:"wins"`
	//Region                    string    `json:"region"`
	Losses int `json:"losses"`
}

type RankedSeasons struct {
	//BoardID             string                `json:"board_id"`
	Seasons []RankedSeason `json:"players_skill_records"`
}

type RegionsPlayerSkillRecords struct {
	//RegionID                 string                     `json:"region_id"`
	BoardsPlayerSkillRecords []RankedSeasons `json:"boards_player_skill_records"`
}

type SeasonsPlayerSkillRecords struct {
	SeasonID                  int                         `json:"season_id"`
	RegionsPlayerSkillRecords []RegionsPlayerSkillRecords `json:"regions_player_skill_records"`
}
