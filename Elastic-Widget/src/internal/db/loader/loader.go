package loader

import (
	"context"
	"goday03/src/internal/app/model"
	"goday03/src/internal/db/client"
	"log"
	"os"
	"strconv"

	"encoding/csv"
	"encoding/json"

	"github.com/olivere/elastic/v7"
)

const (
	dlimCluster string = "http://localhost:9200"
	schemaPath  string = "../../internal/web/json/schema.json" // В Elasticsearch версии 7 и выше был изменен формат маппинга, и ключ properties теперь должен находиться внутри ключа mappings.
)

const (
	dataFile string = "../../../materials/data.csv"
)

func LoadData() {
	ctx := context.Background()
	client := client.GetClient()
	createIndexPlaces(client, ctx)
	loadCSV(client, ctx)
}

func loadCSV(client *elastic.Client, ctx context.Context) {
	file, err := os.Open(dataFile)
	if err != nil {
		log.Fatalf("Bad datafile path: %s", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = '\t'
	places, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Error reading the datafile: %s", err)
	}

	bulkRequest := client.Bulk()

	for _, place := range places[1:] {
		lat, err := strconv.ParseFloat(place[5], 64)
		if err != nil {
			log.Fatalf("Error parsing latitude: %s", err)
		}
		lon, err := strconv.ParseFloat(place[4], 64)
		if err != nil {
			log.Fatalf("Error parsing longitude: %s", err)
		}
		pl := model.Place{
			Name:    place[1],
			Address: place[2],
			Phone:   place[3],
			Location: model.GeoPoint{
				Longitude: lon,
				Latitude:  lat,
			},
		}
		id := place[0]

		request := elastic.NewBulkIndexRequest().Index("places").Id(id).Doc(pl)
		bulkRequest = bulkRequest.Add(request)
	}
	response, err := bulkRequest.Do(ctx)
	if err != nil {
		log.Fatalf("BulkRequest err: %s", err)
	}

	if response.Errors {
		log.Println("Bulk request completed with errors")
		for _, item := range response.Failed() {
			log.Printf("Error indexing document ID %s: %s", item.Id, item.Error.Reason)
		}
	} else {
		cnt, err := client.Count("places").Do(ctx)
		if err != nil {
			log.Fatalf("No data loaded in index: %s", err)
		}
		log.Printf("Bulk request completed successfully, cnt of data`s: %d\n", cnt)
	}
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