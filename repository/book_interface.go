package repository

import (
	"database/sql"

	"github.com/arturbaccarin/library-api-with-tests/model"
)

type BookRepo interface {
	Connection() *sql.DB
	List() ([]model.Book, error)
}
