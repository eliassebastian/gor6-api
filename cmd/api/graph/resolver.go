package graph

import (
	"github.com/eliassebastian/gor6-api/internal/elastic"
	"github.com/eliassebastian/gor6-api/internal/mongodb"
	"net/http"
	"sync"
)

//go:generate go run -mod=mod github.com/99designs/gqlgen
//go run -mod=mod github.com/99designs/gqlgen

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	EC *elastic.ESClient
	MC *mongodb.MongoClient
	SM *sync.Map
	HC *http.Client
}
