package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

var orders = []Order{}

type Handler struct{}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) Check(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "Продюсер запущен",
	})
}

func (h *Handler) Post(c *gin.Context) {
	var order Order

	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Некорректные данные заказа",
			"info":  err.Error(),
		})
		return
	}

	order.CreatedAt = time.Now()
	order.UpdatedAt = time.Now()

	orders = append(orders, order)

	c.JSON(http.StatusCreated, gin.H{
		"message": "Заказ успешно получен",
		"order":   order,
	})
}

func (h *Handler) GetAll(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"orders": orders,
	})
}

func (h *Handler) GetByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Некорректный ID",
		})
		return
	}

	for _, order := range orders {
		if order.ID == id {
			c.JSON(http.StatusOK, gin.H{
				"order": order,
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"error": "Заказ не найден",
	})
}
