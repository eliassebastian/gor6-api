package elastic

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	model "github.com/eliassebastian/gor6-api/cmd/api/models"
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
	//for _, index := range []string{"r6index.uplay", "r6index.psn", "r6index.xbl"} {
	//	if res, err = es.Indices.Delete([]string{index}); err != nil {
	//		log.Fatalf("Cannot delete index: %s", err)
	//	}
	//	res.Body.Close()
	//
	//	res, err = es.Indices.Create(index)
	//	if err != nil {
	//		log.Fatalf("Cannot create index: %s", err)
	//	}
	//	if res.IsError() {
	//		log.Fatalf("Cannot create index: %s", res)
	//	}
	//	res.Body.Close()
	//}

	res.Body.Close()
	return &ESClient{
		Client: es,
	}, nil
}

func (c *ESClient) IndexPlayer(ctx context.Context, player *model.Player, platform string) error {
	var buf bytes.Buffer
	_ = json.NewEncoder(&buf).Encode(player)

	req := esapi.IndexRequest{
		Index:      "r6index." + platform,
		Body:       &buf,
		DocumentID: player.ID,
	}

	res, err := req.Do(ctx, c.Client)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	if res.IsError() {
		log.Println(res)
		return errors.New(fmt.Sprintf("[%s] error indexing document id=%s", res.Status(), player.ID))
	}

	return nil
}
