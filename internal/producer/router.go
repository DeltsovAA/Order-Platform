package producer

import (
	"github.com/gin-gonic/gin"
	"order-platform/internal/producer/handlers"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	handler := handlers.New()

	router.GET("/health", handler.Check)
	router.POST("/order", handler.Post)
	router.GET("/orders", handler.GetAll)
	router.GET("/orders/:id", handler.GetByID)

	return router
}
