package responses

import (
	"net/http"
)

func ResponseWithEmptyBody(w http.ResponseWriter, status int) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(status)
}

func Respond(w http.ResponseWriter, status int, bytes []byte) {
	ResponseWithEmptyBody(w, status)
	w.Write(bytes)
}

func OK(w http.ResponseWriter, bytes []byte) {
	Respond(w, http.StatusOK, bytes)
}

func NoContent(w http.ResponseWriter) {
	ResponseWithEmptyBody(w, http.StatusNoContent)
}
