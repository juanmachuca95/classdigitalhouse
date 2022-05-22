package books

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	cache "github.com/juanmachuca95/classdigitalhouse/cache"
	books "github.com/juanmachuca95/classdigitalhouse/services/books/gateway"
)

type ServiceHTTPBooks struct {
	gtw   books.GatewayBook
	cache cache.CacheMemory
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
	tiempo := time.Now() // tiempo
	reqBody, _ := ioutil.ReadAll(r.Body)
	var request struct{ Book string }
	json.Unmarshal(reqBody, &request)

	book, err := s.cache.GetBook(request.Book)
	if err != nil { // no esta en cache
		log.Println(err.Error())

		book, err = s.gtw.GetBook(request.Book)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "No existe en base de datos %s ⚠️\n", err.Error())
			return
		}

		log.Println("Agregado a cache")
		s.cache.AddBook(book)

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Obteniendo el libro %s desde base de datos 👎 - ⏰ tiempo %v\n", book.Book, time.Since(tiempo))
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Obteniendo el libro %s desde cache 👍 - ⏰ tiempo %v\n", book.Book, time.Since(tiempo))
	return
}
