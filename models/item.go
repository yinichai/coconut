package models

import (
	"math"
)

//Item record data model in db
type Item struct {
	Name  string  `db:"item_name" json:"item_name"`
	Lat   float64 `db:"lat" json:"lat"`
	Lng   float64 `db:"lng" json:"lng"`
	URL   string  `db:"item_url" json:"item_url"`
	Image string  `db:"img_urls" json:"img_urls"`
}

//Location lat and lng
type Location struct {
	Lat float64
	Lng float64
}

//GetLocation get location
func (i Item) GetLocation() Location {
	return Location{
		Lat: i.Lat,
		Lng: i.Lng,
	}
}

// DistanceSquare power of 2 of two location distance
func (l Location) DistanceSquare(des Location) float64 {
	// In order to have one less operation, thereis no need to calulate math.Sqrt here, since the distance will be used as rank order
	return math.Pow((l.Lat-des.Lat), 2) + math.Pow((l.Lng-des.Lng), 2)
}
