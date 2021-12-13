package controllers

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/eliassebastian/gor6-api/cmd/api/models"
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
	//appid? 3587dcbb-7f81-457c-9781-0e3f29f6f56a
	return http.Header{
		"Authorization": []string{fmt.Sprintf("Ubi_v1 t=%s", re["ticket"])},
		"Ubi-AppId":     []string{"39baebad-39e5-4552-8c25-2c9b919064e2"},
		"Ubi-SessionId": []string{re["sessionId"]},
		"Connection":    []string{"keep-alive"},
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

//fetch level

//fetch playtime

func (pc *PlayerController) fetchGeneralStats(ctx context.Context, n, p string) (interface{}, error) {
	var stats = strings.Join([]string{"PPvPtimeplayed", "PClearanceLevel", "PPvPmatchplayed", "PPvPmatchwon", "PPvPmatchlost", "PPvPkills", "PPvPdeath"}[:], ",")
	//var stats = strings.Join([]string{"PPvPtimeplayed", "PClearanceLevel"}[:], ",")
	url := fmt.Sprintf("https://public-ubiservices.ubi.com/v1/profiles/stats?profileIds=ab1ff7ae-13e4-4a6a-9b03-317285f8057b&spaceId=%s&statNames=%s", getSpaceId(p), stats)
	//url2 := fmt.Sprintf("https://public-ubiservices.ubi.com/v1/spaces/%s/sandboxes/%s/playerstats2/statistics?populations=%s&statistics=%s", getSpaceId(p), getPlatformURL(p), "ab1ff7ae-13e4-4a6a-9b03-317285f8057b", "PPvPtimeplayed")
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)

	if err != nil {
		return nil, err
	}

	sd, ok := pc.sm.Load("session")
	if !ok {
		return nil, nil
	}

	re, ok := sd.(map[string]string)
	if !ok {
		return nil, nil
	}

	//req.Proto = "h2"
	//req.Method = http.MethodGet
	req.Header = http.Header{
		//experimenting with same headers as on ubi website for same request
		//":method": []string{"GET"},
		//":scheme":         []string{"https"},
		//":authority":      []string{"public-ubiservices.ubi.com"},
		//":path":           []string{fmt.Sprintf("/v1/profiles/stats?profileIds=ab1ff7ae-13e4-4a6a-9b03-317285f8057b&spaceId=5172a557-50b5-4665-b7db-e3f2e8c5041d&statNames=PPvPTimePlayed")},
		"Accept":          []string{"*/*"},
		"Accept-Language": []string{"en-GB,en;q=0.9"},
		"Accept-Encoding": []string{"gzip", "deflate", "br"},
		"Authorization":   []string{fmt.Sprintf("ubi_v1 t=%s", re["ticket"])},
		"Host":            []string{"public-ubiservices.ubi.com"},
		"Origin":          []string{"https://www.ubisoft.com"},
		"User-Agent":      []string{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.1 Safari/605.1.15"},
		"Ubi-AppId":       []string{"3587dcbb-7f81-457c-9781-0e3f29f6f56a"},
		"Ubi-SessionId":   []string{re["sessionId"]},
		//"content-type":    []string{"application/json"},
		"Connection": []string{"keep-alive"},
		//"Referer":         []string{"https://www.ubisoft.com/"},
		"expiration": []string{genExpiration()},
	}
	//fmt.Println(req)
	//fmt.Println(req.Header)
	log.Print("HC DO")
	//res, err := pc.hc.RoundTrip(req)
	res, err := pc.hc.Do(req)
	if err != nil {
		return nil, err
	}
	//fmt.Println(res.Request, res.ProtoMajor, err)
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

/*
func (pc *PlayerController) fetchProfileId(ctx context.Context, n, p string) (*models.PlayerProfile, error) {
	url := fmt.Sprintf("https://public-ubiservices.ubi.com/v3/profiles?namesOnPlatform=%s&platformType=%s", n, p)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header = pc.getHeader()
	res, _ := pc.hc.Do(req)
	if res.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("error fetching profileId STATUS CODE %v // S: %s", res.StatusCode, res.Status))
	}

	var player models.PlayerIDModel
	de := json.NewDecoder(res.Body).Decode(&player)
	if de != nil {
		return nil, errors.New("error decoding player")
	}
	res.Body.Close()
	return &player.Profiles[0], nil
}*/

func (pc *PlayerController) fetchPlayer(ctx context.Context, n, p string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second) //change to deadline?
	defer cancel()
	//search elastic (check direct name and alias)

	//if found, put in cache
	//return to request

	//if not found check for profileId
	//profileId, errf := pc.fetchProfileId(ctx, n, p)
	//if errf != nil {
	//	log.Println(errf)
	//	return []byte("empty"), errf
	//}

	//player stats
	res, err := pc.fetchGeneralStats(ctx, n, p)
	if err != nil {
		log.Println(err)
		return []byte(""), err
	}
	log.Println(res)
	//seasonal stats
	//operator stats
	//weapon stats
	return []byte("success"), nil
	//return []byte(profileId.ProfileID), nil
}

func (pc *PlayerController) Test(w http.ResponseWriter, r *http.Request) {
	platform := strings.ToLower(chi.URLParam(r, "platform"))
	player := strings.ToLower(chi.URLParam(r, "player"))
	res, err := pc.fetchPlayer(r.Context(), player, platform)
	if err != nil {
		log.Println(err)
		w.Write([]byte("check logs error"))
		return
	}
	w.Write(res)
}

func (pc *PlayerController) GetPlayers(w http.ResponseWriter, r *http.Request) {
	//platform := chi.URLParam(r, "platform")
	//player := chi.URLParam(r, "player")
}

func (pc *PlayerController) Player(w http.ResponseWriter, r *http.Request) {

}
