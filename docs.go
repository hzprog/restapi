// Package classification Book API.
// Documentation for Book API
//
//
//     Schemes: http
//     Host: localhost:8000
//     BasePath: /api
//     Version: 1.0.0
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//     Security:
//     - Bearer:
//
//     SecurityDefinitions:
//     Bearer:
//          type: apiKey
//          name: Authorization
//          in: header
//
// swagger:meta
package main

import (
	"bytes"
	Book "github.com/hzprog/restapi/Models/book"
)

// The total of books
// swagger:response booksResponse
type booksResponseWrapper struct {
	//in:body
	Body struct {
		Success bool              `json:"success"`
		Data    booksResponseData `json:"data"`
	}
}

type booksResponseData struct {
	Books []Book.Book `json:"books"`
	Total int         `json:"total"`
}

//swagger:parameters getBooks
type booksParamsWrapper struct {
	// the book created
	// in:query
	Limit  string `json:"limit"`
	Offset string `json:"offset"`
}

// The book created returns in the response
// swagger:response bookResponse
type bookResponseWrapper struct {
	// the book created
	// in:body
	Body struct {
		Success bool             `json:"success"`
		Data    bookResponseData `json:"data"`
	}
}

type bookResponseData struct {
	Book Book.Book `json:"book"`
}

// swagger:parameters updateBook
type bookParamsWrapper struct {
	// Book data structure to Update or Create.
	// Note: the id field is ignored by update and create operations
	// in: body
	// required: true
	Body struct {
		Isbn   string `json:"isbn"`
		Title  string `json:"title"`
		Author string `json:"author"`
	}
}

// swagger:parameters deleteBook updateBook getOneBook
type bookIDParameterWrapper struct {
	// The ID of the book to delete or get or update from the database
	//in:path
	//required: true
	Id int `json:"id"`
}

// swagger:response noContent
type noContentResponse struct {
	//in: body
	Body struct {
		Success bool   `json:"success"`
		Data    string `json:"data"`
	}
}

// The from contains the create book
// swagger:parameters createBook
type formData struct {
	// in: formData
	Isbn string `json:"isbn"`
	// in: formData
	Title string `json:"title"`
	// in: formData
	Author string `json:"author"`
}

// Image contains the uploaded file data
// swagger:parameters createBook
type createBookFromDataParamsWrapper struct {
	// Image desc.
	//
	// in: formData
	//
	// swagger:file
	Image *bytes.Buffer `json:"image"`
}

// The body to pass to login
//swagger:parameters login signup
type loginParamsWrapper struct {
	//in: body
	//required: true
	Body struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
}

// The total of books
// swagger:response authResponse
type loginResponseWrapper struct {
	//in:body
	Body struct {
		Success bool   `json:"success"`
		Data    string `json:"data"`
	}
}
