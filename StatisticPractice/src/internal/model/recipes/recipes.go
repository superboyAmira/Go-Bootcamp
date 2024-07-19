package recipes

import "goday01/internal/model/cake"

type Recipes struct {
	Cake []cake.Cake `xml:"cake" json:"cake"`
}
