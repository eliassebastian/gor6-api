package controllers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	model "github.com/eliassebastian/gor6-api/cmd/api/graph/models"
	"github.com/eliassebastian/gor6-api/internal/mongodb"
	"log"
	"net/http"
	"sync"
	"time"
)

func getSpaceId(p string) string {
	return model.SpaceIds[p]
}

func getPlatformURL(p string) string {
	return model.PlatformURLNames[p]
}

func genExpiration() string {
	//add 10 minutes to current date+time
	return time.Now().UTC().Add(10 * time.Minute).Format("2006-01-02T15:04:05.999Z07:00")
}

func getDate() string {
	year, month, day := time.Now().AddDate(0, 0, -1).Date()
	s := fmt.Sprintf("%v%02d%02d", year, int(month), day)
	fmt.Println("Date:", s)
	return s
}

func getHeader(p *sync.Map) http.Header {
	sd, ok := p.Load("session")
	if !ok {
		return nil
	}
	re, ok := sd.(map[string]string)
	if !ok {
		return nil
	}
	return http.Header{
		"authorization": []string{fmt.Sprintf("ubi_v1 t=%s", re["ticket"])},
		"Origin":        []string{"https://www.ubisoft.com"},
		"content-type":  []string{"application/json"},
		"user-agent":    []string{"node.js"},
		"ubi-appid":     []string{"3587dcbb-7f81-457c-9781-0e3f29f6f56a"},
		"ubi-sessionid": []string{re["sessionId"]},
		"expiration":    []string{genExpiration()},
	}
}

func fetchPlayerWeapons(ctx context.Context, wg *sync.WaitGroup, sm *sync.Map, hc *http.Client, player *model.Player, id, p string) {
	url := fmt.Sprintf("https://r6s-stats.ubisoft.com/v1/current/weapons/%s?gameMode=all&platform=%s&teamRole=all&startDate=20160101&endDate=%s", id, model.PlatformURLNames2[p], getDate())
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		log.Println("error fetching player weapon specific 1")
		wg.Done()
		return
	}

	req.Header = getHeader(sm)
	res, err := hc.Do(req)
	if err != nil {
		log.Println("error fetching player weapon 2", err)
		wg.Done()
		return
	}

	defer res.Body.Close()
	fmt.Println("Summary:", res)
	if res.StatusCode != 200 {
		fmt.Println("error fetching player weapon 3", res.Status)
		wg.Done()
		return
	}

	var wm model.WeaponsModel
	de := json.NewDecoder(res.Body).Decode(&wm)
	if de != nil {
		log.Println("error fetching player weapon 5", de)
		wg.Done()
		return
	}

	var wl model.WeaponsPlatform
	switch p {
	case "uplay":
		wl = wm.Platforms.Pc
	case "psn":
		wl = wm.Platforms.Ps4
	case "xbl":
		wl = wm.Platforms.Xbox
	default:
		log.Println("error fetching player level 5")
		wg.Done()
		return
	}

	player.Weapons = wl.GameModes
	fmt.Println("weapons done")
	wg.Done()
}

func fetchPlayerPlayTimeLevel(ctx context.Context, wg *sync.WaitGroup, sm *sync.Map, hc *http.Client, player *model.Player, id, p string) {
	url := fmt.Sprintf("https://public-ubiservices.ubi.com/v1/profiles/stats?profileIds=%s&spaceId=%s&statNames=PPvPTimePlayed,PClearanceLevel", id, getSpaceId(p))
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		log.Println("error fetching player level 1")
		wg.Done()
		return
	}

	req.Header = getHeader(sm)
	res, err := hc.Do(req)
	if err != nil {
		log.Println("error fetching player level 3")
		wg.Done()
		return
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Println("error fetching player level 4")
		wg.Done()
		return
	}

	var tm model.TimeAndLevelModel
	de := json.NewDecoder(res.Body).Decode(&tm)
	if de != nil {
		log.Println("error fetching player level 5")
		wg.Done()
		return
	}
	//log.Println(tm)
	player.Level = tm.Profiles[0].Stats.Level
	player.TimePlayed = tm.Profiles[0].Stats.TimePlayed
	wg.Done()
}

func fetchPlayerProfile(ctx context.Context, sm *sync.Map, hc *http.Client, n, p string) (*model.PlayerProfile, error) {
	url := fmt.Sprintf("https://public-ubiservices.ubi.com/v3/profiles?namesOnPlatform=%s&platformType=%s", n, p)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header = getHeader(sm)
	res, err := hc.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("error fetching profileId STATUS CODE %v // S: %s", res.StatusCode, res.Status))
	}

	var player model.PlayerProfiles
	de := json.NewDecoder(res.Body).Decode(&player)
	if de != nil {
		return nil, errors.New("error decoding player")
	}
	if len(player.Profiles) == 0 {
		return nil, nil
	}
	return &player.Profiles[0], nil
}

func SearchForPlayer(ctx context.Context, n, p string) (bool, interface{}, error) {
	//TODO: Switch to Automatic persisted queries for GraphQL

	//TODO: full text search of mongodb

	//TODO: Check whether user is found and index
	log.Println("Search For Player...")
	return false, nil, nil
}

func FetchNewPlayer(ctx context.Context, md *mongodb.MongoClient, sm *sync.Map, hc *http.Client, n, p string) (*model.Player, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	res, err := fetchPlayerProfile(ctx, sm, hc, n, p)
	if err != nil {
		return nil, err
	}

	if res == nil {
		return nil, errors.New("player not found")
	}

	wg := &sync.WaitGroup{}

	a := []*model.Alias{{
		Name: res.NameOnPlatform,
		Date: time.Now().UTC(),
	}}

	//ab := append(a, &model.Alias{
	//	Name: res.NameOnPlatform,
	//	Date: time.Now().UTC(),
	//})

	player := &model.Player{
		ID:         res.ProfileID,
		PlatformID: res.IDOnPlatform,
		Platform:   res.PlatformType,
		NickName:   res.NameOnPlatform,
		Aliases:    a,
		LastUpdate: time.Now().UTC(),
	}
	//if found, put in cache
	wg.Add(2)
	//return to request
	go fetchPlayerPlayTimeLevel(ctx, wg, sm, hc, player, res.ProfileID, p)
	//go pc.fetchPlayerSummary(ctx, wg, player, res.ProfileID, p)
	//go pc.fetchPlayerSummarySpecific(ctx, wg, player, res.ProfileID, p)
	//go pc.fetchPlayerRanked(ctx, wg, player, res.ProfileID, p)
	//go pc.fetchPlayerOperators(ctx, wg, player, res.ProfileID, p)
	go fetchPlayerWeapons(ctx, wg, sm, hc, player, res.ProfileID, p)
	//go fetchPlayerMaps(ctx, wg, player, res.ProfileID, p)
	wg.Wait()

	select {
	case <-ctx.Done():
		return nil, errors.New("fetch new player context cancelled")
	default:
		err := md.NewPlayer(ctx, p, player)
		if err != nil {
			return nil, err
		}
		log.Println("Fetch New Player Default select & inserted into mongodb")
		return player, nil
	}
}
