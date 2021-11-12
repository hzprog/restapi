package main

import (
	"log"
	"net/http"

	configdb "github.com/hzprog/restapi/DBConfig"
	Env "github.com/hzprog/restapi/Helpers"
	models "github.com/hzprog/restapi/Models"
	routeHandler "github.com/hzprog/restapi/RouteHandler"
)

func main() {

	configdb.Config()
	models.InitilizeMigation()
	router := routeHandler.InitilizeRouter()

	log.Fatal(http.ListenAndServe(":"+Env.GetEnvVar("APP_PORT"), router))
}
