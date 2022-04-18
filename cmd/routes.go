package main

import (
	"github.com/eliassebastian/gor6-api/cmd/api/controllers"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func routes(c serverConfig) http.Handler {
	mux := chi.NewRouter()
	//initialise controllers
	pc := controllers.NewPlayerController(c.ElasticSearch, c.IndexCache, c.ProfileCache, c.TLS, c.Rabbit.Cache)
	sc := controllers.NewSearchController(c.ElasticSearch, c.IndexCache, c.TLS, c.Rabbit.Cache)
	uc := controllers.NewUpdateController(c.ElasticSearch, c.ProfileCache, c.TLS, c.Rabbit.Cache)
	//middleware

	//curl -H "Content-Type: application/json" -X POST -d '{"player":"Kanzen","platform":"uplay"}' https://localhost:8090/test
	mux.Post("/test", pc.Test)
	//curl -H "Content-Type: application/json" -X POST -d '{"id":"81bde55f-a30f-450a-94cb-4151b1a32130", "score": 0.320886, "platform":"uplay"}' https://localhost:8090/player
	//curl -H "Content-Type: application/json" -X POST -d '{"id":"4a0e6a89-4603-402d-a2b9-89a71612a95a", "score": -1, "platform":"uplay"}' https://localhost:8090/player
	mux.Post("/player", pc.Player)
	//curl -H "Content-Type: application/json" -X POST -d '{"player":"Kanz","platform":"uplay"}' https://localhost:8090/search
	mux.Post("/search", sc.SearchPlayer)
	//curl -H "Content-Type: application/json" -X POST -d '{"id":"81bde55f-a30f-450a-94cb-4151b1a32130", "platform":"uplay"}' https://localhost:8090/update
	mux.Post("/update", uc.UpdatePlayer)

	return mux
}
