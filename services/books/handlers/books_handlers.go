package books

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
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

func (s *ServiceHTTPBooks) GetBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category: %v\n", vars["category"])
}
