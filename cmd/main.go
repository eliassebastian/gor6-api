package main

import (
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
)

func main() {

}

func run() {

}

type serverConfig struct {
	Address       string
	MongoDB       *mongo.Client
	ElasticSearch *elasticsearch.Client
}

func newServer(c serverConfig) *http.Server {
	mux := chi.NewRouter()
	log.Println(mux)
	return &http.Server{}
}
