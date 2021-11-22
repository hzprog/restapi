package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	configdb "github.com/hzprog/restapi/DBConfig"
	Env "github.com/hzprog/restapi/Helpers"
	models "github.com/hzprog/restapi/Models"
	routeHandler "github.com/hzprog/restapi/RouteHandler"
)

func main() {
	configdb.Config()
	models.InitializeMigration()
	router := routeHandler.InitializeRouter()

	port := Env.GetEnvVar("APP_PORT")

	methodsAllowed := []string{"POST", "GET", "DELETE", "OPTIONS", "HEAD", "PUT"}

	origins := handlers.AllowedOrigins([]string{"*"})
	methods := handlers.AllowedMethods(methodsAllowed)
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Accept", "Accept-Language", "Content-Type", "Authorization"})

	log.Printf("Server started on: http://localhost:%s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), handlers.CORS(headers, methods, origins)(router)))
}
