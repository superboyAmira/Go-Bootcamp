package handler

import (
	"context"
	"goday03/src/internal/app/model"
	"log"
	"os"
	"strconv"

	"encoding/csv"

	"github.com/olivere/elastic/v7"
)

const (
	dataFile string = "../../../materials/data.csv"
)


func LoadData(client *elastic.Client, ctx context.Context) {
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
