package repository

import (
	"context"
	"encoding/json"
	"goday03/src/internal/app/model"
	"goday03/src/internal/db/client"
	"io"

	"github.com/olivere/elastic/v7"
)

type Store interface {
	GetPlaces(limit int, offset int) ([]model.Place, int, error)
	ScrollApiPlaces(limit int) ([]model.Place, error)
}

type PlaceRepository struct {
	client *elastic.Client
}

func NewPlaceRepository() *PlaceRepository {
	return &PlaceRepository{client: client.GetClient()}
}

func (r *PlaceRepository) GetPlaces(limit int, offset int) ([]model.Place, int, error) {
	var ret []model.Place

	res, err := r.client.Search("places").From(offset).Size(limit).MaxResponseSize(20000).Do(context.Background())
	if err != nil {
		return nil, 0, err
	}

	if res.TotalHits() == 0 {
		return nil, 0, nil
	}

	for _, hit := range res.Hits.Hits {
		var tmp model.Place
		err := json.Unmarshal(hit.Source, &tmp)
		if err != nil {
			return nil, 0, err
		}
		ret = append(ret, tmp)
	}

	return ret, int(res.TotalHits()), nil
}

func (r *PlaceRepository) ScrollApiPlaces(limit int) ([]model.Place, error) {
	var allPlaces []model.Place

	scroll := r.client.Scroll("places").
		Size(limit).
		FetchSourceContext(elastic.NewFetchSourceContext(true)).
		Pretty(true)

	for {
		res, err := scroll.Do(context.Background())
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		for _, hit := range res.Hits.Hits {
			var tmp model.Place
			err := json.Unmarshal(hit.Source, &tmp)
			if err != nil {
				return nil, err
			}
			allPlaces = append(allPlaces, tmp)
		}
	}

	return allPlaces, nil
}