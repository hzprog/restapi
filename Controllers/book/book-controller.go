package Book

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"strconv"

	helpers "github.com/hzprog/restapi/Helpers"

	configdb "github.com/hzprog/restapi/DBConfig"
	Book "github.com/hzprog/restapi/Models/book"

	"github.com/gorilla/mux"
)

// GetBooks godoc
// @Summary Get details of all books
// @Description Get details of all books
// @Tags orders
// @Accept  json
// @Produce  json
// @Success 200 {array} Books
// @Router /books [get]

//get all books
func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var books []Book.Book
	var total int64

	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))

	offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))

	configdb.Db.Model(&Book.Book{}).Count(&total)

	err := configdb.Db.Offset(offset).Limit(limit).Find(&books).Error

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Error couldn't find when retreiving books")
		fmt.Println(err)
		return
	}

	if len(books) < 1 {
		json.NewEncoder(w).Encode("no book found try adding a book")
		return
	}

	booksData := map[string]interface{}{
		"books": books,
		"total": total,
	}

	data, _ := json.Marshal(booksData)

	w.Write(data)
}

//get a book with his id
func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	var book Book.Book

	err := configdb.Db.First(&book, params["id"]).Error
	if err != nil {
		json.NewEncoder(w).Encode("Can't find a book with that id")
		fmt.Println(err)
		return
	}

	json.NewEncoder(w).Encode(book)
}

// Create Bookgodoc
// @Summary Create a new Book
// @Description Create a new Book with the input paylod
// @Tags Books
// @Accept  json
// @Produce  json
// @Param order body Book true "Create book"
// @Success 200 {object} Order
// @Router /orders [post]

//create a book
func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	uploadedFile := helpers.UploadFile(r, "image")

	var book Book.Book

	book.Isbn = r.FormValue("isbn")
	book.Title = r.FormValue("title")
	book.Author = r.FormValue("author")
	book.Image = path.Base(uploadedFile.Name())

	err := configdb.Db.Create(&book).Error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Error couldn't create the book")
		fmt.Println(err)
		return
	}

	json.NewEncoder(w).Encode(book)
}

//update a book
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	var book Book.Book

	err := configdb.Db.First(&book, params["id"]).Error
	if err != nil {
		json.NewEncoder(w).Encode("Can't find a book with that id")
		fmt.Println(err)
		return
	}

	json.NewDecoder(r.Body).Decode(&book)

	err = configdb.Db.Save(&book).Error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Error couldn't update the book")
		fmt.Println(err)
		return
	}

	json.NewEncoder(w).Encode(book)
}

//delete a book
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	var book Book.Book

	err := configdb.Db.First(&book, params["id"]).Error
	if err != nil {
		json.NewEncoder(w).Encode("Can't find a book with that id")
		fmt.Println(err)
		return
	}

	helpers.DeleteFile(book.Image)

	err = configdb.Db.Delete(&book).Error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Error couldn't Delete the book")
		fmt.Println(err)
		return
	}

	json.NewEncoder(w).Encode("The book has been deleted successfully")
}
