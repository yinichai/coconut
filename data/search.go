package data

import (
	"sort"
	"strings"

	. "github.com/yinichai/coconut/env"
	"github.com/yinichai/coconut/models"
)

// GetBestResults - search
func GetBestResults(input models.InputQuery) (*models.DistanceMapSlice, error) {
	var results []models.Item

	if err := DB.Select(&results, exactSearchQuery(input.SearchTerm)); err != nil {
		return nil, err
	}

	scoredResults := calculateResults(input, results)

	if len(scoredResults) > 20 {
		bestTwenty := scoredResults[:20]
		return &bestTwenty, nil
	}
	return &scoredResults, nil

}

func calculateResults(input models.InputQuery, results []models.Item) models.DistanceMapSlice {
	var (
		length    = len(results)
		distLists models.DistanceMapSlice
	)

	for i := 0; i < length; i++ {
		var (
			count    = 0
			distList = models.DistanceMap{&results[i], &input, 0, 0.0}
		)

		for _, k := range input.SearchTerm {
			if strings.Contains(results[i].Name, k) {
				count++
			}
			distList.TotalScore(count)
			distLists = append(distLists, distList)
		}
	}
	sort.Sort(distLists)
	return distLists
}

func exactSearchQuery(keywords []string) string {
	//TODO: expand results by semantic analysis
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
