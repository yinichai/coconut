package main

import (
	"log"
	"net/http"
	"os"

	. "github.com/yinichai/coconut/env"

	"github.com/yinichai/coconut/routes"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	router := routes.NewRouter()
	log.Fatal(http.ListenAndServe(":9060", router))
	defer DB.Close()
}
