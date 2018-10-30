package handlers

import (
	"net/http"

	"github.com/yinichai/coconut/handlers/responses"
)

//Ping health check
func Ping() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			bytes = []byte("PONG")
		)
		responses.OK(w, bytes)
	}
}
