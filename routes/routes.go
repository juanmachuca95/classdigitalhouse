package routes

import (
	"github.com/gorilla/mux"
	serviceBooks "github.com/juanmachuca95/classdigitalhouse/services/books/handlers"
)

func NewRoutes() *mux.Router {
	router := mux.NewRouter()

	/* Servicios*/
	books := serviceBooks.NewServiceHTTPBooks()

	// GET
	/* router.HandleFunc("/getbook/{book}", books.GetBook).
	Methods("GET") */

	// POST
	router.HandleFunc("/getbook", books.GetBook).
		Methods("POST")
	return router
}
