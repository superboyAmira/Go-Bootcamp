package dbreader

import (
	"errors"
	"log/slog"
	"goday01/internal/model/recipes"
)

type DBReader interface {
	Load(path string, log *slog.Logger) (error, *recipes.Recipes)
	MustProcess(log *slog.Logger)
}

func GetReader(ext string) (DBReader, error) {
	if ext == ".xml" {
		return &XMLReader{}, nil
	} else if ext == ".json" {
		return &JSONReader{}, nil
	} else {
		return nil, errors.New("incorrect file type")
	}
}