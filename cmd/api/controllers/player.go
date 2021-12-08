package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/eliassebastian/gor6-api/internal/elastic"
	"github.com/eliassebastian/gor6-api/internal/mongodb"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"sync"
	"time"
)

type PlayerController struct {
	ec *elastic.ESClient
	mc *mongodb.MongoClient
	sm *sync.Map
}

func NewPlayerController(c *elastic.ESClient, m *mongodb.MongoClient, p *sync.Map) *PlayerController {
	return &PlayerController{
		ec: c,
		mc: m,
		sm: p,
	}
}

func (pc *PlayerController) Test(w http.ResponseWriter, r *http.Request) {
	platform := chi.URLParam(r, "platform")
	player := chi.URLParam(r, "player")

	client := &http.Client{
		Timeout: time.Second * 10,
	}

	url := fmt.Sprintf("https://public-ubiservices.ubi.com/v3/profiles?namesOnPlatform=%s&platformType=%s", player, platform)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Println(req, err)
	}

	sd, ok := pc.sm.Load("session")
	if !ok {
		log.Println("Error Retrieving Sync Map Session")
		return
	}

	re, ok := sd.(map[string]string)
	if !ok {
		log.Println("Type Assertion Error")
		return
	}
	req.Header = http.Header{
		"Authorization": []string{fmt.Sprintf("Ubi_v1 t=%s", re["ticket"])},
		"Ubi-AppId":     []string{"39baebad-39e5-4552-8c25-2c9b919064e2"},
		"Ubi-SessionId": []string{re["sessionKey"]},
		"Connection":    []string{"keep-alive"},
	}

	res, err := client.Do(req)

	fmt.Println("STATUS CODE:::::", res.StatusCode)

	if err != nil {
		log.Println(res, err)
	}

	log.Println("SUCCESS:  ", res, err, res.Body)
	defer res.Body.Close()

	var result map[string]interface{}

	err2 := json.NewDecoder(res.Body).Decode(&result)

	if err2 != nil {
		log.Fatalln(err)
	}

	fmt.Println(result)
}

func (pc *PlayerController) GetPlayers(w http.ResponseWriter, r *http.Request) {
	//platform := chi.URLParam(r, "platform")
	//player := chi.URLParam(r, "player")
}

func (pc *PlayerController) Player(w http.ResponseWriter, r *http.Request) {

}
