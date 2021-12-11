package main

import (
	"github.com/eliassebastian/gor6-api/cmd/api/controllers"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

func routes(c serverConfig) http.Handler {
	mux := chi.NewRouter()
	log.Println(mux)

	//middleware
	pc := controllers.NewPlayerController(c.ElasticSearch, c.MongoDB, c.TLS, c.Kafka.Cache)
	//routes
	//mux.Get("/search/{platform}/{player}", pc.GetPlayers)
	//mux.Get("/player/{platform}/{id}", pc.Player)

	mux.Get("/test/{platform}/{player}", pc.Test)

	return mux
}
