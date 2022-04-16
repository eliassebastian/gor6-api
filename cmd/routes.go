package main

import (
	"github.com/eliassebastian/gor6-api/cmd/api/controllers"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func routes(c serverConfig) http.Handler {
	mux := chi.NewRouter()
	//initialise controllers
	pc := controllers.NewPlayerController(c.ElasticSearch, c.TLS, c.Rabbit.Cache)
	sc := controllers.NewSearchController(c.ElasticSearch, c.TLS, c.Rabbit.Cache)
	//middleware

	//curl -H "Content-Type: application/json" -X POST -d '{"player":"Kanzen","platform":"uplay"}' https://localhost:8090/test
	mux.Post("/test", pc.Test)
	//curl -H "Content-Type: application/json" -X POST -d '{"player":"Kanz","platform":"uplay"}' https://localhost:8090/search
	mux.Post("/search", sc.SearchPlayer)

	return mux
}
