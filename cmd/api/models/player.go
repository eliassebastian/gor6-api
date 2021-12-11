package models

type PlayerProfile struct {
	ProfileID      string `json:"profileId"`
	UserID         string `json:"userId"`
	PlatformType   string `json:"platformType"`
	IDOnPlatform   string `json:"idOnPlatform"`
	NameOnPlatform string `json:"nameOnPlatform"`
}

type PlayerIDModel struct {
	Profiles []PlayerProfile `json:"profiles"`
}

type PlayerGeneralStats struct {
}

type PlayerSeasonStats struct {
}

type PlayerWeaponStats struct {
}

type PlayerFullProfile struct {
	//PlayerProfile https://public-ubiservices.ubi.com/v1/profiles/stats?profileIds=ab1ff7ae-13e4-4a6a-9b03-317285f8057b&spaceId=5172a557-50b5-4665-b7db-e3f2e8c5041d&statNames=PPvPTimePlayed
	//PlayerAliases
	//GeneralStats - https://r6s-stats.ubisoft.com/v1/current/summary/ab1ff7ae-13e4-4a6a-9b03-317285f8057b?gameMode=all,ranked,unranked,casual&platform=PC&startDate=20210811&endDate=202112
	//PlayerSeasons - https://public-ubiservices.ubi.com/v1/spaces/5172a557-50b5-4665-b7db-e3f2e8c5041d/sandboxes/OSBOR_PC_LNCH_A/r6karma/player_skill_records?board_ids=pvp_ranked&season_ids=-1,-2,-3,-4,-5,-6,-7,-8,-9,-10,-11,-12,-13,-14,-15,-16,-17,-18,-19,-20,-21,-22,-23,-24&region_ids=ncsa&profile_ids=ab1ff7ae-13e4-4a6a-9b03-317285f8057b
	//Operators - https://r6s-stats.ubisoft.com/v1/current/operators/ab1ff7ae-13e4-4a6a-9b03-317285f8057b?gameMode=all,ranked,casual,unranked&platform=PC&teamRole=attacker,defender&startDate=20210811&endDate=20211209
	//Weapons - https://r6s-stats.ubisoft.com/v1/current/weapons/ab1ff7ae-13e4-4a6a-9b03-317285f8057b?gameMode=all,ranked,casual,unranked&platform=PC&teamRole=all&startDate=20210811&endDate=20211209
	//Trends - https://r6s-stats.ubisoft.com/v1/current/trend/ab1ff7ae-13e4-4a6a-9b03-317285f8057b?gameMode=all,ranked,casual,unranked&startDate=20210811&endDate=20211209&teamRole=all,attacker,defender&trendType=weeks
}
