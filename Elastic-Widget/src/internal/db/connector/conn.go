package connector

import (
	"context"
	"encoding/json"
	"goday03/src/internal/db/loader"
	"log"
	"os"

	"github.com/olivere/elastic/v7"
)

const (
	dlimCluster string = "http://localhost:9200"
	schemaPath  string = "../../internal/web/json/schema.json" // В Elasticsearch версии 7 и выше был изменен формат маппинга, и ключ properties теперь должен находиться внутри ключа mappings.
)

func Conn() {
	ctx := context.Background()
	client := getClient()
	createIndexPlaces(client, ctx)
	handler.LoadData(client, ctx)
}

func getClient() *elastic.Client {
	client, err := elastic.NewClient(elastic.SetURL(dlimCluster))
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}
	return client
}

func createIndexPlaces(client *elastic.Client, ctx context.Context) {
	info, code, err := client.Ping(dlimCluster).Do(ctx)
	if err != nil {
		log.Fatalf("Error pinging Elasticsearch: %s", err)
	}
	log.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)


	if exist, _ := client.IndexExists("places").Do(ctx); exist {
		client.DeleteIndex("places").Do(ctx)
	}

	dataJSON := loadSchema(schemaPath)

	createIndex, err := client.CreateIndex("places").BodyJson(json.RawMessage(dataJSON)).Do(ctx)
	if err != nil {
		log.Fatalf("Error creating the index: %s", err)
	}
	if !createIndex.Acknowledged {
		log.Fatalf("CreateIndex was not acknowledged. Check that timeout value is correct.")
	}
}

func loadSchema(pathJSON string) []byte {
	schemaData, err := os.ReadFile(pathJSON)
	if err != nil {
		log.Fatalf("Err with opening JOSN schema with path [%s] : %s", pathJSON, err)
	}

	return schemaData
}
