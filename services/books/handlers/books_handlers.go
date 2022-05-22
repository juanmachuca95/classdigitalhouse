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
	tiempo := time.Now()
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "No ha especificado los parametros.")
		return
	}

	var request struct{ Book string }
	err = json.Unmarshal(reqBody, &request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "No ha sido posible parsear la request.")
		return
	}

	book, err := s.cache.GetBook(request.Book)
	if err != nil {
		log.Printf("El libro %s no esta en cache.\n", request.Book)

		book, err = s.gtw.GetBook(request.Book)
		if err != nil {
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "No existe el libro %s en nuestra base de datos\nTiempo de consulta ⏰ %v", request.Book, time.Since(tiempo))
			return
		}

		/* Agregamos a cache */
		s.cache.AddBook(book)

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Datos desde base de datos 👎\nLibro: %s\nAutor:%s\nTiempo de busqueda ⏰  %v", book.Book, book.Author, time.Since(tiempo))
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Datos desde cache 👍\nLibro: %s\nAutor:%s\nTiempo de busqueda ⏰ %v", book.Book, book.Author, time.Since(tiempo))
}
