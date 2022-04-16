package controllers

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/eliassebastian/gor6-api/cmd/api/response"
	"github.com/eliassebastian/gor6-api/internal/elastic"
	"log"
	"net/http"
	"sync"
	"time"
)

type SearchController struct {
	ec *elastic.ESClient
	sm *sync.Map
	hc *http.Client
}

func NewSearchController(c *elastic.ESClient, tlsc *tls.Config, p *sync.Map) *SearchController {
	return &SearchController{
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

type searchPayload struct {
	Player   string
	Platform string
}

//curl -H "Content-Type: application/json" -X POST -d '{"player":"Kanzen","platform":"uplay"}' https://localhost:8090/test
func (sc *SearchController) SearchPlayer(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	log.Println("Search Player Running???")

	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	var sp searchPayload
	err := json.NewDecoder(r.Body).Decode(&sp)
	if err != nil {
		response.ErrorJSON(w, err)
		return
	}

	defer r.Body.Close()
	fmt.Println(sp.Player, sp.Platform)

	n, res, err := sc.ec.SearchPlayer(ctx, sp.Player, sp.Platform)
	if err != nil {
		response.ErrorJSON(w, err)
		return
	}

	log.Println("Search Controller Search Player Func - MaxScore:", n, "for Player:", sp.Player)

	response.SuccessJSON(w, startTime, res)
}
