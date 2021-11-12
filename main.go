package main

import (
	"log"
	"net/http"

	configdb "github.com/hzprog/restapi/DBConfig"
	models "github.com/hzprog/restapi/Models"
	routeHandler "github.com/hzprog/restapi/RouteHandler"
)

func main() {
	configdb.Config()
	models.InitilizeMigation()
	router := routeHandler.InitilizeRouter()

	log.Fatal(http.ListenAndServe(":8000", router))
}
