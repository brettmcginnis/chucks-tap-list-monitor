package site

import (
	"encoding/json"
)

type beerJson struct {
	Tap      int         `json:"tap,omitempty"`
	Beer     string      `json:"beerJson,omitempty"`
	Shop     string      `json:"shop,omitempty"`
	No16     string      `json:"No16,omitempty"`
	Special  string      `json:"Special,omitempty"`
	NoGr     string      `json:"NoGr,omitempty"`
	KegCost  string      `json:"Keg$,omitempty"`
	ABV      string      `json:"abv,omitempty"`
	Origin   string      `json:"origin,omitempty"`
	Type     string      `json:"type,omitempty"`
	Color    string      `json:"color,omitempty"`
	Size     string      `json:"Size,omitempty"`
	Serving  string      `json:"serving,omitempty"`
	Oz       int         `json:"oz,omitempty"`
	CostOz   float64     `json:"costOz,omitempty"`
	PriceOz  float64     `json:"priceOz,omitempty"`
	Full     json.Number `json:"Full,omitempty"`
	Half     string      `json:"Half,omitempty"`
	Quarter  string      `json:"Quarter,omitempty"`
	Growler  json.Number `json:"Growler,omitempty"`
	Crowler  json.Number `json:"Crowler,omitempty"`
	Override string      `json:"$Override,omitempty"`
}
