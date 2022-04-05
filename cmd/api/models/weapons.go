package model

type WeaponsModel struct {
	ProfileID string           `json:"profileId"`
	Platforms WeaponsPlatforms `json:"platforms"`
}

type Weapon struct {
	WeaponName          *string  `json:"weaponName"`
	RoundsPlayed        *int     `json:"roundsPlayed"`
	RoundsWon           *int     `json:"roundsWon"`
	RoundsLost          *int     `json:"roundsLost"`
	Kills               *int     `json:"kills"`
	Headshots           *int     `json:"headshots"`
	RoundsWithMultiKill *float32 `json:"roundsWithMultiKill" bson:"roundsWithMultiKill_F32"`
}

type WeaponTypes struct {
	Weapons []*Weapon `json:"weapons"`
}

type Weapons struct {
	WeaponTypes []*WeaponTypes `json:"weaponTypes"`
}

type WeaponSlots struct {
	SecondaryWeapons *Weapons `json:"secondaryWeapons"`
	PrimaryWeapons   *Weapons `json:"primaryWeapons"`
}

type WeaponsAll struct {
	WeaponSlots *WeaponSlots `json:"weaponSlots"`
}

type WeaponsTeamRoles struct {
	All *WeaponsAll `json:"all"`
}

type WeaponsGameMode struct {
	TeamRoles *WeaponsTeamRoles `json:"teamRoles"`
}

type WeaponsGameModes struct {
	All      *WeaponsGameMode `json:"all"`
	Casual   *WeaponsGameMode `json:"casual"`
	Ranked   *WeaponsGameMode `json:"ranked"`
	Unranked *WeaponsGameMode `json:"unranked"`
}

type WeaponsPlatform struct {
	GameModes WeaponsGameModes `json:"gameModes"`
}

type WeaponsPlatforms struct {
	Pc   WeaponsPlatform `json:"PC,omitempty"`
	Xbox WeaponsPlatform `json:"XONE,omitempty"`
	Ps4  WeaponsPlatform `json:"PS4,omitempty"`
}
