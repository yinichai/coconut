package models

//DistanceMap model for distance
type DistanceMap struct {
	Item     *Item   `json:"item"`
	Score    int64   `json:"-"`
	Distance float64 `json:"-"`
}

//DistanceMapSlice slice of distance map
type DistanceMapSlice []DistanceMap

func (slice DistanceMapSlice) Len() int {
	return len(slice)
}

func (slice DistanceMapSlice) Less(i, j int) bool {
	return slice[i].Distance < slice[j].Distance
}

func (slice DistanceMapSlice) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}
