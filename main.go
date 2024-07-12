package main

import (
	"net/http"

	"github.com/KiteShi/golang-exercise/cmd"
)

var (
	//TODO: remove temporary hardcoded data
	apiPort = ":8080"
)

func main() {

	http.ListenAndServe(apiPort, cmd.GetRouter())
}
