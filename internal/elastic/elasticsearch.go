package elastic

import (
	"errors"
	"github.com/elastic/go-elasticsearch/v7"
)

type ESClient struct {
	client *elasticsearch.Client
}

func NewElasticClient() (*ESClient, error) {
	//TODO: Secure ElasticSearch Client
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		return nil, errors.New("error creating new elasticsearch client")
	}

	res, err := es.Info()
	if err != nil {
		return nil, errors.New("error pinging new elasticsearch client")
	}

	res.Body.Close()
	return &ESClient{
		client: es,
	}, nil
}
