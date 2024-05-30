package handler

import (
	"net/http"

	"github.com/arturbaccarin/library-api-with-tests/model"
	"github.com/gin-gonic/gin"
)

type BookMock struct{}

func (BookMock) GetList(c *gin.Context) {
	books := []model.Book{
		{Title: "Lord of the Rings: The Fellowship of the Ring", AuthorID: 1},
		{Title: "The Hobbit", AuthorID: 1},
		{Title: "Dune", AuthorID: 2},
	}

	c.JSON(http.StatusOK, books)
}
