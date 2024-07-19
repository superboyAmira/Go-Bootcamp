package dbreader

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"goday01/internal/model/recipes"
	"io"
	"log/slog"
	"os"
)

type XMLReader struct {
	Data recipes.Recipes `xml:"recipes"`
}

func (r *XMLReader) Load(path string, log *slog.Logger) (error, *recipes.Recipes) {
	file, err := os.Open(path)
	if err != nil {
		return err, nil
	}
	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		return err, nil
	}
	err = xml.Unmarshal(byteValue, &r.Data)
	log.Debug(r.Data.Cake[0].Name)
	if err != nil {
		return err, nil
	}
	return nil, &r.Data
}

func (r *XMLReader) MustProcess(log *slog.Logger) {
	if r == nil {
		panic("No data loaded!")
	}

	jsonData, err := json.MarshalIndent(r.Data, "", "    ")
	if err != nil {
		panic("Error marshaling to JSON:" + err.Error())
	}
	fmt.Println(string(jsonData))
}
