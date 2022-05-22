package cache

import (
	"errors"
	"log"
	"sync"

	"github.com/juanmachuca95/classdigitalhouse/services/books/models"
)

type Cache interface {
	GetBook(string) (models.Book, error)
}

type CacheMemory struct {
	cache map[string]models.Book
	lock  sync.Mutex
}

func NewCacheMemory() Cache {
	return &CacheMemory{
		cache: make(map[string]models.Book),
	}
}

func (c *CacheMemory) GetBook(name string) (models.Book, error) {
	results, exists := c.cache[name]
	if !exists {
		return models.Book{}, errors.New("no existe en cache")
	}
	log.Println("Si existe en cache")
	return results, nil
}

func (c *CacheMemory) AddBook(book models.Book) bool {
	if len(c.cache) == 0 {
		log.Println("El cache de libros esta vacio - ", len(c.cache))
		c.cache = make(map[string]models.Book)
		c.cache[book.Book] = book
		return true
	}

	c.cache[book.Book] = book
	return true
}
