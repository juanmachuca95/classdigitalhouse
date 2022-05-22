package books

import (
	"net/http"

	books "github.com/juanmachuca95/classdigitalhouse/services/books/gateway"
)

type ServiceHTTPBooks struct {
	gtw books.GatewayBook
}

func NewServiceHTTPBooks() *ServiceHTTPBooks {
	return &ServiceHTTPBooks{
		gtw: books.NewGatewayBook(),
	}
}

/* GET */
/* func (s *ServiceHTTPBooks) GetBook(w http.ResponseWriter, r *http.Request) {
	book := mux.Vars(r)["book"]
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "book: %v\n", book)
} */

func (s *ServiceHTTPBooks) GetBook(w http.ResponseWriter, r *http.Request) {
	// TODO implement
}
