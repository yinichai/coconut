package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

func RespondError(w http.ResponseWriter, method, url string, code int, err error) {
	RespondErrorWithEmptyBody(w, method, url, code, err)

	json.NewEncoder(w).Encode(map[string]interface{}{
		"error":       err.Error(),
		"code":        code,
		"description": http.StatusText(code),
	})
}

func RespondErrorWithEmptyBody(w http.ResponseWriter, method, url string, code int, err error) {
	log.Printf("%s %s: HTTP %d: %s", method, url, code, err)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(code)
}

func BadRequest(w http.ResponseWriter, r *http.Request, err error) {
	RespondError(w, r.Method, r.URL.String(), http.StatusBadRequest, err)
}

func InternalServerError(w http.ResponseWriter, r *http.Request, err error) {
	RespondError(w, r.Method, r.URL.String(), http.StatusInternalServerError, err)
}

func NotFound(w http.ResponseWriter, r *http.Request, err error) {
	RespondError(w, r.Method, r.URL.String(), http.StatusNotFound, err)
}
