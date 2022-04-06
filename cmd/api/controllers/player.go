package controllers

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/eliassebastian/gor6-api/cmd/api/models"
	"github.com/eliassebastian/gor6-api/cmd/api/response"
	"github.com/eliassebastian/gor6-api/internal/elastic"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"sync"
	"time"
)

type PlayerController struct {
	ec *elastic.ESClient
	//mc *mongodb.MongoClient
	sm *sync.Map
	hc *http.Client
}

type testPayload struct {
	Player   string
	Platform string
}

func NewPlayerController(c *elastic.ESClient, tlsc *tls.Config, p *sync.Map) *PlayerController {
	return &PlayerController{
		ec: c,
		//mc: m,
		sm: p,
		hc: &http.Client{
			Timeout: 10 * time.Second,
			//Transport: &http.Transport{
			//	ForceAttemptHTTP2: true,
			//	TLSClientConfig:   tlsc,
			//},
		},
	}
}

func (pc *PlayerController) getHeader() http.Header {
	sd, ok := pc.sm.Load("session")
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

func getDate(day int) string {
	year, month, day := time.Now().AddDate(0, 0, day).Date()
	s := fmt.Sprintf("%v%02d%02d", year, int(month), day)
	fmt.Println("Date:", s)
	return s
}

func (pc *PlayerController) searchForPlayer(ctx context.Context, n, p string) (bool, interface{}, error) {
	//TODO: Redis Cache Player?

	//TODO: search elastic search

	//Not Indexed? Fetch

	return false, nil, nil
}

func (pc *PlayerController) fetchPlayerTrends(ctx context.Context, player *model.Player, id, p string) error {
	url := fmt.Sprintf("https://r6s-stats.ubisoft.com/v1/current/trend/%s?gameMode=all,ranked,casual,unranked&startDate=%s&endDate=%s&teamRole=all,attacker,defender&trendType=days", id, getDate(-7), getDate(-1))
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	req.Header = pc.getHeader()
	res, err := pc.hc.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		return errors.New("(trends) response status code is not 200")
	}

	var m model.TrendsModel
	de := json.NewDecoder(res.Body).Decode(&m)
	if de != nil {
		return de
	}

	var wl model.TrendsPlatform
	switch p {
	case "uplay":
		wl = m.Platforms.Pc
	case "psn":
		wl = m.Platforms.Ps4
	case "xbl":
		wl = m.Platforms.Xbox
	default:
		return errors.New("error fetching trends platform")
	}

	player.Trends = &wl.GameModes
	return nil
}

func (pc *PlayerController) fetchPlayerMaps(ctx context.Context, player *model.Player, id, p string) error {
	url := fmt.Sprintf("https://r6s-stats.ubisoft.com/v1/current/maps/%s?gameMode=all,ranked,casual,unranked&platform=%s&teamRole=all,attacker,defender&startDate=20160101&endDate=%s", id, model.PlatformURLNames2[p], getDate(-1))
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	req.Header = pc.getHeader()
	res, err := pc.hc.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		return errors.New("(map) response status code is not 200")
	}

	var m model.MapsModel
	de := json.NewDecoder(res.Body).Decode(&m)
	if de != nil {
		return de
	}

	var wl model.MapsPlatform
	switch p {
	case "uplay":
		wl = m.Platforms.Pc
	case "psn":
		wl = m.Platforms.Ps4
	case "xbl":
		wl = m.Platforms.Xbox
	default:
		return errors.New("fetching map platform error")
	}

	player.Maps = &wl.GameModes
	return nil
}

func (pc *PlayerController) fetchPlayerWeapons(ctx context.Context, player *model.Player, id, p string) error {
	url := fmt.Sprintf("https://r6s-stats.ubisoft.com/v1/current/weapons/%s?gameMode=all&platform=%s&teamRole=all&startDate=20160101&endDate=%s", id, model.PlatformURLNames2[p], getDate(-1))
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	req.Header = pc.getHeader()
	res, err := pc.hc.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		return errors.New("response status code is not 200")
	}

	var wm model.WeaponsModel
	de := json.NewDecoder(res.Body).Decode(&wm)
	if de != nil {
		return de
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
		return errors.New("weapons platform fetch error")
	}

	player.Weapons = &wl.GameModes
	return nil
}

func (pc *PlayerController) fetchPlayerOperators(ctx context.Context, player *model.Player, id, p string) error {
	url := fmt.Sprintf("https://r6s-stats.ubisoft.com/v1/current/operators/%s?gameMode=all,ranked,casual,unranked&platform=%s&teamRole=attacker,defender&startDate=20160101&endDate=%s", id, model.PlatformURLNames2[p], getDate(-1))
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	req.Header = pc.getHeader()
	res, err := pc.hc.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		return errors.New("response status code is not 200")
	}

	var m model.OperatorModel
	de := json.NewDecoder(res.Body).Decode(&m)
	if de != nil {
		return de
	}

	var wl model.OperatorPlatform
	switch p {
	case "uplay":
		wl = m.Platforms.Pc
	case "psn":
		wl = m.Platforms.Ps4
	case "xbl":
		wl = m.Platforms.Xbox
	default:
		return errors.New("operator platform error")
	}

	player.Operators = &wl.GameModes
	return nil
}

