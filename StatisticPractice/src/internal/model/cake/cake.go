package cake

import (
	"goday01/internal/model/item"
)

type Cake struct {
	Name        string           `xml:"name" json:"name"`
	Stovetime   string           `xml:"stovetime" json:"time"`
	Ingredients []item.Itemsdata `xml:"ingredients>item" json:"ingredients"`
}
