package elasticsearch

import (
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"strings"
)

func InitializeElasticsearchClient(address string) (*elasticsearch.Client, error) {
	cfg := elasticsearch.Config{
		Addresses: []string{
			address,
		},
	}

	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return nil, err
	}
	return es, nil
}

func ElasticsearchSearch(es *elasticsearch.Client, index string, query string) (map[string]interface{}, error) {
	res, err := es.Search(
		es.Search.WithIndex(index),
		es.Search.WithBody(strings.NewReader(query)),
	)
	if err != nil {
		return nil, err
	}
	fmt.Println(res)
	// Process and return search results
	return nil, nil
}
