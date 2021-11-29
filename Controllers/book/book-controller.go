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

// swagger:route GET /books books getBooks
// Return all books.
// responses:
//   200: booksResponse
//
// swagger:response booksResponse

//get all books
func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var books []Book.Book
	var total int64

	limit, noValidLimit := strconv.Atoi(r.URL.Query().Get("limit"))
	if noValidLimit != nil {
		limit = 5
	}

	offset, noValidOffset := strconv.Atoi(r.URL.Query().Get("offset"))

	if noValidOffset != nil {
		offset = 0
	}

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

// swagger:route GET /books/{id} books getOneBook
// returns a book by his id.
// responses:
//   200: bookResponse
//
// swagger:response bookResponse

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

// swagger:route POST /books books createBook
// Create a book.
// responses:
//   200: bookResponse
//
// swagger:response bookResponse

//create a book
func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	fmt.Println("working")

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

// swagger:route PUT /books/{id} books updateBook
// Update a book.
// responses:
//   200: bookResponse
//
// swagger:response bookResponse

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

// swagger:route DELETE /books/{id} books deleteBook
// Delete a book from the database
//
// responses:
//	201: noContent

//delete a book
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	fmt.Println("works")

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

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("The book has been deleted successfully")
}
