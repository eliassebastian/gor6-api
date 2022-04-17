package controllers

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	model "github.com/eliassebastian/gor6-api/cmd/api/models"
	"github.com/eliassebastian/gor6-api/cmd/api/response"
	"github.com/eliassebastian/gor6-api/internal/cache"
	"github.com/eliassebastian/gor6-api/internal/elastic"
	"github.com/vmihailenco/msgpack/v5"
	"log"
	"net/http"
	"sync"
	"time"
)

type SearchController struct {
	ec *elastic.ESClient
	ic *cache.IndexCache
	sm *sync.Map
	hc *http.Client
}

func NewSearchController(c *elastic.ESClient, i *cache.IndexCache, tlsc *tls.Config, p *sync.Map) *SearchController {
	return &SearchController{
		ec: c,
		ic: i,
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

func (sc *SearchController) getHeader() http.Header {
	sd, ok := sc.sm.Load("session")
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

func (sc *SearchController) fetchPlayerProfile(ctx context.Context, n, p string) (*model.PlayerProfile, error) {
	url := fmt.Sprintf("https://public-ubiservices.ubi.com/v3/profiles?namesOnPlatform=%s&platformType=%s", n, p)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header = sc.getHeader()
	res, err := sc.hc.Do(req)
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

type searchPayload struct {
	Player   string
	Platform string
}

//curl -H "Content-Type: application/json" -X POST -d '{"player":"Kanzen","platform":"uplay"}' https://localhost:8090/test
func (sc *SearchController) SearchPlayer(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()

	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	var sp searchPayload
	err := json.NewDecoder(r.Body).Decode(&sp)
	if err != nil {
		log.Println("msgpack error #1")
		response.ErrorJSON(w, err)
		return
	}

	defer r.Body.Close()

	n, res, err := sc.ec.SearchPlayer(ctx, sp.Player, sp.Platform)
	if err != nil {
		response.ErrorJSON(w, err)
		return
	}

	//experiment with maxscore to define boundary for fetching profile from ubisoft
	log.Println("Search Controller Search Player Func - MaxScore:", n, "for Player:", sp.Player)
	if n < 1.0 || res.Total.Value == 0 {
		p, err := sc.fetchPlayerProfile(ctx, sp.Player, sp.Platform)
		if err != nil {
			response.ErrorJSON(w, err)
			return
		}
		//if profile is found add to existing slice
		if p != nil {
			//iterate through existing results to check if profile is already included
			for _, result := range res.Hits {
				if result.Fields.ProfileID[0] == p.ProfileID {
					response.SuccessJSON(w, startTime, res.Hits)
					return
				}
			}

			//TODO: cache profileID for index
			b, err := msgpack.Marshal(p)
			if err != nil {
				log.Println("msgpack error #2")
				response.ErrorJSON(w, err)
				return
			}

			err = sc.ic.DB.Set(ctx, p.ProfileID, b, 30*time.Minute).Err()
			if err != nil {
				response.ErrorJSON(w, err)
				return
			}

			sf := []model.SearchResults{
				{
					Index: sp.Platform,
					ID:    p.ProfileID,
					Score: -1,
					Fields: model.SearchFields{
						ProfileID: []string{p.ProfileID},
						NickName:  []string{p.NameOnPlatform},
						Platform:  []string{p.PlatformType},
					},
				},
			}

			tot := res.Total.Value + 1
			if tot <= cap(res.Hits) {
				s2 := res.Hits[:tot]
				copy(s2[1:], res.Hits[0:])
				copy(s2[0:], sf)

				response.SuccessJSON(w, startTime, &s2)
				return
			}

			s2 := make([]model.SearchResults, tot)
			copy(s2, res.Hits[:0])
			copy(s2[0:], sf)
			copy(s2[1:], res.Hits[0:])

			response.SuccessJSON(w, startTime, &s2)
			return
		}
	}

	response.SuccessJSON(w, startTime, &res.Hits)
}
