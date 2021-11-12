package routehandler

import (
	Book "github.com/hzprog/restapi/Controllers/book"

	"github.com/gorilla/mux"
)

func InitilizeRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/books", Book.GetBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", Book.GetBook).Methods("GET")
	router.HandleFunc("/api/books", Book.CreateBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", Book.UpdateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", Book.DeleteBook).Methods("DELETE")

	return router
}
