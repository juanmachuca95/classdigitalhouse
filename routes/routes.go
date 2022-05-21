package routes

import (
	"github.com/gorilla/mux"
	serviceBooks "github.com/juanmachuca95/classdigitalhouse/services/books/handlers"
)

func NewRoutes() *mux.Router {
	router := mux.NewRouter()

	/* Servicios*/
	books := serviceBooks.NewServiceHTTPBooks()
	router.HandleFunc("/getbook", books.GetBook)

	return router
}
