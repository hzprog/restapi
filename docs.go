// Documentation for Book API
//
//
//     Schemes: http
//     Host: localhost
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
	Book "github.com/hzprog/restapi/Models/book"
)

// A list of all books returns in the response
// swagger:response booksResponse
type booksResponseWrapper struct {
	// all books in the system
	//in:body
	books []Book.Book
}

// The book created returns in the response
// swagger:response bookResponse
type bookResponseWrapper struct {
	// the book created
	//in:body
	book Book.Book
}

// swagger:parameters createBook
type bookParamsWrapper struct {
	// Book data structure to Update or Create.
	// Note: the id field is ignored by update and create operations
	// in: body
	// required: true
	book Book.Book
}

// swagger:parameters deleteBook
type bookIDParameterWrapper struct {
	// The ID of the book to delete from the database
	//in:path
	//required: true
	Id int `json:"id"`
}

// swagger:response noContent
type noContentResponse struct {
	message string `json:"book deleted successfully"`
}

// title  string `json:"title"`
// 	isbn   string `json:"isbn"`
// 	author string `json:"author"`
// 	image  string `json:"image"`
