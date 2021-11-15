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
	api.HandleFunc("/books/{limit:[0-9]+}", Book.GetBooks).Methods("GET")
	api.HandleFunc("/book/{id:[0-9]+}", Book.GetBook).Methods("GET")
	api.HandleFunc("/books", Book.CreateBook).Methods("POST")
	api.HandleFunc("/books/{id:[0-9]+}", Book.UpdateBook).Methods("PUT")
	api.HandleFunc("/books/{id:[0-9]+}", Book.DeleteBook).Methods("DELETE")

	return router
}
