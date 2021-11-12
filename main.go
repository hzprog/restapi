package main

import (
	"fmt"
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

	port := Env.GetEnvVar("APP_PORT")

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), router))
}
