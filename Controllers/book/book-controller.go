package Book

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path"

	helpers "github.com/hzprog/restapi/Helpers"

	configdb "github.com/hzprog/restapi/DBConfig"
	Book "github.com/hzprog/restapi/Models/book"

	"github.com/gorilla/mux"
)

//get all books
func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var books []Book.Book

	configdb.Db.Find(&books)

	json.NewEncoder(w).Encode(books)
}

//get a book with his id
func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	var book Book.Book
	configdb.Db.First(&book, params["id"])

	json.NewEncoder(w).Encode(book)
}

//create a book
func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	tempFile, err := helpers.Upload(r, "image")

	if err != nil {
		fmt.Println(err)
	}

	var book Book.Book

	book.Isbn = r.FormValue("isbn")
	book.Title = r.FormValue("title")
	book.Author = r.FormValue("author")
	book.Image = path.Base(tempFile.Name())

	configdb.Db.Create(&book)

	json.NewEncoder(w).Encode(book)
}

//update a book
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	var book Book.Book
	configdb.Db.First(&book, params["id"])

	json.NewDecoder(r.Body).Decode(&book)

	configdb.Db.Save(&book)

	json.NewEncoder(w).Encode(book)
}

//delete a book
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	var book Book.Book
	configdb.Db.Delete(&book, params["id"])

	json.NewEncoder(w).Encode("The book has been deleted successfully")
}
