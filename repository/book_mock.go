package repository

import (
	"database/sql"

	"github.com/arturbaccarin/library-api-with-tests/model"
)

type BookMockDB struct{}

func (BookMockDB) Connection() *sql.DB {
	return nil
}

func (BookMockDB) List() ([]model.Book, error) {
	books := []model.Book{
		{Title: "Lord of the Rings: The Fellowship of the Ring", AuthorID: 1},
		{Title: "The Hobbit", AuthorID: 1},
		{Title: "Dune", AuthorID: 2},
	}

	return books, nil
}
