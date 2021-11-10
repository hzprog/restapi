package user

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const Dsn = "root:password@tcp(localhost)/demo?parseTime=true"

var Db *gorm.DB
var Err error

type Book struct {
	gorm.Model

	Isbn   string `json:"isbn"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

func InitilizeMigation() {
	Db, Err = gorm.Open(mysql.Open(Dsn), &gorm.Config{})
	if Err != nil {
		panic("failed to connect database")
	}

	//migrate the schemes
	Db.AutoMigrate(&Book{})
}

//get all books
func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var books []Book

	Db.Find(&books)

	json.NewEncoder(w).Encode(books)
}

//get a book with his id
func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	var book Book
	Db.First(&book, params["id"])

	json.NewEncoder(w).Encode(book)
}

//create a book
func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var book Book
	json.NewDecoder(r.Body).Decode(&book)
	Db.Create(&book)

	json.NewEncoder(w).Encode(book)
}

//update a book
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	var book Book
	Db.First(&book, params["id"])

	json.NewDecoder(r.Body).Decode(&book)

	Db.Save(&book)

	json.NewEncoder(w).Encode(book)
}

//delete a book
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	var book Book
	Db.Delete(&book, params["id"])

	json.NewEncoder(w).Encode("The book has been deleted successfully")
}
