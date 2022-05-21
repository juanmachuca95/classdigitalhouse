package books

import (
	"log"

	db "github.com/juanmachuca95/classdigitalhouse/database"
	"github.com/juanmachuca95/classdigitalhouse/services/books/models"
)

type GatewayBook interface {
	GetBook(string) (models.Book, error)
}

type ServiceBook struct {
	*db.MySQLClient
}

func NewGatewayBook() GatewayBook {
	return &ServiceBook{
		db.NewMySQLClient(),
	}
}

func (s *ServiceBook) GetBook(bookName string) (models.Book, error) {
	stmt, err := s.Prepare("SELECT id, book, author, created_at, updated_at FROM books")
	if err != nil {
		log.Fatal("Error al preparar la consulta - error: ", err)
	}
	defer stmt.Close()

	var book models.Book
	err = stmt.QueryRow(bookName).Scan(&book.Id, &book.Book, &book.Author, &book.Created_at, &book.Updated_at)
	if err != nil {
		return book, err
	}

	return book, nil
}
