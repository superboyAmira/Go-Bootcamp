package client

import (
	"log"

	"github.com/olivere/elastic/v7"
)

const (
	connClusterURL string = "http://localhost:9200"
)

func GetClient() *elastic.Client {
	client, err := elastic.NewClient(elastic.SetURL(connClusterURL))
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}
	return client
}