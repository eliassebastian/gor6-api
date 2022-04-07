package model

import "time"

type SearchShards struct {
	Total      int `json:"total"`
	Successful int `json:"successful"`
	Skipped    int `json:"skipped"`
	Failed     int `json:"failed"`
}

type SearchTotal struct {
	Value    int    `json:"value"`
	Relation string `json:"relation"`
}

type SearchFields struct {
	LevelValue              []int       `json:"level.value"`
	AliasesName             []string    `json:"aliases.name"`
	TimePlayedLastModified  []time.Time `json:"timePlayed.lastModified"`
	ProfileID               []string    `json:"profileID"`
	NickName                []string    `json:"nickName"`
	TimePlayedValue         []int       `json:"timePlayed.value"`
	RankedCurrentSeasonRank []int       `json:"ranked.currentSeason.rank"`
}

type SearchResults struct {
	Index  string       `json:"_index"`
	Type   string       `json:"_type"`
	ID     string       `json:"_id"`
	Score  float32      `json:"_score"`
	Fields SearchFields `json:"fields"`
}

type SearchHits struct {
	Total    SearchTotal     `json:"total"`
	MaxScore float32         `json:"max_score"`
	Hits     []SearchResults `json:"hits"`
}

type SearchOutput struct {
	Took     int          `json:"took"`
	TimedOut bool         `json:"timed_out"`
	Shards   SearchShards `json:"_shards"`
	Hits     SearchHits   `json:"hits"`
}
