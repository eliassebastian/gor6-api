package models

type RankedModel struct {
	SeasonsPlayerSkillRecords []SeasonsPlayerSkillRecords `json:"seasons_player_skill_records"`
}

type RankedSeason struct {
	Season    int     `json:"season"`
	MaxMmr    float32 `json:"max_mmr"`
	SkillMean float32 `json:"skill_mean"`
	Deaths    int     `json:"deaths"`
	//ProfileID                 string    `json:"profile_id"`
	NextRankMmr float32 `json:"next_rank_mmr"`
	Rank        int     `json:"rank"`
	MaxRank     int     `json:"max_rank"`
	//BoardID                   string    `json:"board_id"`
	SkillStdev                float32 `json:"skill_stdev"`
	Kills                     int     `json:"kills"`
	LastMatchSkillStdevChange float32 `json:"last_match_skill_stdev_change"`
	//PastSeasonsWins           int       `json:"past_seasons_wins"`
	//UpdateTime                time.Time `json:"update_time"`
	LastMatchMmrChange float32 `json:"last_match_mmr_change"`
	Abandons           int     `json:"abandons"`
	//PastSeasonsLosses         int       `json:"past_seasons_losses"`
	//TopRankPosition          int `json:"top_rank_position"`
	LastMatchSkillMeanChange float32 `json:"last_match_skill_mean_change"`
	Mmr                      float32 `json:"mmr"`
	PreviousRankMmr          float32 `json:"previous_rank_mmr"`
	LastMatchResult          int     `json:"last_match_result"`
	PastSeasonsAbandons      int     `json:"past_seasons_abandons"`
	Wins                     int     `json:"wins"`
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
