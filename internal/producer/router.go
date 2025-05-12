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
	router.DELETE("/order/:id", handler.Delete)
	router.PUT("/order/:id", handler.Update)

	return router
}
