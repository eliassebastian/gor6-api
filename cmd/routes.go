package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/eliassebastian/gor6-api/cmd/api/graph"
	"github.com/eliassebastian/gor6-api/cmd/api/graph/generated"
	"github.com/eliassebastian/gor6-api/internal/elastic"
	"github.com/eliassebastian/gor6-api/internal/mongodb"
	"github.com/go-chi/chi/v5"
	"net/http"
	"sync"
	"time"
)

func graphQLHandler(c *elastic.ESClient, m *mongodb.MongoClient, p *sync.Map) http.HandlerFunc {
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: &graph.Resolver{
			EC: c,
			MC: m,
			SM: p,
			HC: &http.Client{
				Timeout: 10 * time.Second,
			},
		},
	}))
	return h.ServeHTTP
}

func playgroundQLHandler(endpoint string) http.HandlerFunc {
	//endpoint argument must be same as graphql handler path
	playgroundHandler := playground.Handler("R6Index", endpoint)
	return playgroundHandler
}

func routes(c serverConfig) http.Handler {
	mux := chi.NewRouter()
	//middleware
	//routes
	//mux.Get("/search/{platform}/{player}", pc.GetPlayers)
	//mux.Get("/player/{platform}/{id}", pc.Player)
	//GraphQL Endpoint
	mux.Post("/query", graphQLHandler(c.ElasticSearch, c.MongoDB, c.Kafka.Cache))
	//DEV ONLY
	mux.Get("/playground", playgroundQLHandler("/query"))

	//mux.Get("/test/{platform}/{player}", pc.Test)

	return mux
}
