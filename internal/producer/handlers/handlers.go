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

	// Проверка на корректность JSON
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Некорректные данные заказа",
			"info":  err.Error(),
		})
		return
	}

	// Проверка на отсутствие ID в запросе
	if order.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Поле ID обязательно для заполнения",
		})
		return
	}

	// Проверка на существующий ID
	for _, existingOrder := range orders {
		if existingOrder.ID == order.ID {
			c.JSON(http.StatusConflict, gin.H{
				"error":   "Заказ с таким ID уже существует",
				"orderID": order.ID,
			})
			return
		}
	}

	// Устанавливаем время создания и обновления
	order.CreatedAt = time.Now()
	order.UpdatedAt = time.Now()

	// Добавляем заказ в хранилище
	orders = append(orders, order)

	c.JSON(http.StatusCreated, gin.H{
		"message": "Заказ успешно получен",
		"order":   order,
	})
}

func (h *Handler) GetAll(c *gin.Context) {
	if len(orders) == 0 {
		// Если заказов нет, возвращаем пустой массив
		c.JSON(http.StatusOK, gin.H{
			"error": "Список заказов пуст",
		})
		return
	}

	// Если заказы есть, возвращаем их
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

func (h *Handler) Put(c *gin.Context) {
	var updatedOrder Order

	// Проверяем, что тело запроса содержит правильные данные
	if err := c.ShouldBindJSON(&updatedOrder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Некорректные данные заказа",
			"info":  err.Error(),
		})
		return
	}

	// Проверка на обязательность ID
	if updatedOrder.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Поле ID обязательно для обновления",
		})
		return
	}

	// Ищем заказ с таким ID
	var orderIndex int
	var orderFound bool
	for i, order := range orders {
		if order.ID == updatedOrder.ID {
			// Обновляем найденный заказ
			orders[i].Status = updatedOrder.Status
			orders[i].TotalAmount = updatedOrder.TotalAmount
			orders[i].Items = updatedOrder.Items
			orders[i].ShippingInfo = updatedOrder.ShippingInfo
			orders[i].UpdatedAt = time.Now() // Обновляем время

			orderIndex = i
			orderFound = true
			break
		}
	}

	if !orderFound {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Заказ с таким ID не найден",
		})
		return
	}

	// Возвращаем обновленный заказ
	c.JSON(http.StatusOK, gin.H{
		"message": "Заказ успешно обновлен",
		"order":   orders[orderIndex],
	})
}

func (h *Handler) Patch(c *gin.Context) {
	// Получаем ID заказа из параметра URL
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Некорректный ID",
		})
		return
	}

	// Пробуем распарсить JSON-данные из тела запроса
	if err := c.ShouldBindJSON(&patchData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Некорректные данные",
			"info":  err.Error(),
		})
		return
	}

	// Проверяем, чтобы ID в запросе не изменялся
	if patchData.ID != 0 && patchData.ID != id {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Нельзя изменить поле ID",
		})
		return
	}

	// Ищем заказ с таким ID
	for i, order := range orders {
		if order.ID == id {
			// Обновляем данные заказа
			if patchData.Status != "" {
				orders[i].Status = OrderStatus(patchData.Status)
			}
			if patchData.TotalAmount != 0 {
				orders[i].TotalAmount = patchData.TotalAmount
			}
			if len(patchData.Items) > 0 {
				orders[i].Items = patchData.Items
			}
			if patchData.ShippingInfo != (ShippingInfo{}) {
				orders[i].ShippingInfo = patchData.ShippingInfo
			}

			// Обновляем дату последнего обновления
			orders[i].UpdatedAt = time.Now()

			// Возвращаем обновленный заказ
			c.JSON(http.StatusOK, gin.H{
				"message": "Заказ успешно обновлен",
				"order":   orders[i],
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"error": "Заказ не найден",
	})
}
