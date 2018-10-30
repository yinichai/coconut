package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/yinichai/coconut/data"
	"github.com/yinichai/coconut/handlers/responses"
	"github.com/yinichai/coconut/models"
)

/*
search?searchTerm=camera&lat=51.948&lng=0.172943
*/

//Search - search endpoint
func Search(w http.ResponseWriter, r *http.Request) {
	var (
		params = r.URL.Query()
		lat    float64
		lng    float64
		err    error

		inputQuery models.InputQuery
	)

	if params.Get("lat") == "" {
		responses.BadRequest(w, r, fmt.Errorf("missing_lat"))
		return
	}

	if params.Get("lng") == "" {
		responses.BadRequest(w, r, fmt.Errorf("missing_lng"))
		return
	}

	if params.Get("searchTerm") == "" {
		responses.BadRequest(w, r, fmt.Errorf("missing_keywords"))
		return
	}

	lat, err = strconv.ParseFloat(params.Get("lat"), 64)
	if err != nil {
		panic(err)
	}

	lng, err = strconv.ParseFloat(params.Get("lng"), 64)
	if err != nil {
		panic(err)
	}

	inputQuery = models.InputQuery{
		SearchTerm: strings.Split(params.Get("searchTerm"), " "),
		Lat:        lat,
		Lng:        lng,
	}

	results, err := data.GetBestResults(inputQuery)
	if err != nil {
		responses.InternalServerError(w, r, err)
		return
	}

	if results == nil {
		responses.NotFound(w, r, fmt.Errorf("no results"))
		return
	}
	if err := json.NewEncoder(w).Encode(results); err != nil {
		responses.InternalServerError(w, r, err)
		return
	}
}
