package cache

import (
	"errors"
	"sync"

	models "github.com/juanmachuca95/classdigitalhouse/services/books/models"
)

type Cache interface {
	GetBook(string) (models.Book, error)
	AddBook(models.Book) bool
}

type CacheMemory struct {
	cache map[string]models.Book
	l     sync.Mutex
}

func NewCacheMemory() Cache {
	return &CacheMemory{
		cache: make(map[string]models.Book),
	}
}

func (c *CacheMemory) GetBook(bookName string) (models.Book, error) {
	results, exists := c.cache[bookName]
	if !exists {
		return models.Book{}, errors.New("Este libro o existe en cache")
	}
	return results, nil
}

func (c *CacheMemory) AddBook(book models.Book) bool {
	if len(c.cache) == 0 {
		c.l.Lock()
		c.cache = make(map[string]models.Book)
		c.cache[book.Book] = book
		c.l.Unlock()
		return true
	}

	c.l.Lock()
	c.cache[book.Book] = book
	c.l.Unlock()
	return true
}
