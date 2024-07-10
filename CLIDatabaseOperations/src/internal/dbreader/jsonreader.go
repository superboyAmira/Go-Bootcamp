package dbreader

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"goday01/internal/model/recipes"
	"log/slog"
	"os"
)

type JSONReader struct {
	Data recipes.Recipes `xml:"recipes"`
}

func (r *JSONReader) Load(path string, log *slog.Logger) (error, *recipes.Recipes) {
	file, err := os.Open(path)
	if err != nil {
		return err, nil
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&r.Data)
	log.Debug(r.Data.Cake[0].Name)
	if err != nil {
		return err, nil
	}
	return nil, &r.Data
}

func (r *JSONReader) MustProcess(log *slog.Logger) {
	if r == nil {
		panic("No data loaded")
	}

	type WrappedRecipes struct {
		XMLName xml.Name          `xml:"recipes"`
		Recipes recipes.Recipes `xml:"cake"`
	}

	wrappedData := WrappedRecipes{
		Recipes: r.Data,
	}

	xmlData, err := xml.MarshalIndent(wrappedData, "", "    ")
	if err != nil {
		panic("Error marshaling to XML:" + err.Error())
	}
	fmt.Println(string(xmlData))
}
