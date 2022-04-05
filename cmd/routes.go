package main

import (
	"github.com/eliassebastian/gor6-api/cmd/api/controllers"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func routes(c serverConfig) http.Handler {
	mux := chi.NewRouter()
	//initialise controllers
	pc := controllers.NewPlayerController(c.ElasticSearch, c.TLS, c.Kafka.Cache)

	//middleware

	//routes
	//mux.Get("/search/{platform}/{player}", pc.GetPlayers)
	//mux.Get("/player/{platform}/{id}", pc.Player)
	//GraphQL Endpoint
	//mux.Post("/query", graphQLHandler(c.ElasticSearch, c.MongoDB, c.Kafka.Cache))
	//DEV ONLY
	//mux.Get("/playground", playgroundQLHandler("/query"))

	//curl -X POST -H "Content-Type: application/json" \
	//-d '{"player": "Kanzen", "platform": "uplay"}' \
	//https://localhost:8090/test

	mux.Post("/test", pc.Test)

	return mux
}
