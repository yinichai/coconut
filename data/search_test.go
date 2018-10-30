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

func Test_calculateResults(t *testing.T) {
	var (
		input = models.InputQuery{
			SearchTerm: []string{"camera", "canon"},
			Lat:        0,
			Lng:        0,
		}

		results = []models.Item{
			models.Item{"canon", 1, 1, "url_closest", "url"},
			models.Item{"camera canon", 2, 2, "url", "url"},
		}
	)

	assert.Equal(t, calculateResults(input, results)[0].Item.URL, "url_closest")
}
