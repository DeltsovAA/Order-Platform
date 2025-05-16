package handlers

import "time"

type OrderStatus string

const (
	StatusPending   OrderStatus = "pending"
	StatusPaid      OrderStatus = "paid"
	StatusShipped   OrderStatus = "shipped"
	StatusDelivered OrderStatus = "delivered"
	StatusCanceled  OrderStatus = "canceled"
)

type Order struct {
	ID           uint64       `json:"id"`            // уникальный идентификатор заказа
	CustomerID   uint64       `json:"customer_id"`   // ID покупателя
	CreatedAt    time.Time    `json:"created_at"`    // дата создания заказа
	UpdatedAt    time.Time    `json:"updated_at"`    // дата последнего обновления
	Status       OrderStatus  `json:"status"`        // статус заказа
	Items        []OrderItem  `json:"items"`         // список товаров в заказе
	TotalAmount  float64      `json:"total_amount"`  // общая сумма заказа
	ShippingInfo ShippingInfo `json:"shipping_info"` // информация о доставке
}

type OrderItem struct {
	ProductID uint64  `json:"product_id"` // ID товара
	Quantity  int     `json:"quantity"`   // количество
	Price     float64 `json:"price"`      // цена за единицу
}

type ShippingInfo struct {
	Address string `json:"address"`
	City    string `json:"city"`
	Country string `json:"country"`
	ZipCode string `json:"zip_code"`
	Contact string `json:"contact"` // имя получателя
	Phone   string `json:"phone"`   // телефон получателя
}

var patchData struct {
	ID           uint64       `json:"id"` // Это поле будет игнорироваться
	Status       string       `json:"status"`
	TotalAmount  float64      `json:"total_amount"`
	Items        []OrderItem  `json:"items"`
	ShippingInfo ShippingInfo `json:"shipping_info"`
}
