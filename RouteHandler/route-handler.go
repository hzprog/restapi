package routehandler

import (
	"net/http"

	_ "github.com/hzprog/restapi/docs"

	Book "github.com/hzprog/restapi/Controllers/book"
	User "github.com/hzprog/restapi/Controllers/user"
	Mid "github.com/hzprog/restapi/Middlewares"

	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/gorilla/mux"
)

func InitializeRouter() *mux.Router {
	router := mux.NewRouter()
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static", http.FileServer(http.Dir("./static"))))

	api := router.PathPrefix("/api").Subrouter()

	//routes for book

	// Read-all
	api.HandleFunc("/books", Mid.IsAuthorized(Book.GetBooks)).Methods("GET")

	// Read one
	api.HandleFunc("/books/{id:[0-9]+}", Mid.IsAuthorized(Book.GetBook)).Methods("GET")

	// Create
	api.HandleFunc("/books", Mid.IsAuthorized(Book.CreateBook)).Methods("POST")
	api.HandleFunc("/books/{id:[0-9]+}", Mid.IsAuthorized(Book.UpdateBook)).Methods("put")
	api.HandleFunc("/books/{id:[0-9]+}", Mid.IsAuthorized(Book.DeleteBook)).Methods("DELETE")

	//routes for user
	api.HandleFunc("/signup", User.Signup).Methods("POST")
	api.HandleFunc("/login", User.Login).Methods("POST")
	api.HandleFunc("/users/{id:[0-9]+}", User.UpdateUser).Methods("PUT")
	api.HandleFunc("/users/{id:[0-9]+}", User.DeleteUser).Methods("DELETE")

	// Swagger
	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	return router
}
