package controllers

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/eliassebastian/gor6-api/cmd/api/models"
	"github.com/eliassebastian/gor6-api/cmd/api/util"
	"github.com/eliassebastian/gor6-api/internal/elastic"
	"github.com/eliassebastian/gor6-api/internal/mongodb"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
)

type PlayerController struct {
	ec *elastic.ESClient
	mc *mongodb.MongoClient
	sm *sync.Map
	hc *http.Client
}

func NewPlayerController(c *elastic.ESClient, m *mongodb.MongoClient, tlsc *tls.Config, p *sync.Map) *PlayerController {
	return &PlayerController{
		ec: c,
		mc: m,
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
		"Accept": []string{"*/*"},
		//"Accept-Language": []string{"en-GB,en;q=0.9"},
		//"Accept-Encoding": []string{"gzip", "deflate", "br"},
		"Authorization": []string{fmt.Sprintf("ubi_v1 t=%s", re["ticket"])},
		//"Host":            []string{"public-ubiservices.ubi.com"},
		//"Origin":          []string{"https://www.ubisoft.com"},
		"User-Agent":    []string{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.1 Safari/605.1.15"},
		"Ubi-AppId":     []string{"3587dcbb-7f81-457c-9781-0e3f29f6f56a"},
		"Ubi-SessionId": []string{re["sessionId"]},
		"Connection":    []string{"keep-alive"},
		"expiration":    []string{genExpiration()},
	}
}

func getSpaceId(p string) string {
	return models.SpaceIds[p]
}

func getPlatformURL(p string) string {
	return models.PlatformURLNames[p]
}

func genExpiration() string {
	return (time.Now().Add(1 * time.Hour)).String()
}

func (pc *PlayerController) searchForPlayer(ctx context.Context, n, p string) (bool, interface{}, error) {
	//TODO: Redis Cache Player?

	//TODO: search elastic search

	//Not Indexed? Fetch

	return false, nil, nil
}

func (pc *PlayerController) fetchGeneralStats(ctx context.Context, n, p string) (interface{}, error) {
	var stats = strings.Join([]string{"PPvPtimeplayed", "PClearanceLevel", "PPvPmatchplayed", "PPvPmatchwon", "PPvPmatchlost", "PPvPkills", "PPvPdeath"}[:], ",")
	//var stats = strings.Join([]string{"PPvPtimeplayed", "PClearanceLevel"}[:], ",")
	url := fmt.Sprintf("https://public-ubiservices.ubi.com/v1/profiles/stats?profileIds=ab1ff7ae-13e4-4a6a-9b03-317285f8057b&spaceId=%s&statNames=%s", getSpaceId(p), stats)
	//url2 := fmt.Sprintf("https://public-ubiservices.ubi.com/v1/spaces/%s/sandboxes/%s/playerstats2/statistics?populations=%s&statistics=%s", getSpaceId(p), getPlatformURL(p), "ab1ff7ae-13e4-4a6a-9b03-317285f8057b", "PPvPtimeplayed")
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)

	if err != nil {
		return nil, err
	}

	req.Header = pc.getHeader()
	res, err := pc.hc.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		log.Println("Error with Client Request", res.Status)
		return nil, err
	}
	var output map[string]interface{}
	json.NewDecoder(res.Body).Decode(&output)
	fmt.Println(output)
	res.Body.Close()
	return nil, nil
}

func (pc *PlayerController) fetchPlayerPlayTimeLevel(ctx context.Context, wg *sync.WaitGroup, player *models.PlayerFullProfile, id, p string) {
	url := fmt.Sprintf("https://public-ubiservices.ubi.com/v1/profiles/stats?profileIds=%s&spaceId=%s&statNames=PPvPTimePlayed,PClearanceLevel", id, getSpaceId(p))
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		log.Println("error fetching player level 1")
		wg.Done()
		return
	}

	req.Header = pc.getHeader()
	res, _ := pc.hc.Do(req)
	if res.StatusCode != 200 {
		log.Println("error fetching player level 2")
		wg.Done()
		return
	}
	defer res.Body.Close()

	log.Println()

	var tm models.TimeAndLevelModel
	de := json.NewDecoder(res.Body).Decode(&tm)
	if de != nil {
		log.Println("error fetching player level 3")
		wg.Done()
		return
	}
	log.Println(tm)
	player.Level = tm.Profiles[0].StatsO.LevelO
	player.TimePlayed = tm.Profiles[0].StatsO.TimePlayedO
	wg.Done()
}

func (pc *PlayerController) fetchPlayerProfile(ctx context.Context, n, p string) (*models.PlayerProfile, error) {
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

	var player models.PlayerIDModel
	de := json.NewDecoder(res.Body).Decode(&player)
	if de != nil {
		return nil, errors.New("error decoding player")
	}
	return &player.Profiles[0], nil
}

func (pc *PlayerController) fetchNewPlayer(ctx context.Context, n, p string) (interface{}, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second) //change to deadline?
	defer cancel()

	res, err := pc.fetchPlayerProfile(ctx, n, p)
	if err != nil {
		return nil, err
	}

	//TODO create channel for communication between goroutines
	wg := &sync.WaitGroup{}
	player := &models.PlayerFullProfile{
		ProfileID: res.IDOnPlatform,
		Platform:  res.PlatformType,
		NickName:  res.NameOnPlatform,
	}
	//if found, put in cache
	wg.Add(1)
	//return to request
	go pc.fetchPlayerPlayTimeLevel(ctx, wg, player, res.IDOnPlatform, p)
	//go pc.fetchPlayerLevel(ctx, wg, player, res.IDOnPlatform, p)
	//go pc.fetchPlayerPlaytime(ctx, wg, player, res.IDOnPlatform, p)

	wg.Wait()

	select {
	case <-ctx.Done():
		return nil, errors.New("fetch new player context cancelled")
	default:
		log.Println("Fetch New Player Default select")
		return player, nil
	}
}

func (pc *PlayerController) Test(w http.ResponseWriter, r *http.Request) {
	platform := strings.ToLower(chi.URLParam(r, "platform"))
	player := strings.ToLower(chi.URLParam(r, "player"))
	startTime := time.Now()

	s, _, err := pc.searchForPlayer(r.Context(), player, platform)
	if s {
		if err != nil {
			//util.ErrorJSON()
			w.Write([]byte("found player"))
			return
		}
		//util.ReturnJSON()
		w.Write([]byte("found player"))
		return
	}

	res, err := pc.fetchNewPlayer(r.Context(), player, platform)
	if err != nil {
		util.ErrorJSON(w, startTime, http.StatusNotFound, err.Error())
		return
	}

	util.ReturnJSON(w, startTime, http.StatusOK, "OK", res)
}
