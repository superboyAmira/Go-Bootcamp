package handler

import (
	"context"
	"log"
	"os"
	"strconv"

	"encoding/csv"

	"github.com/olivere/elastic"
)

const (
	dataFile string = "../../../materials/data.csv"
)

type GeoPoint struct {
	Longitude float64 `json:"lon"`
	Latitude  float64 `json:"lat"`
}

type Place struct {
	Name     string   `json:"name"`
	Address  string   `json:"address"`
	Phone    string   `json:"phone"`
	Location GeoPoint `json:"location"`
}

func LoadData(client *elastic.Client, ctx *context.Context) {
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
	places = places[1:]
	log.Println(places[0][0])
	log.Println(places[0][1])
	log.Println(places[0][2])
	log.Println(places[0][3])
	log.Println(places[0][4])
	log.Println(places[0][5])

	for _, place := range places {
		
		lat, err := strconv.ParseFloat(place[4], 64)
		if err != nil {
			log.Fatalf("Error parsing latitude: %s", err)
		}
		lon, err := strconv.ParseFloat(place[5], 64)
		if err != nil {
			log.Fatalf("Error parsing longitude: %s", err)
		}
		pl := Place{
			Name:    place[1],
			Address: place[2],
			Phone:   place[3],
			Location: GeoPoint{
				Longitude: lon,
				Latitude:  lat,
			},
		}
		request := elastic.NewBulkIndexRequest().Index("places").Doc(pl)
		bulkRequest = bulkRequest.Add(request)
	}
	bulkRequest.ErrorTrace(true)
	response, err := bulkRequest.Do(*ctx)
	if err != nil {
		log.Fatalf("BulkRequest err: %s", err)
	}
	if response.Errors {
		log.Println("Bulk request completed with errors")
		// for _, item := range response.Failed() {
        //     log.Printf("Error indexing document ID %s: %s", item.Id, item.Error.Reason)
        // }
	} else {
		log.Println("Bulk request completed successfully")
	}
}
