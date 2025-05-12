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

func (h *Handler) Delete(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Некорректный ID",
		})
		return
	}

	for i, order := range orders {
		if order.ID == id {
			orders = append(orders[:i], orders[i+1:]...)
			c.JSON(http.StatusOK, gin.H{
				"message": "Заказ успешно удалён",
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"error": "Заказ не найден",
	})
}

func (h *Handler) Update(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Некорректный ID",
		})
		return
	}

	var updatedOrder Order
	if err := c.ShouldBindJSON(&updatedOrder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Некорректные данные для обновления",
			"info":  err.Error(),
		})
		return
	}

	for i, order := range orders {
		if order.ID == id {

			orders[i].Status = updatedOrder.Status
			orders[i].TotalAmount = updatedOrder.TotalAmount
			orders[i].UpdatedAt = time.Now()
			orders[i].Items = updatedOrder.Items
			orders[i].ShippingInfo = updatedOrder.ShippingInfo

			c.JSON(http.StatusOK, gin.H{
				"message": "Заказ успешно обновлён",
				"order":   orders[i],
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"error": "Заказ не найден",
	})
}
