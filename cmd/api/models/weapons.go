package models

type WeaponsModel struct {
	ProfileID string           `json:"profileId"`
	Platforms WeaponsPlatforms `json:"platforms"`
}

type Weapons struct {
	WeaponName   string `json:"weaponName"`
	RoundsPlayed int    `json:"roundsPlayed"`
	RoundsWon    int    `json:"roundsWon"`
	RoundsLost   int    `json:"roundsLost"`
	Kills        int    `json:"kills"`
	Headshots    int    `json:"headshots"`
	//HeadshotAccuracy    float32 `json:"headshotAccuracy"`
	//RoundsWithAKill     float32 `json:"roundsWithAKill"`
	RoundsWithMultiKill float32 `json:"roundsWithMultiKill"`
}

type WeaponTypes struct {
	WeaponType string    `json:"weaponType"`
	Weapons    []Weapons `json:"weapons"`
}

type SecondaryWeapons struct {
	WeaponTypes []WeaponTypes `json:"weaponTypes"`
}

type PrimaryWeapons struct {
	WeaponTypes []WeaponTypes `json:"weaponTypes"`
}

type WeaponSlots struct {
	SecondaryWeapons SecondaryWeapons `json:"secondaryWeapons"`
	PrimaryWeapons   PrimaryWeapons   `json:"primaryWeapons"`
}

type WeaponsAll struct {
	WeaponSlots WeaponSlots `json:"weaponSlots"`
}

type WeaponsTeamRoles struct {
	All WeaponsAll `json:"all"`
}

type WeaponsGameMode struct {
	Type      string           `json:"type"`
	TeamRoles WeaponsTeamRoles `json:"teamRoles"`
}

type WeaponsGameModes struct {
	All      WeaponsGameMode `json:"all"`
	Casual   WeaponsGameMode `json:"casual"`
	Ranked   WeaponsGameMode `json:"ranked"`
	Unranked WeaponsGameMode `json:"unranked"`
}

type WeaponsPlatform struct {
	GameModes WeaponsGameModes `json:"gameModes"`
}

type WeaponsPlatforms struct {
	Pc   WeaponsPlatform `json:"PC,omitempty"`
	Xbox WeaponsPlatform `json:"XBOX,omitempty"`
	Ps4  WeaponsPlatform `json:"PS4,omitempty"`
}
