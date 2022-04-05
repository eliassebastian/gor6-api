package elastic

import (
	"errors"
	"github.com/elastic/go-elasticsearch/v7"
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

	res.Body.Close()
	return &ESClient{
		Client: es,
	}, nil
}
