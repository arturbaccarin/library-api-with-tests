package handler

import (
	"net/http"

	"github.com/arturbaccarin/library-api-with-tests/repository"
	"github.com/gin-gonic/gin"
)

type Book struct {
	DBRepo repository.BookRepo
}

func NewBookHandler(dbRepo repository.BookRepo) *Book {
	return &Book{
		DBRepo: dbRepo,
	}
}

func (b Book) GetList(c *gin.Context) {
	books, err := b.DBRepo.List()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, books)
}
