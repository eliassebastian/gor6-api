package controllers

import (
	"context"
	"encoding/json"
	"errors"
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

func NewPlayerController(c *elastic.ESClient, m *mongodb.MongoClient, p *sync.Map) *PlayerController {
	return &PlayerController{
		ec: c,
		mc: m,
		sm: p,
		hc: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (pc *PlayerController) fetchProfileId(ctx context.Context, n, p string) (string, error) {
	log.Println("FetchProfileId", n, p)
	url := fmt.Sprintf("https://public-ubiservices.ubi.com/v3/profiles?namesOnPlatform=%s&platformType=%s", n, p)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}

	sd, ok := pc.sm.Load("session")
	if !ok {
		return "", errors.New("error retrieving sync map session")
	}

	re, ok := sd.(map[string]string)
	if !ok {
		return "", errors.New("type assertion error")
	}
	req.Header = http.Header{
		"Authorization": []string{fmt.Sprintf("Ubi_v1 t=%s", re["ticket"])},
		"Ubi-AppId":     []string{"39baebad-39e5-4552-8c25-2c9b919064e2"},
		"Ubi-SessionId": []string{re["sessionKey"]},
		"Connection":    []string{"keep-alive"},
	}

	//appid? 3587dcbb-7f81-457c-9781-0e3f29f6f56a
	res, _ := pc.hc.Do(req)
	fmt.Println(res)
	if res.StatusCode != 200 {
		return "", errors.New(fmt.Sprintf("error fetching profileId STATUS CODE %v // S: %s", res.StatusCode, res.Status))
	}
	log.Println("BODY: ", res.Body)
	var player models.PlayerIDModel
	//var player map[string]string
	de := json.NewDecoder(res.Body).Decode(&player)
	if de != nil {
		return "", errors.New("error decoding player")
	}
	res.Body.Close()
	return player.Profiles[0].IDOnPlatform, nil
}

func (pc *PlayerController) fetchPlayer(ctx context.Context, n, p string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	//search elastic (check direct name and alias)

	//if found, put in cache
	//return to request

	//if not found check for profileId
	profileId, err := pc.fetchProfileId(ctx, n, p)
	log.Println(profileId)
	return []byte(profileId), err
}

func (pc *PlayerController) Test(w http.ResponseWriter, r *http.Request) {
	platform := strings.ToLower(chi.URLParam(r, "platform"))
	player := strings.ToLower(chi.URLParam(r, "player"))
	res, _ := pc.fetchPlayer(r.Context(), player, platform)
	w.Write(res)
}

func (pc *PlayerController) GetPlayers(w http.ResponseWriter, r *http.Request) {
	//platform := chi.URLParam(r, "platform")
	//player := chi.URLParam(r, "player")
}

func (pc *PlayerController) Player(w http.ResponseWriter, r *http.Request) {

}
