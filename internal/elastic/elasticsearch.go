package elastic

import (
	"errors"
	"github.com/elastic/go-elasticsearch/v7"
	"log"
)

type ESClient struct {
	Client *elasticsearch.Client
}

func NewElasticClient() (*ESClient, error) {
	//TODO: Secure ElasticSearch Client
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200",
		},
	}

	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return nil, errors.New("error creating new elasticsearch client")
	}

	res, err := es.Info()
	if err != nil {
		return nil, errors.New("error pinging new elasticsearch client")
	}

	//DEV ONLY
	for _, index := range []string{"r6index.uplay", "r6index.psn", "r6index.xbl"} {
		if res, err = es.Indices.Delete([]string{index}); err != nil {
			log.Fatalf("Cannot delete index: %s", err)
		}

		res, err = es.Indices.Create(index)
		if err != nil {
			log.Fatalf("Cannot create index: %s", err)
		}

		if res.IsError() {
			log.Fatalf("Cannot create index: %s", res)
		}
	}

	res.Body.Close()
	return &ESClient{
		Client: es,
	}, nil
}
