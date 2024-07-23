package repository

import (
	"context"
	"goday03/src/internal/app/model"
	"log"

	"github.com/olivere/elastic/v7"
)

type Store interface {
	GetPlaces(limit int, offset int) ([]model.Place, int, error)
}

type PlacesRepo struct {
	client *elastic.Client
}

func (r *PlacesRepo) GetPlaces(limit int, offset int) ([]model.Place, int, error) {

	res, err := r.client.Search("places").From(offset).Size(limit).Do(context.Background())
	if err != nil {
		return nil, 0, err
	}
	

	if (res.TotalHits() == 0) {
		return nil, 0, nil
	}
	var ret []model.Place
	for _, hit := range res.Hits.Hits {
		var _tmp model.Place
		hit.Source.UnmarshalJSON()
		ret = append(ret, )
	}
	return ret, int(res.TotalHits()), nil
	
}