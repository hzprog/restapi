package routehandler

import (
	"net/http"

	Book "github.com/hzprog/restapi/Controllers/book"

	"github.com/gorilla/mux"
)

func InitilizeRouter() *mux.Router {
	router := mux.NewRouter()
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static", http.FileServer(http.Dir("./static"))))

	api := router.PathPrefix("/api").Subrouter()
	api.HandleFunc("/books", Book.GetBooks).Methods("GET")
	api.HandleFunc("/books/{id}", Book.GetBook).Methods("GET")
	api.HandleFunc("/books", Book.CreateBook).Methods("POST")
	api.HandleFunc("/books/{id}", Book.UpdateBook).Methods("PUT")
	api.HandleFunc("/books/{id}", Book.DeleteBook).Methods("DELETE")

	return router
}
