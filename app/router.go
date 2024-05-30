package main

import (
	"github.com/arturbaccarin/library-api-with-tests/handler"
	"github.com/gin-gonic/gin"
)

func SetupRouter(bookHandler handler.BookHandler) *gin.Engine {
	router := gin.Default()
	router.GET("/books", bookHandler.GetList)

	return router
}
