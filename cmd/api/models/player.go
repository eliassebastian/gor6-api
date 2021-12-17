package models

import "time"

type PlayerFullProfile struct {
	ProfileID  string            `json:"profileID"`
	PlatformID string            `json:"platformID"`
	Platform   string            `json:"platform"`
	NickName   string            `json:"nickName"`
	LastUpdate time.Time         `json:"lastUpdate"`
	TimePlayed TimePlayed        `json:"timePlayed"`
	Level      Level             `json:"level"`
	Summary    SummaryGameModes  `json:"summary"`
	Weapons    WeaponsGameModes  `json:"weapons"`
	Operators  OperatorGameModes `json:"operators"`
	Maps       MapsGameModes     `json:"maps"`
	//PlayerProfile
	//Level https://public-ubiservices.ubi.com/v1/profiles/stats?profileIds=ab1ff7ae-13e4-4a6a-9b03-317285f8057b&spaceId=5172a557-50b5-4665-b7db-e3f2e8c5041d&statNames=PClearanceLevel
	//Playtime https://public-ubiservices.ubi.com/v1/profiles/stats?profileIds=ab1ff7ae-13e4-4a6a-9b03-317285f8057b&spaceId=5172a557-50b5-4665-b7db-e3f2e8c5041d&statNames=PPvPTimePlayed
	//PlayerAliases - TODO
	//GeneralStats - https://r6s-stats.ubisoft.com/v1/current/summary/ab1ff7ae-13e4-4a6a-9b03-317285f8057b?gameMode=all,ranked,unranked,casual&platform=PC&startDate=20210811&endDate=202112
	//PlayerSeasons - https://public-ubiservices.ubi.com/v1/spaces/5172a557-50b5-4665-b7db-e3f2e8c5041d/sandboxes/OSBOR_PC_LNCH_A/r6karma/player_skill_records?board_ids=pvp_ranked&season_ids=-1,-2,-3,-4,-5,-6,-7,-8,-9,-10,-11,-12,-13,-14,-15,-16,-17,-18,-19,-20,-21,-22,-23,-24&region_ids=ncsa&profile_ids=ab1ff7ae-13e4-4a6a-9b03-317285f8057b
	//Operators - https://r6s-stats.ubisoft.com/v1/current/operators/ab1ff7ae-13e4-4a6a-9b03-317285f8057b?gameMode=all,ranked,casual,unranked&platform=PC&teamRole=attacker,defender&startDate=20210811&endDate=20211209
	//Weapons - https://r6s-stats.ubisoft.com/v1/current/weapons/ab1ff7ae-13e4-4a6a-9b03-317285f8057b?gameMode=all,ranked,casual,unranked&platform=PC&teamRole=all&startDate=20210811&endDate=20211209
	//Trends - https://r6s-stats.ubisoft.com/v1/current/trend/ab1ff7ae-13e4-4a6a-9b03-317285f8057b?gameMode=all,ranked,casual,unranked&startDate=20210811&endDate=20211209&teamRole=all,attacker,defender&trendType=weeks
	//Maps - https://r6s-stats.ubisoft.com/v1/current/maps/ab1ff7ae-13e4-4a6a-9b03-317285f8057b?gameMode=all,ranked,casual,unranked&platform=PC&teamRole=all,attacker,defender&startDate=20210813&endDate=20211211
}
