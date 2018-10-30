package models

//InputQuery - parameters in query
type InputQuery struct {
	SearchTerm []string
	Lat        float64
	Lng        float64
}

//GetLocation get location
func (i InputQuery) GetLocation() Location {
	return Location{
		Lat: i.Lat,
		Lng: i.Lng,
	}
}