func (pc *PlayerController) fetchPlayerRanked(ctx context.Context, player *model.Player, id, p string) error {
	url := fmt.Sprintf("https://public-ubiservices.ubi.com/v1/spaces/%s/sandboxes/%s/r6karma/player_skill_records?board_ids=pvp_ranked&season_ids=-1,-2,-3,-4,-5,-6,-7,-8,-9,-10,-11,-12,-13,-14,-15,-16,-17,-18,-19&region_ids=ncsa&profile_ids=%s", getSpaceId(p), getPlatformURL(p), id)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	req.Header = pc.getHeader()
	res, err := pc.hc.Do(req)
	fmt.Println(res.Request.Header)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		return errors.New("response status code is not 200")
	}

	var m model.RankedModel
	de := json.NewDecoder(res.Body).Decode(&m)
	if de != nil {
		return de
	}

	var seasons []model.RankedSeason
	for _, season := range m.SeasonsPlayerSkillRecords {
		seasons = append(seasons, season.RegionsPlayerSkillRecords[0].BoardsPlayerSkillRecords[0].Seasons[0])
	}

	player.Ranked = &seasons
	return nil
}

//func (pc *PlayerController) fetchPlayerSummarySpecific(ctx context.Context, wg *sync.WaitGroup, player *model.Player, id, p string) {
//	fmt.Println(id, p)
//	url := fmt.Sprintf("https://r6s-stats.ubisoft.com/v1/current/summary/%s?gameMode=all,ranked,unranked,casual&platform=%s&startDate=20210811&endDate=20211214", id, model.PlatformURLNames2[p])
//	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
//	if err != nil {
//		log.Println("error fetching player summary specific 1")
//		wg.Done()
//		return
//	}
//	req.Header = pc.getHeader()
//	res, err := pc.hc.Do(req)
//	if err != nil {
//		fmt.Println("error fetching player summary specific 2", err)
//		wg.Done()
//		return
//	}
//
//	fmt.Print("Specific:", res)
//	wg.Done()
//}

func (pc *PlayerController) fetchPlayerSummary(ctx context.Context, player *model.Player, id, p string) error {
	url := fmt.Sprintf("https://r6s-stats.ubisoft.com/v1/seasonal/summary/%s?gameMode=all,ranked,casual,unranked&platform=%s", id, model.PlatformURLNames2[p])
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	req.Header = pc.getHeader()
	res, err := pc.hc.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		return errors.New("response status code is not 200")
	}

	var sm model.SummaryModel
	de := json.NewDecoder(res.Body).Decode(&sm)
	if de != nil {
		return de
	}

	var sl model.SummaryPlatform
	switch p {
	case "uplay":
		sl = sm.Platforms.Pc
	case "psn":
		sl = sm.Platforms.Ps4
	case "xbl":
		sl = sm.Platforms.Xbox
	default:
		return errors.New("error fetching player level 5")
	}

	player.Summary = &sl.SGameModes
	return nil
}

func (pc *PlayerController) fetchPlayerPlayTimeLevel(ctx context.Context, player *model.Player, id, p string) error {
	url := fmt.Sprintf("https://public-ubiservices.ubi.com/v1/profiles/stats?profileIds=%s&spaceId=%s&statNames=PPvPTimePlayed,PClearanceLevel", id, getSpaceId(p))
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		log.Println("error fetching player level 1")
		return err
	}

	req.Header = pc.getHeader()
	res, err := pc.hc.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		return errors.New("response status code is not 200")
	}

	var tm model.TimeAndLevelModel
	de := json.NewDecoder(res.Body).Decode(&tm)
	if de != nil {
		return de
	}

	player.Level = &tm.Profiles[0].Stats.Level
	player.TimePlayed = &tm.Profiles[0].Stats.TimePlayed
	return nil
}

func (pc *PlayerController) fetchPlayerProfile(ctx context.Context, n, p string) (*model.PlayerProfile, error) {
	url := fmt.Sprintf("https://public-ubiservices.ubi.com/v3/profiles?namesOnPlatform=%s&platformType=%s", n, p)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header = pc.getHeader()
	res, err := pc.hc.Do(req)
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

func (pc *PlayerController) fetchNewPlayer(ctx context.Context, n, p string) (*model.Player, error) {
	res, err := pc.fetchPlayerProfile(ctx, n, p)
	if res == nil || err != nil {
		return nil, err
	}
	//set up aliases
	a := &[]model.Alias{{
		Name: res.NameOnPlatform,
		Date: time.Now().UTC(),
	}}

	player := &model.Player{
		ID:         res.ProfileID,
		PlatformID: res.IDOnPlatform,
		Platform:   res.PlatformType,
		NickName:   res.NameOnPlatform,
		Aliases:    a,
		LastUpdate: time.Now().UTC(),
	}

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		err := pc.fetchPlayerPlayTimeLevel(ctx, player, res.ProfileID, p)
		return err
	})

	g.Go(func() error {
		err := pc.fetchPlayerSummary(ctx, player, res.ProfileID, p)
		return err
	})

	g.Go(func() error {
		err := pc.fetchPlayerRanked(ctx, player, res.ProfileID, p)
		return err
	})

	g.Go(func() error {
		err := pc.fetchPlayerOperators(ctx, player, res.ProfileID, p)
		return err
	})

	g.Go(func() error {
		err := pc.fetchPlayerWeapons(ctx, player, res.ProfileID, p)
		return err
	})

	g.Go(func() error {
		err := pc.fetchPlayerMaps(ctx, player, res.ProfileID, p)
		return err
	})

	g.Go(func() error {
		err := pc.fetchPlayerTrends(ctx, player, res.ProfileID, p)
		return err
	})

	if err := g.Wait(); err != nil {
		return nil, err
	}

	return player, nil
}

func (pc *PlayerController) Test(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()

	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	var p testPayload
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		response.ErrorJSON(w, err)
		return
	}

	defer r.Body.Close()
	res, err := pc.fetchNewPlayer(ctx, p.Player, p.Platform)
	if err != nil {
		response.ErrorJSON(w, err)
		return
	}

	ie := pc.ec.IndexPlayer(ctx, res, p.Platform)
	if ie != nil {
		response.ErrorJSON(w, ie)
		return
	}

	response.SuccessJSON(w, startTime, res)
}
