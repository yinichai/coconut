package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"runtime/debug"
	"strconv"
	"strings"

	"github.com/yinichai/coconut/data"
	"github.com/yinichai/coconut/handlers/responses"
	"github.com/yinichai/coconut/models"
)

//Recover handle panic
func Recover(w http.ResponseWriter, r *http.Request) {
	if re := recover(); re != nil {
		log.Println("Internal Server error.", re, string(debug.Stack()))
		responses.InternalServerError(w, r, fmt.Errorf("%v", re))
	}
}

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

	defer Recover(w, r)

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
