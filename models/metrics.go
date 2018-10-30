package models

//DistanceMap model for distance
type DistanceMap struct {
	Item     *Item       `json:"item"`
	Input    *InputQuery `json:"-"`
	Score    int         `json:"-"`
	Distance float64     `json:"-"`
}

//TotalScore calculate score for results
func (d *DistanceMap) TotalScore(match int) {
	lw := getLocationWeight(d.Input.GetLocation().DistanceSquare(d.Item.GetLocation()))
	// hard code weight, location is treated more important than keyword match
	d.Score = match*2 + lw*5
}

func getLocationWeight(distanceSquare float64) int {
	switch {
	case distanceSquare < 2:
		return 10
	case distanceSquare < 4:
		return 5
	default:
		return 1
	}

}

//DistanceMapSlice slice of distance map
type DistanceMapSlice []DistanceMap

func (slice DistanceMapSlice) Len() int {
	return len(slice)
}

func (slice DistanceMapSlice) Less(i, j int) bool {
	return slice[i].Score > slice[j].Score
}

func (slice DistanceMapSlice) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}
