package main

import (
	"log"
	"net/http"

	"github.com/hzprog/restapi/user"

	"github.com/gorilla/mux"
)

func initilizeRouter() {
	r := mux.NewRouter()
	r.HandleFunc("/api/books", user.GetBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", user.GetBook).Methods("GET")
	r.HandleFunc("/api/books", user.CreateBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", user.UpdateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", user.DeleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}

func main() {

	user.InitilizeMigation()
	initilizeRouter()
}
