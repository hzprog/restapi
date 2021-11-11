package main

import (
	"log"
	"net/http"

	book "github.com/hzprog/restapi/book"

	"github.com/gorilla/mux"
)

func initilizeRouter() {
	r := mux.NewRouter()
	r.HandleFunc("/api/books", book.GetBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", book.GetBook).Methods("GET")
	r.HandleFunc("/api/books", book.CreateBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", book.UpdateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", book.DeleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}

func main() {

	book.InitilizeMigation()
	initilizeRouter()
}
