package main

import (
	"net/http"

	"github.com/KiteShi/golang-exercise/cmd"
	"github.com/KiteShi/golang-exercise/pkg/db"
)

var (
	//TODO: remove temporary hardcoded data
	apiPort = ":8080"

	dbCfg = db.DBConfig{
		Host:     "localhost",
		Port:     5432,
		User:     "test_user",
		Password: "test_password",
		DBName:   "test_db",
	}
)

func main() {
	db.InitDB(dbCfg)
	defer db.CloseDB()

	http.ListenAndServe(apiPort, cmd.GetRouter())
}
