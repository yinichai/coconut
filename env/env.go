package env

import (
	"log"

	"github.com/jmoiron/sqlx"

	"github.com/kelseyhightower/envconfig"
	_ "github.com/mattn/go-sqlite3"
)

const appname = "fatlama"

// Env defines enviroment variables
var Env struct{}

var (
	DB *sqlx.DB
)

func init() {
	envconfig.MustProcess(appname, &Env)

	var err error
	DB, err = sqlx.Open("sqlite3", "./fatlama.sqlite3")
	log.Print("db: ", DB)
	if err != nil {
		panic(err)
	}
}
