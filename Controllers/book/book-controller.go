package Book

import (
	"encoding/json"
	// "fmt"
	"net/http"
	"path"
	"strconv"

	helpers "github.com/hzprog/restapi/Helpers"

	configdb "github.com/hzprog/restapi/DBConfig"
	Response "github.com/hzprog/restapi/Helpers"
	Book "github.com/hzprog/restapi/Models/book"

	"github.com/gorilla/mux"
)

// swagger:route GET /books Books getBooks
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

	if err := configdb.Db.Model(&Book.Book{}).Count(&total).Error; err != nil {
		Response.HttpError(w, http.StatusInternalServerError, "Error couldn't find when retreiving books", err)
		return
	}

	if err := configdb.Db.Offset(offset).Limit(limit).Find(&books).Error; err != nil {
		Response.HttpError(w, http.StatusInternalServerError, "Error couldn't find when retreiving books", err)
		return
	}

	if len(books) < 1 {
		Response.HttpError(w, http.StatusInternalServerError, "Error couldn't find when retreiving books", nil)
		return
	}

	booksData := map[string]interface{}{
		"books": books,
		"total": total,
	}

	Response.HttpResponse(w, http.StatusOK, booksData)
}

// swagger:route GET /books/{id} Books getOneBook
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

	if err := configdb.Db.First(&book, params["id"]).Error; err != nil {
		Response.HttpError(w, http.StatusInternalServerError, "Can't find a book with that id", err)
		return
	}

	// json.NewEncoder(w).Encode(book)
	bookData := map[string]interface{}{
		"book": book,
	}

	Response.HttpResponse(w, http.StatusOK, bookData)
}

// swagger:route POST /books Books createBook
// Create a book.
// responses:
//   200: bookResponse
//
// swagger:response bookResponse

//create a book
func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	uploadedFile, error := helpers.UploadFile(r, "image")

	if error != nil {
		Response.HttpError(w, http.StatusInternalServerError, "Error couldn't create the book", error)
		return
	}

	var book Book.Book

	book.Isbn = r.FormValue("isbn")
	book.Title = r.FormValue("title")
	book.Author = r.FormValue("author")
	book.Image = path.Base(uploadedFile)

	if err := configdb.Db.Create(&book).Error; err != nil {
		Response.HttpError(w, http.StatusInternalServerError, "Error couldn't create the book", err)
		return
	}

	bookData := map[string]interface{}{
		"book": book,
	}

	Response.HttpResponse(w, http.StatusOK, bookData)
}

// swagger:route PUT /books/{id} Books updateBook
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

	if err := configdb.Db.First(&book, params["id"]).Error; err != nil {
		Response.HttpError(w, http.StatusInternalServerError, "Can't find a book with that id", err)
		return
	}

	json.NewDecoder(r.Body).Decode(&book)

	if err := configdb.Db.Save(&book).Error; err != nil {
		Response.HttpError(w, http.StatusInternalServerError, "Error couldn't update the book", err)
		return
	}

	bookData := map[string]interface{}{
		"book": book,
	}

	Response.HttpResponse(w, http.StatusOK, bookData)
}

// swagger:route DELETE /books/{id} Books deleteBook
// Delete a book from the database
//
// responses:
//	201: noContent

//delete a book
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	var book Book.Book

	if err := configdb.Db.First(&book, params["id"]).Error; err != nil {
		Response.HttpError(w, http.StatusInternalServerError, "Can't find a book with that id", err)
		return
	}

	helpers.DeleteFile(book.Image)

	if err := configdb.Db.Delete(&book).Error; err != nil {
		Response.HttpError(w, http.StatusInternalServerError, "Error couldn't Delete the book", err)
		return
	}

	Response.HttpResponse(w, http.StatusCreated, "The book has been deleted successfully")
}
