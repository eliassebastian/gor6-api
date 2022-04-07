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
		DocumentID: player.ProfileId,
	}

	res, err := req.Do(ctx, c.Client)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	if res.IsError() {
		log.Println(res)
		return errors.New(fmt.Sprintf("[%s] error indexing document id=%s", res.Status(), player.ProfileId))
	}

	return nil
}

func (c *ESClient) SearchPlayer(ctx context.Context, name, platform string) (*model.SearchHits, error) {

	fmt.Println(name, platform)

	var buf bytes.Buffer
	query := model.SearchInput{
		Query: model.SearchQuery{
			Match: model.SearchMatch{AliasesName: name}},
		Fields: []string{"profileId", "platform", "nickName", "aliases.name", "level.value", "timePlayed.value", "timePlayed.lastModified", "ranked.currentSeason.rank"},
		Source: false,
		Size:   10,
	}

	//s := fmt.Sprintf(`{"query": {"match": {"aliases.name": %s} }, fields": ["profileId", "platform", "nickName", "aliases.name", "level.value", "timePlayed.value", "timePlayed.lastModified", "ranked.currentSeason.rank"], "_source": false}`, name)
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		return nil, err
	}

	res, err := c.Client.Search(
		c.Client.Search.WithContext(ctx),
		c.Client.Search.WithIndex("r6index."+platform),
		c.Client.Search.WithBody(&buf),
		//c.Client.Search.WithStoredFields("profileId", "platform", "nickName", "aliases.name", "level.value", "timePlayed.value", "timePlayed.lastModified", "ranked.currentSeason.rank"),
		//c.Client.Search.WithSize(10),
		//c.Client.Search.WithSource("false"),
		c.Client.Search.WithTrackTotalHits(true),
	)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	if res.IsError() {
		log.Println(res)
		return nil, errors.New(fmt.Sprintf("[%s] error searching index %s for player=%s", res.Status(), platform, name))
	}

	var output model.SearchOutput
	if err := json.NewDecoder(res.Body).Decode(&output); err != nil {
		return nil, err
	}

	return &output.Hits, nil
}
