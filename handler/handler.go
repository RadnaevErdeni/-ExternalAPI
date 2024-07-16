package handler

import (
	"tt/service"

	"github.com/gin-gonic/gin"
	swagFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	user := router.Group("/users")
	{
		user.GET("", h.getUserByPas)

	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swagFiles.Handler))

	return router
}
