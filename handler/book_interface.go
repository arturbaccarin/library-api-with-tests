package handler

import (
	"github.com/gin-gonic/gin"
)

type BookHandler interface {
	GetList(c *gin.Context)
}
