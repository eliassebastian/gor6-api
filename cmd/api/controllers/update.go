package controllers

import (
	"context"
	"crypto/tls"
	"github.com/eliassebastian/gor6-api/cmd/api/response"
	"github.com/eliassebastian/gor6-api/internal/cache"
	"github.com/eliassebastian/gor6-api/internal/elastic"
	"net/http"
	"sync"
	"time"
)

type UpdateController struct {
	ec *elastic.ESClient
	//mc *mongodb.MongoClient
	pc *cache.ProfileCache
	sm *sync.Map
	hc *http.Client
}

func NewUpdateController(c *elastic.ESClient, i *cache.ProfileCache, tlsc *tls.Config, p *sync.Map) *UpdateController {
	return &UpdateController{
		ec: c,
		pc: i,
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

type updatePayload struct {
	ID string
}

func (uc *UpdateController) UpdatePlayer(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()

	_, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	response.SuccessJSON(w, startTime, nil)
}
