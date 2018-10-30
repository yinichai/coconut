package data

import (
	"log"
	"sort"

	. "github.com/yinichai/coconut/env"
	"github.com/yinichai/coconut/models"
)

// GetBestResults - search
func GetBestResults(input models.InputQuery) (*models.DistanceMapSlice, error) {
	var results []models.Item
	if err := DB.Select(&results, exactSearchQuery(input.SearchTerm)); err != nil {
		return nil, err
	}

	distanceResults := sortResultsByDistance(input, results)

	if len(distanceResults) > 20 {
		bestTwenty := distanceResults[:20]
		log.Print("1...")
		return &bestTwenty, nil
	}
	log.Print("2...")
	return &distanceResults, nil

}

func sortResultsByDistance(input models.InputQuery, results []models.Item) models.DistanceMapSlice {
	var (
		length   = len(results)
		distList models.DistanceMapSlice
	)

	for i := 0; i < length; i++ {
		distList = append(distList, models.DistanceMap{
			Item:     &results[i],
			Distance: results[i].GetLocation().DistanceSquare(input.GetLocation()),
		})
	}

	// sort array by distance O(n*log(n))
	sort.Sort(distList)
	return distList
}

func exactSearchQuery(keywords []string) string {
	var (
		baseQuery = "SELECT * FROM items WHERE"
		length    = len(keywords)
	)

	for i := 0; i < length; i++ {
		baseQuery += " item_name LIKE '%" + keywords[i] + "%'"
		if i != (length - 1) {
			baseQuery += " OR"
		}
	}

	return baseQuery
}
