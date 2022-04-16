package elastic

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"log"
)

var settings = `
{
  "settings": {
    "analysis": {
      "analyzer": {
        "search": {
          "tokenizer": "search",
          "filter": [
            "lowercase"
          ]
        }
      },
      "tokenizer": {
        "search": {
          "type": "edge_ngram",
          "min_gram": 3,
          "max_gram": 20
        }
      }
    }
  }
}
`

var mappings = `
{
  "properties": {
    "profileId": {
      "type": "keyword"
    },
    "userId": {
      "type": "keyword"
    },
    "platform": {
      "type": "keyword",
      "index": "false"
    },
    "nickName": {
      "type": "keyword"
    },
    "lastUpdate": {
      "type": "date",
      "index": "false"
    },
    "timePlayed": {
      "type": "object",
      "properties": {
        "value": {
          "type": "long",
          "index": "false"
        },
        "lastModified": {
          "type": "date",
          "index": "false"
        }
      }
    },
    "aliases": {
      "properties": {
        "name": {
          "type": "text",
          "analyzer": "search"
        },
        "date": {
          "type": "date"
        }
      }
    },
    "level": {
      "type": "object",
      "properties": {
        "value": {
          "type": "long",
          "index": "false"
        },
        "lastModified": {
          "type": "date",
          "index": "false"
        }
      }
    },
    "summary": {
      "type": "object",
      "enabled": "false"
    },
    "ranked": {
      "type": "object",
      "properties": {
        "currentSeason": {
          "type": "object",
          "properties": {
            "season": {
              "type": "long",
              "index": "false"
            },
            "max_mmr": {
              "type": "float",
              "index": "false"
            },
            "skill_mean": {
              "type": "float",
              "index": "false"
            },
            "deaths": {
              "type": "long",
              "index": "false"
            },
            "rank": {
              "type": "long",
              "index": "false"
            },
            "max_rank": {
              "type": "long",
              "index": "false"
            },
            "skill_stdev": {
              "type": "float",
              "index": "false"
            },
            "kills": {
              "type": "long",
              "index": "false"
            },
            "last_match_mmr_change": {
              "type": "float",
              "index": "false"
            },
            "abandons": {
              "type": "long",
              "index": "false"
            },
            "mmr": {
              "type": "float",
              "index": "false"
            },
            "last_match_result": {
              "type": "long",
              "index": "false"
            },
            "wins": {
              "type": "long",
              "index": "false"
            },
            "losses": {
              "type": "long",
              "index": "false"
            }
          }
        },
        "rankedSeasons": {
          "type": "object",
          "properties": {
            "season": {
              "type": "long",
              "index": "false"
            },
            "max_mmr": {
              "type": "float",
              "index": "false"
            },
            "skill_mean": {
              "type": "float",
              "index": "false"
            },
            "deaths": {
              "type": "long",
              "index": "false"
            },
            "rank": {
              "type": "long",
              "index": "false"
            },
            "max_rank": {
              "type": "long",
              "index": "false"
            },
            "skill_stdev": {
              "type": "float",
              "index": "false"
            },
            "kills": {
              "type": "long",
              "index": "false"
            },
            "last_match_mmr_change": {
              "type": "float",
              "index": "false"
            },
            "abandons": {
              "type": "long",
              "index": "false"
            },
            "mmr": {
              "type": "float",
              "index": "false"
            },
            "last_match_result": {
              "type": "long",
              "index": "false"
            },
            "wins": {
              "type": "long",
              "index": "false"
            },
            "losses": {
              "type": "long",
              "index": "false"
            }
          }
        }
      }
    },
    "weapons": {
      "type": "object",
      "enabled": "false"
    },
    "operators": {
      "type": "object",
      "enabled": "false"
    },
    "maps": {
      "type": "object",
      "enabled": "false"
    },
    "trends": {
      "type": "object",
      "enabled": "false"
    }
  }
}
`

// ElasticSetup migration for elasticsearch indices on server start (dev only or first time setup)
func InitialSetup(ctx context.Context, es *elasticsearch.Client) error {
	res, err := es.Indices.Exists([]string{"r6index.uplay", "r6index.psn", "r6index.xbl"})
	if err != nil {
		return err
	}

	//if res.IsError() {
	//	return errors.New(res.Status())
	//}

	//indices exist so don't run migration
	if res.StatusCode == 200 {
		return nil
	}

	for _, index := range []string{"r6index.uplay", "r6index.psn", "r6index.xbl"} {
		//Delete Index
		res, err := es.Indices.Delete([]string{index})
		if err != nil {
			fmt.Printf("Cannot delete index: %s", err)
		}
		if res.IsError() {
			fmt.Printf("cannot delete index %s", index)
		}

		res.Body.Close()

		//Create Index
		res, err = es.Indices.Create(index)
		if err != nil {
			log.Fatalf("Cannot create index: %s", err)
		}
		if res.IsError() {
			log.Fatalf("Cannot create index: %s", res)
		}
		res.Body.Close()

		//Close index to apply settings
		res, err = esapi.IndicesCloseRequest{Index: []string{index}}.Do(ctx, es)
		if err != nil {
			log.Fatalf("cannot close index: %s", index)
		}

		if res.IsError() {
			log.Fatalf("cannot close index: %s", index)
		}
		res.Body.Close()

		//Apply Settings to Index
		res, err = esapi.IndicesPutSettingsRequest{Index: []string{index}, Body: bytes.NewBufferString(settings)}.Do(ctx, es)
		if err != nil {
			log.Fatalf("cannot apply settings to index: %s", index)
		}

		if res.IsError() {
			log.Fatalf("cannot apply settings to index: %s", index)
		}
		res.Body.Close()

		//Open index after applying settings
		res, err = esapi.IndicesOpenRequest{Index: []string{index}}.Do(ctx, es)
		if err != nil {
			log.Fatalf("cannot apply settings to index: %s", index)
		}

		if res.IsError() {
			log.Fatalf("cannot apply settings to index: %s", index)
		}
		res.Body.Close()
	}

	res, err = esapi.IndicesPutMappingRequest{Index: []string{"r6index.uplay", "r6index.psn", "r6index.xbl"}, Body: bytes.NewBufferString(mappings)}.Do(ctx, es)
	if err != nil {
		return err
	}

	if res.IsError() {
		return errors.New("error applying mappings to all indices")
	}

	log.Println("finished running elasticsearch initial setup")
	return nil
}
