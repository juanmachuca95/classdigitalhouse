package main

import (
	"log"
	"net/http"

	"github.com/juanmachuca95/classdigitalhouse/routes"
	"github.com/juanmachuca95/classdigitalhouse/services/books/models"
)

type CacheMemory struct {
	cache map[string]models.Book
}

func NewCacheMemory() *CacheMemory {
	return &CacheMemory{
		cache: make(map[string]models.Book),
	}
}

func (c *CacheMemory) GetBook(name string) {}

func main() {
	cacheMemory := NewCacheMemory()
	routes := routes.NewRoutes()

	log.Println("Cache", cacheMemory.cache)
	err := http.ListenAndServe(":8080", routes)
	if err != nil {
		log.Fatal("No se puede inicializar el servidor - error: ", err)
	}

}
