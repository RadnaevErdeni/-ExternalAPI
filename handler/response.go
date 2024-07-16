package handler

import (
	"log"

	"github.com/gin-gonic/gin"
)

type errorz struct {
	Message string `json:"message"`
}

func errResponse(c *gin.Context, statusCode int, message string) {
	log.Fatal(message)
	c.AbortWithStatusJSON(statusCode, errorz{Message: message})
}
