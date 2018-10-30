package data

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yinichai/coconut/models"
)

func Test_extactSearchQuery(t *testing.T) {

	var keywords = []string{"camera", "new"}

	assert.Equal(
		t,
		exactSearchQuery(keywords),
		"SELECT * FROM items WHERE item_name LIKE '%camera%' OR item_name LIKE '%new%'",
	)
}

func Test_sortResultByDistance(t *testing.T) {
	var (
		input = models.InputQuery{
			SearchTerm: []string{"camera", "canon"},
			Lat:        0,
			Lng:        0,
		}
		results = []models.Item{
			models.Item{"canon", 1, 2, "url_closest", "url"},
			models.Item{"canon", 2, 2, "url", "url"},
			models.Item{"canon", 3, 2, "url", "url"},
			models.Item{"canon", 4, 2, "url", "url"},
			models.Item{"canon", 5, 2, "url", "url"},
			models.Item{"canon", 6, 2, "url", "url"},
			models.Item{"canon", 7, 2, "url", "url"},
			models.Item{"canon", 8, 2, "url", "url"},
			models.Item{"canon", 9, 2, "url", "url"},
			models.Item{"canon", 10, 2, "url", "url"},
			models.Item{"canon", 11, 2, "url", "url"},
			models.Item{"canon", 12, 2, "url", "url"},
			models.Item{"canon", 13, 2, "url", "url"},
			models.Item{"canon", 14, 2, "url", "url"},
			models.Item{"canon", 15, 2, "url", "url"},
			models.Item{"canon", 16, 2, "url", "url"},
			models.Item{"canon", 17, 2, "url", "url"},
			models.Item{"canon", 18, 2, "url", "url"},
			models.Item{"canon", 19, 2, "url", "url"},
			models.Item{"canon", 20, 2, "url", "url"},
			models.Item{"canon", 21, 2, "url", "url"},
		}
	)

	assert.Equal(t, sortResultsByDistance(input, results)[0].Item.URL, "url_closest")

}
